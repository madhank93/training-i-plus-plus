// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	domain "order/domain"
	errs "order/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// OrderService is an autogenerated mock type for the OrderService type
type OrderService struct {
	mock.Mock
}

// CheckAndCalcAmount provides a mock function with given fields: _a0
func (_m *OrderService) CheckAndCalcAmount(_a0 []domain.OrderItem) (int, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func([]domain.OrderItem) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func([]domain.OrderItem) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: _a0
func (_m *OrderService) CreateOrder(_a0 domain.Order) (*domain.Order, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Order
	if rf, ok := ret.Get(0).(func(domain.Order) *domain.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.Order) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// FindById provides a mock function with given fields: _a0
func (_m *OrderService) FindById(_a0 string) (*domain.Order, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Order
	if rf, ok := ret.Get(0).(func(string) *domain.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// FindProductById provides a mock function with given fields: _a0
func (_m *OrderService) FindProductById(_a0 string) (*domain.Product, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(string) *domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: _a0, _a1
func (_m *OrderService) UpdateProduct(_a0 string, _a1 domain.Product) *errs.AppError {
	ret := _m.Called(_a0, _a1)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(string, domain.Product) *errs.AppError); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}
