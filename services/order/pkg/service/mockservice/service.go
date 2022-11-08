// Code generated by mockery v2.13.1. DO NOT EDIT.

package mockservice

import (
	domain "github.com/nndergunov/deliveryApp/services/order/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: order, accountID
func (_m *App) CreateOrder(order domain.Order, accountID int) (*domain.Order, error) {
	ret := _m.Called(order, accountID)

	var r0 *domain.Order
	if rf, ok := ret.Get(0).(func(domain.Order, int) *domain.Order); ok {
		r0 = rf(order, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Order, int) error); ok {
		r1 = rf(order, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnOrder provides a mock function with given fields: orderID
func (_m *App) ReturnOrder(orderID int) (*domain.Order, error) {
	ret := _m.Called(orderID)

	var r0 *domain.Order
	if rf, ok := ret.Get(0).(func(int) *domain.Order); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnOrderList provides a mock function with given fields: params
func (_m *App) ReturnOrderList(params domain.SearchParameters) ([]domain.Order, error) {
	ret := _m.Called(params)

	var r0 []domain.Order
	if rf, ok := ret.Get(0).(func(domain.SearchParameters) []domain.Order); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.SearchParameters) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrder provides a mock function with given fields: order
func (_m *App) UpdateOrder(order domain.Order) (*domain.Order, error) {
	ret := _m.Called(order)

	var r0 *domain.Order
	if rf, ok := ret.Get(0).(func(domain.Order) *domain.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: status
func (_m *App) UpdateStatus(status domain.OrderStatus) (*domain.OrderStatus, error) {
	ret := _m.Called(status)

	var r0 *domain.OrderStatus
	if rf, ok := ret.Get(0).(func(domain.OrderStatus) *domain.OrderStatus); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OrderStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.OrderStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewApp creates a new instance of App. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApp(t mockConstructorTestingTNewApp) *App {
	mock := &App{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}