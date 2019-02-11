// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moiot/gravity/pkg/position_store (interfaces: PositionCacheInterface)

// Package mock_position_store is a generated GoMock package.
package mock_position_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	position_store "github.com/moiot/gravity/pkg/position_store"
)

// MockPositionCacheInterface is a mock of PositionCacheInterface interface
type MockPositionCacheInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPositionCacheInterfaceMockRecorder
}

// MockPositionCacheInterfaceMockRecorder is the mock recorder for MockPositionCacheInterface
type MockPositionCacheInterfaceMockRecorder struct {
	mock *MockPositionCacheInterface
}

// NewMockPositionCacheInterface creates a new mock instance
func NewMockPositionCacheInterface(ctrl *gomock.Controller) *MockPositionCacheInterface {
	mock := &MockPositionCacheInterface{ctrl: ctrl}
	mock.recorder = &MockPositionCacheInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPositionCacheInterface) EXPECT() *MockPositionCacheInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method
func (m *MockPositionCacheInterface) Clear() error {
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockPositionCacheInterfaceMockRecorder) Clear() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockPositionCacheInterface)(nil).Clear))
}

// Close mocks base method
func (m *MockPositionCacheInterface) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockPositionCacheInterfaceMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPositionCacheInterface)(nil).Close))
}

// Flush mocks base method
func (m *MockPositionCacheInterface) Flush() error {
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockPositionCacheInterfaceMockRecorder) Flush() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockPositionCacheInterface)(nil).Flush))
}

// Get mocks base method
func (m *MockPositionCacheInterface) Get() (position_store.Position, bool, error) {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(position_store.Position)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockPositionCacheInterfaceMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPositionCacheInterface)(nil).Get))
}

// Put mocks base method
func (m *MockPositionCacheInterface) Put(arg0 position_store.Position) error {
	ret := m.ctrl.Call(m, "Put", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockPositionCacheInterfaceMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPositionCacheInterface)(nil).Put), arg0)
}

// Start mocks base method
func (m *MockPositionCacheInterface) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockPositionCacheInterfaceMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPositionCacheInterface)(nil).Start))
}
