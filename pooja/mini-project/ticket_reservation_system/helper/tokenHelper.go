package helper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"ticket_reservation_system/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails
type SignedDetails struct {
	UserName string `json:"username"`
	EmailId  string `json:"email_id"`
	jwt.StandardClaims
}

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(username string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	fmt.Println(claims, refreshClaims)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

//UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, username string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"username": username}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	return
}

func CreateToken(username, email_id string) (string, error) {
	claims := &SignedDetails{
		UserName: username,
		EmailId:  email_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	fmt.Print("create token claims", claims)
	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	fmt.Print("create token claims", claims)
	return signedToken, err
}

//ValidateToken validates the jwt token
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return false, errors.New("Token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return false, errors.New("Session expired. Login again!!!")
	}

	return true, nil
}

func GetClaimsFromToken(tokenString string) (SignedDetails, error) {
	var claims SignedDetails
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		return claims, err
	}
	if token.Valid {
		return claims, nil
	}
	return claims, errors.New("invalid token")
}
