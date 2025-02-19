// Code generated by MockGen. DO NOT EDIT.
// Source: go.arcalot.io/arcaflow-container-toolkit/internal/ce_service (interfaces: ContainerEngineService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	docker "go.arcalot.io/arcaflow-container-toolkit/internal/docker"
)

// MockContainerEngineService is a mock of ContainerEngineService interface.
type MockContainerEngineService struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEngineServiceMockRecorder
}

// MockContainerEngineServiceMockRecorder is the mock recorder for MockContainerEngineService.
type MockContainerEngineServiceMockRecorder struct {
	mock *MockContainerEngineService
}

// NewMockContainerEngineService creates a new mock instance.
func NewMockContainerEngineService(ctrl *gomock.Controller) *MockContainerEngineService {
	mock := &MockContainerEngineService{ctrl: ctrl}
	mock.recorder = &MockContainerEngineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerEngineService) EXPECT() *MockContainerEngineServiceMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockContainerEngineService) Build(arg0, arg1 string, arg2 []string, arg3 *docker.BuildOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockContainerEngineServiceMockRecorder) Build(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockContainerEngineService)(nil).Build), arg0, arg1, arg2, arg3)
}

// Push mocks base method.
func (m *MockContainerEngineService) Push(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockContainerEngineServiceMockRecorder) Push(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockContainerEngineService)(nil).Push), arg0, arg1, arg2, arg3)
}

// Tag mocks base method.
func (m *MockContainerEngineService) Tag(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockContainerEngineServiceMockRecorder) Tag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockContainerEngineService)(nil).Tag), arg0, arg1)
}
