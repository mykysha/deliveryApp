// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/accountingservice/accountingservice.go

// Package mock_accountingservice is a generated GoMock package.
package mock_accountingservice

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/nndergunov/deliveryApp/services/accounting/pkg/domain"
)

// MockAccountService is a mock of Service interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// DeleteAccount mocks base method.
func (m *MockAccountService) DeleteAccount(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockAccountServiceMockRecorder) DeleteAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountService)(nil).DeleteAccount), id)
}

// DeleteTransaction mocks base method.
func (m *MockAccountService) DeleteTransaction(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockAccountServiceMockRecorder) DeleteTransaction(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockAccountService)(nil).DeleteTransaction), id)
}

// GetAccountByID mocks base method.
func (m *MockAccountService) GetAccountByID(ID string) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ID)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockAccountServiceMockRecorder) GetAccountByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockAccountService)(nil).GetAccountByID), ID)
}

// GetAccountListByParam mocks base method.
func (m *MockAccountService) GetAccountListByParam(param domain.SearchParam) ([]domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountListByParam", param)
	ret0, _ := ret[0].([]domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountListByParam indicates an expected call of GetAccountListByParam.
func (mr *MockAccountServiceMockRecorder) GetAccountListByParam(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountListByParam", reflect.TypeOf((*MockAccountService)(nil).GetAccountListByParam), param)
}

// InsertNewAccount mocks base method.
func (m *MockAccountService) InsertNewAccount(account domain.Account) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewAccount", account)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNewAccount indicates an expected call of InsertNewAccount.
func (mr *MockAccountServiceMockRecorder) InsertNewAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewAccount", reflect.TypeOf((*MockAccountService)(nil).InsertNewAccount), account)
}

// InsertTransaction mocks base method.
func (m *MockAccountService) InsertTransaction(transaction domain.Transaction) (*domain.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTransaction", transaction)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTransaction indicates an expected call of InsertTransaction.
func (mr *MockAccountServiceMockRecorder) InsertTransaction(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockAccountService)(nil).InsertTransaction), transaction)
}
