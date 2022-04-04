package domain

import (
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/product/utils/errs"
)

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

func NewProduct(name string, description string, quantity int, price int) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Price:       price,
	}
}

type ProductRepositoryDB interface {
	Save(Product) (*Product, *errs.AppError)
	FetchProductById(string) (*Product, *errs.AppError)
	UpdateProduct(string, Product) (*Product, *errs.AppError)
	DeleteProductById(string) *errs.AppError
}
