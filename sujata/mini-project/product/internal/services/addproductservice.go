package services

import (
	"context"
	"product/internal/dao/mongodao"
	model "product/internal/dao/mongodao/models"
	"product/internal/errors"
	"product/util"
	"sync"

	log "github.com/sirupsen/logrus"
)

type AddProductService interface {
	ValidateRequest(product model.Product) *errors.ServerError
	ProcessRequest(ctx context.Context, product model.Product) *errors.ServerError
}

var addProductServiceStruct AddProductService
var addProductServiceOnce sync.Once

type addProductService struct {
	config *util.RouterConfig
}

func InitAddProductService(config *util.RouterConfig) AddProductService {
	addProductServiceOnce.Do(func() {
		addProductServiceStruct = &addProductService{
			config: config,
		}
	})

	return addProductServiceStruct
}

func GetAddProductService() AddProductService {
	if addProductServiceStruct == nil {
		panic("Add product service not initialised")
	}

	return addProductServiceStruct
}

func (s *addProductService) ValidateRequest(product model.Product) *errors.ServerError {

	return nil
}

func (service *addProductService) ProcessRequest(ctx context.Context, product model.Product) *errors.ServerError {
	dao := mongodao.GetMongoDAO()

	err := dao.AddProduct(ctx, product)
	if err != nil {
		log.WithField("Error: ", err).Error("an error occurred while inserting product in db")
		return err
	}

	return nil
}
