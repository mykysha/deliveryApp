// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/clients/clients.go

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	consumer_proto "github.com/nndergunov/deliveryApp/services/consumer/api/v1/grpc/proto"
	courier_proto "github.com/nndergunov/deliveryApp/services/courier/api/v1/grpc/proto"
	pb "github.com/nndergunov/deliveryApp/services/restaurant/api/v1/grpclogic/pb"
)

// MockRestaurantClient is a mock of RestaurantClient interface.
type MockRestaurantClient struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantClientMockRecorder
}

// MockRestaurantClientMockRecorder is the mock recorder for MockRestaurantClient.
type MockRestaurantClientMockRecorder struct {
	mock *MockRestaurantClient
}

// NewMockRestaurantClient creates a new mock instance.
func NewMockRestaurantClient(ctrl *gomock.Controller) *MockRestaurantClient {
	mock := &MockRestaurantClient{ctrl: ctrl}
	mock.recorder = &MockRestaurantClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantClient) EXPECT() *MockRestaurantClientMockRecorder {
	return m.recorder
}

// GetRestaurant mocks base method.
func (m *MockRestaurantClient) GetRestaurant(restaurantID int) (*pb.RestaurantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", restaurantID)
	ret0, _ := ret[0].(*pb.RestaurantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockRestaurantClientMockRecorder) GetRestaurant(restaurantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockRestaurantClient)(nil).GetRestaurant), restaurantID)
}

// MockCourierClient is a mock of CourierClient interface.
type MockCourierClient struct {
	ctrl     *gomock.Controller
	recorder *MockCourierClientMockRecorder
}

// MockCourierClientMockRecorder is the mock recorder for MockCourierClient.
type MockCourierClientMockRecorder struct {
	mock *MockCourierClient
}

// NewMockCourierClient creates a new mock instance.
func NewMockCourierClient(ctrl *gomock.Controller) *MockCourierClient {
	mock := &MockCourierClient{ctrl: ctrl}
	mock.recorder = &MockCourierClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourierClient) EXPECT() *MockCourierClientMockRecorder {
	return m.recorder
}

// GetLocation mocks base method.
func (m *MockCourierClient) GetLocation(city string) (*courier_proto.LocationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", city)
	ret0, _ := ret[0].(*courier_proto.LocationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MockCourierClientMockRecorder) GetLocation(city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockCourierClient)(nil).GetLocation), city)
}

// UpdateCourierAvailable mocks base method.
func (m *MockCourierClient) UpdateCourierAvailable(courierID, available string) (*courier_proto.CourierResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourierAvailable", courierID, available)
	ret0, _ := ret[0].(*courier_proto.CourierResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourierAvailable indicates an expected call of UpdateCourierAvailable.
func (mr *MockCourierClientMockRecorder) UpdateCourierAvailable(courierID, available interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourierAvailable", reflect.TypeOf((*MockCourierClient)(nil).UpdateCourierAvailable), courierID, available)
}

// MockConsumerClient is a mock of ConsumerClient interface.
type MockConsumerClient struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerClientMockRecorder
}

// MockConsumerClientMockRecorder is the mock recorder for MockConsumerClient.
type MockConsumerClientMockRecorder struct {
	mock *MockConsumerClient
}

// NewMockConsumerClient creates a new mock instance.
func NewMockConsumerClient(ctrl *gomock.Controller) *MockConsumerClient {
	mock := &MockConsumerClient{ctrl: ctrl}
	mock.recorder = &MockConsumerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerClient) EXPECT() *MockConsumerClientMockRecorder {
	return m.recorder
}

// GetLocation mocks base method.
func (m *MockConsumerClient) GetLocation(id int64) (*consumer_proto.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", id)
	ret0, _ := ret[0].(*consumer_proto.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MockConsumerClientMockRecorder) GetLocation(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockConsumerClient)(nil).GetLocation), id)
}
