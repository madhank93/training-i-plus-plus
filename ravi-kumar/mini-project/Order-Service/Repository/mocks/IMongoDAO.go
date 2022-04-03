// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	mockdata "Order-Service/model"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// IMongoDAO is an autogenerated mock type for the IMongoDAO type
type IMongoDAO struct {
	mock.Mock
}

// MongoDeleteOrderById provides a mock function with given fields: orderId
func (_m *IMongoDAO) MongoDeleteOrderById(orderId primitive.ObjectID) {
	_m.Called(orderId)
}

// MongoDeliverOrderByOrderId provides a mock function with given fields: orderId, order
func (_m *IMongoDAO) MongoDeliverOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (*mockdata.Order, error) {
	ret := _m.Called(orderId, order)

	var r0 *mockdata.Order
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, mockdata.Order) *mockdata.Order); ok {
		r0 = rf(orderId, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mockdata.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, mockdata.Order) error); ok {
		r1 = rf(orderId, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MongoGetOrderByOrderId provides a mock function with given fields: orderId
func (_m *IMongoDAO) MongoGetOrderByOrderId(orderId primitive.ObjectID) (*mockdata.Order, error) {
	ret := _m.Called(orderId)

	var r0 *mockdata.Order
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *mockdata.Order); ok {
		r0 = rf(orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mockdata.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MongoGetOrderByUserId provides a mock function with given fields: userId
func (_m *IMongoDAO) MongoGetOrderByUserId(userId string) ([]mockdata.Order, error) {
	ret := _m.Called(userId)

	var r0 []mockdata.Order
	if rf, ok := ret.Get(0).(func(string) []mockdata.Order); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mockdata.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MongoPlaceOrder provides a mock function with given fields: orderPlaced
func (_m *IMongoDAO) MongoPlaceOrder(orderPlaced mockdata.Order) string {
	ret := _m.Called(orderPlaced)

	var r0 string
	if rf, ok := ret.Get(0).(func(mockdata.Order) string); ok {
		r0 = rf(orderPlaced)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MongoUpdateOrderByOrderId provides a mock function with given fields: orderId, order
func (_m *IMongoDAO) MongoUpdateOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (*mockdata.Order, error) {
	ret := _m.Called(orderId, order)

	var r0 *mockdata.Order
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, mockdata.Order) *mockdata.Order); ok {
		r0 = rf(orderId, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mockdata.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, mockdata.Order) error); ok {
		r1 = rf(orderId, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
