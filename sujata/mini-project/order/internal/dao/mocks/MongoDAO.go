// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"
	errors "order/internal/errors"

	mock "github.com/stretchr/testify/mock"

	model "order/internal/dao/models"
)

// MongoDAO is an autogenerated mock type for the MongoDAO type
type MongoDAO struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, cartProduct
func (_m *MongoDAO) CreateOrder(ctx context.Context, cartProduct model.Order) (interface{}, *errors.ServerError) {
	ret := _m.Called(ctx, cartProduct)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, model.Order) interface{}); ok {
		r0 = rf(ctx, cartProduct)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 *errors.ServerError
	if rf, ok := ret.Get(1).(func(context.Context, model.Order) *errors.ServerError); ok {
		r1 = rf(ctx, cartProduct)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ServerError)
		}
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, email
func (_m *MongoDAO) GetOrders(ctx context.Context, email string) (model.AllOrders, *errors.ServerError) {
	ret := _m.Called(ctx, email)

	var r0 model.AllOrders
	if rf, ok := ret.Get(0).(func(context.Context, string) model.AllOrders); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(model.AllOrders)
	}

	var r1 *errors.ServerError
	if rf, ok := ret.Get(1).(func(context.Context, string) *errors.ServerError); ok {
		r1 = rf(ctx, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ServerError)
		}
	}

	return r0, r1
}

// SetOrderStatus provides a mock function with given fields: ctx, orderInfo
func (_m *MongoDAO) SetOrderStatus(ctx context.Context, orderInfo model.OrderInfo) *errors.ServerError {
	ret := _m.Called(ctx, orderInfo)

	var r0 *errors.ServerError
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderInfo) *errors.ServerError); ok {
		r0 = rf(ctx, orderInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ServerError)
		}
	}

	return r0
}
