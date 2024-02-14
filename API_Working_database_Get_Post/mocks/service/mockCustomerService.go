// Code generated by MockGen. DO NOT EDIT.
// Source: api/service (interfaces: CustomerService)

// Package service is a generated GoMock package.
package service

import (
	domain "api/domain"
	dto "api/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockCustomerService) CreateCustomer(arg0 domain.Customer) (*dto.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", arg0)
	ret0, _ := ret[0].(*dto.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockCustomerServiceMockRecorder) CreateCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockCustomerService)(nil).CreateCustomer), arg0)
}

// GetAllCustomers mocks base method.
func (m *MockCustomerService) GetAllCustomers() ([]dto.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomers")
	ret0, _ := ret[0].([]dto.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCustomers indicates an expected call of GetAllCustomers.
func (mr *MockCustomerServiceMockRecorder) GetAllCustomers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomers", reflect.TypeOf((*MockCustomerService)(nil).GetAllCustomers))
}

// GetCustomerById mocks base method.
func (m *MockCustomerService) GetCustomerById(arg0 int) (*dto.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerById", arg0)
	ret0, _ := ret[0].(*dto.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerById indicates an expected call of GetCustomerById.
func (mr *MockCustomerServiceMockRecorder) GetCustomerById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerById", reflect.TypeOf((*MockCustomerService)(nil).GetCustomerById), arg0)
}