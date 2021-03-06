// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tchorzewski1991/flusher/flusher (interfaces: Exporter)

// Package flusherMock is a generated GoMock package.
package flusherMock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExporter is a mock of Exporter interface.
type MockExporter struct {
	ctrl     *gomock.Controller
	recorder *MockExporterMockRecorder
}

// MockExporterMockRecorder is the mock recorder for MockExporter.
type MockExporterMockRecorder struct {
	mock *MockExporter
}

// NewMockExporter creates a new mock instance.
func NewMockExporter(ctrl *gomock.Controller) *MockExporter {
	mock := &MockExporter{ctrl: ctrl}
	mock.recorder = &MockExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExporter) EXPECT() *MockExporterMockRecorder {
	return m.recorder
}

// Export mocks base method.
func (m *MockExporter) Export() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export")
	ret0, _ := ret[0].(error)
	return ret0
}

// Export indicates an expected call of Export.
func (mr *MockExporterMockRecorder) Export() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockExporter)(nil).Export))
}
