// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/exec/executable.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// GetCmd mocks base method.
func (m *MockExecutor) GetCmd() (string, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCmd")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetCmd indicates an expected call of GetCmd.
func (mr *MockExecutorMockRecorder) GetCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCmd", reflect.TypeOf((*MockExecutor)(nil).GetCmd))
}

// Run mocks base method.
func (m *MockExecutor) Run(args ...string) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Run indicates an expected call of Run.
func (mr *MockExecutorMockRecorder) Run(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockExecutor)(nil).Run), args...)
}

// RunCombined mocks base method.
func (m *MockExecutor) RunCombined(args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCombined", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCombined indicates an expected call of RunCombined.
func (mr *MockExecutorMockRecorder) RunCombined(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCombined", reflect.TypeOf((*MockExecutor)(nil).RunCombined), args...)
}