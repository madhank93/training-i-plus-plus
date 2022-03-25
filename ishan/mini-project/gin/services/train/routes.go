package train

import (
	middleware "swiggy/gin/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	router.POST("/train", middleware.CheckAuthMiddleware, createNewTrain)
	router.GET("/train", FetchTrains)
}
