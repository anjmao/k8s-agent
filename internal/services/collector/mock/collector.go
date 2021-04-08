// Code generated by MockGen. DO NOT EDIT.
// Source: castai-agent/internal/services/collector (interfaces: Collector)

// Package mock_collector is a generated GoMock package.
package mock_collector

import (
	collector "castai-agent/internal/services/collector"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	version "k8s.io/apimachinery/pkg/version"
)

// MockCollector is a mock of Collector interface.
type MockCollector struct {
	ctrl     *gomock.Controller
	recorder *MockCollectorMockRecorder
}

// MockCollectorMockRecorder is the mock recorder for MockCollector.
type MockCollectorMockRecorder struct {
	mock *MockCollector
}

// NewMockCollector creates a new mock instance.
func NewMockCollector(ctrl *gomock.Controller) *MockCollector {
	mock := &MockCollector{ctrl: ctrl}
	mock.recorder = &MockCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollector) EXPECT() *MockCollectorMockRecorder {
	return m.recorder
}

// Collect mocks base method.
func (m *MockCollector) Collect(arg0 context.Context) (*collector.ClusterData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", arg0)
	ret0, _ := ret[0].(*collector.ClusterData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect.
func (mr *MockCollectorMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockCollector)(nil).Collect), arg0)
}

// GetVersion mocks base method.
func (m *MockCollector) GetVersion() *version.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(*version.Info)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockCollectorMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockCollector)(nil).GetVersion))
}
