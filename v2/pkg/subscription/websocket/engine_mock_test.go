// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/matthewmcneely/graphql-go-tools/v2/pkg/subscription (interfaces: Engine)

// Package websocket is a generated GoMock package.
package websocket

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	subscription "github.com/matthewmcneely/graphql-go-tools/v2/pkg/subscription"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// StartOperation mocks base method.
func (m *MockEngine) StartOperation(arg0 context.Context, arg1 string, arg2 []byte, arg3 subscription.EventHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartOperation", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartOperation indicates an expected call of StartOperation.
func (mr *MockEngineMockRecorder) StartOperation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartOperation", reflect.TypeOf((*MockEngine)(nil).StartOperation), arg0, arg1, arg2, arg3)
}

// StopSubscription mocks base method.
func (m *MockEngine) StopSubscription(arg0 string, arg1 subscription.EventHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopSubscription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopSubscription indicates an expected call of StopSubscription.
func (mr *MockEngineMockRecorder) StopSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopSubscription", reflect.TypeOf((*MockEngine)(nil).StopSubscription), arg0, arg1)
}

// TerminateAllSubscriptions mocks base method.
func (m *MockEngine) TerminateAllSubscriptions(arg0 subscription.EventHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateAllSubscriptions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateAllSubscriptions indicates an expected call of TerminateAllSubscriptions.
func (mr *MockEngineMockRecorder) TerminateAllSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateAllSubscriptions", reflect.TypeOf((*MockEngine)(nil).TerminateAllSubscriptions), arg0)
}
