// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infrastructure/router/interface.go

// Package server is a generated GoMock package.
package server

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/maykonlf/mocker/internal/model/entities"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockRouter) Set(route, method string, response *entities.APIResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", route, method, response)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockRouterMockRecorder) Set(route, method, response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRouter)(nil).Set), route, method, response)
}

// Listen mocks base method
func (m *MockRouter) Listen() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen")
	ret0, _ := ret[0].(error)
	return ret0
}

// Listen indicates an expected call of Listen
func (mr *MockRouterMockRecorder) Listen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockRouter)(nil).Listen))
}
