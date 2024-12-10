// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -source=controller.go -destination=./mock/controller_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v2 "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockIBookController is a mock of IBookController interface.
type MockIBookController struct {
	ctrl     *gomock.Controller
	recorder *MockIBookControllerMockRecorder
	isgomock struct{}
}

// MockIBookControllerMockRecorder is the mock recorder for MockIBookController.
type MockIBookControllerMockRecorder struct {
	mock *MockIBookController
}

// NewMockIBookController creates a new mock instance.
func NewMockIBookController(ctrl *gomock.Controller) *MockIBookController {
	mock := &MockIBookController{ctrl: ctrl}
	mock.recorder = &MockIBookControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBookController) EXPECT() *MockIBookControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIBookController) Create(ctx *v2.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIBookControllerMockRecorder) Create(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBookController)(nil).Create), ctx)
}