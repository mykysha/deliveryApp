// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/courierservice/courierservice.go

// Package mock_courierservice is a generated GoMock package.
package mock_courierservice

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/nndergunov/deliveryApp/services/courier/pkg/domain"
)

// MockCourierService is a mock of CourierService interface.
type MockCourierService struct {
	ctrl     *gomock.Controller
	recorder *MockCourierServiceMockRecorder
}

// MockCourierServiceMockRecorder is the mock recorder for MockCourierService.
type MockCourierServiceMockRecorder struct {
	mock *MockCourierService
}

// NewMockCourierService creates a new mock instance.
func NewMockCourierService(ctrl *gomock.Controller) *MockCourierService {
	mock := &MockCourierService{ctrl: ctrl}
	mock.recorder = &MockCourierServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourierService) EXPECT() *MockCourierServiceMockRecorder {
	return m.recorder
}

// DeleteCourier mocks base method.
func (m *MockCourierService) DeleteCourier(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCourier", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCourier indicates an expected call of DeleteCourier.
func (mr *MockCourierServiceMockRecorder) DeleteCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCourier", reflect.TypeOf((*MockCourierService)(nil).DeleteCourier), id)
}

// GetCourier mocks base method.
func (m *MockCourierService) GetCourier(id string) (*domain.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourier", id)
	ret0, _ := ret[0].(*domain.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourier indicates an expected call of GetCourier.
func (mr *MockCourierServiceMockRecorder) GetCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourier", reflect.TypeOf((*MockCourierService)(nil).GetCourier), id)
}

// GetCourierList mocks base method.
func (m *MockCourierService) GetCourierList(params domain.SearchParam) ([]domain.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierList", params)
	ret0, _ := ret[0].([]domain.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierList indicates an expected call of GetCourierList.
func (mr *MockCourierServiceMockRecorder) GetCourierList(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierList", reflect.TypeOf((*MockCourierService)(nil).GetCourierList), params)
}

// GetLocation mocks base method.
func (m *MockCourierService) GetLocation(userID string) (*domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation", userID)
	ret0, _ := ret[0].(*domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MockCourierServiceMockRecorder) GetLocation(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockCourierService)(nil).GetLocation), userID)
}

// GetLocationList mocks base method.
func (m *MockCourierService) GetLocationList(param domain.SearchParam) ([]domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationList", param)
	ret0, _ := ret[0].([]domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationList indicates an expected call of GetLocationList.
func (mr *MockCourierServiceMockRecorder) GetLocationList(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationList", reflect.TypeOf((*MockCourierService)(nil).GetLocationList), param)
}

// InsertCourier mocks base method.
func (m *MockCourierService) InsertCourier(courier domain.Courier) (*domain.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCourier", courier)
	ret0, _ := ret[0].(*domain.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCourier indicates an expected call of InsertCourier.
func (mr *MockCourierServiceMockRecorder) InsertCourier(courier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCourier", reflect.TypeOf((*MockCourierService)(nil).InsertCourier), courier)
}

// InsertLocation mocks base method.
func (m *MockCourierService) InsertLocation(courier domain.Location, userID string) (*domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLocation", courier, userID)
	ret0, _ := ret[0].(*domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertLocation indicates an expected call of InsertLocation.
func (mr *MockCourierServiceMockRecorder) InsertLocation(courier, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLocation", reflect.TypeOf((*MockCourierService)(nil).InsertLocation), courier, userID)
}

// UpdateCourier mocks base method.
func (m *MockCourierService) UpdateCourier(courier domain.Courier, id string) (*domain.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourier", courier, id)
	ret0, _ := ret[0].(*domain.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourier indicates an expected call of UpdateCourier.
func (mr *MockCourierServiceMockRecorder) UpdateCourier(courier, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourier", reflect.TypeOf((*MockCourierService)(nil).UpdateCourier), courier, id)
}

// UpdateCourierAvailable mocks base method.
func (m *MockCourierService) UpdateCourierAvailable(id, available string) (*domain.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourierAvailable", id, available)
	ret0, _ := ret[0].(*domain.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourierAvailable indicates an expected call of UpdateCourierAvailable.
func (mr *MockCourierServiceMockRecorder) UpdateCourierAvailable(id, available interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourierAvailable", reflect.TypeOf((*MockCourierService)(nil).UpdateCourierAvailable), id, available)
}

// UpdateLocation mocks base method.
func (m *MockCourierService) UpdateLocation(courier domain.Location, id string) (*domain.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocation", courier, id)
	ret0, _ := ret[0].(*domain.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLocation indicates an expected call of UpdateLocation.
func (mr *MockCourierServiceMockRecorder) UpdateLocation(courier, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocation", reflect.TypeOf((*MockCourierService)(nil).UpdateLocation), courier, id)
}
