// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"
	errors "product/internal/errors"

	mock "github.com/stretchr/testify/mock"

	model "product/internal/dao/mongodao/models"
)

// MongoDAO is an autogenerated mock type for the MongoDAO type
type MongoDAO struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, user
func (_m *MongoDAO) AddProduct(ctx context.Context, user model.Product) *errors.ServerError {
	ret := _m.Called(ctx, user)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) *errors.ServerError); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}

// GetProduct provides a mock function with given fields: ctx, productId
func (_m *MongoDAO) GetProduct(ctx context.Context, productId string) (model.Product, *errors.ServerError) {
	ret := _m.Called(ctx, productId)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Product); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 *errors.ServerError
	if rf, ok := ret.Get(1).(func(context.Context, string) *errors.ServerError); ok {
		r1 = rf(ctx, productId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ServerError)
		}
	}

	return r0, r1
}
