// Code generated by MockGen. DO NOT EDIT.
// Source: viodata.go

// Package viodata is a generated GoMock package.
package viodata

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVioData is a mock of VioData interface.
type MockVioData struct {
	ctrl     *gomock.Controller
	recorder *MockVioDataMockRecorder
}

// MockVioDataMockRecorder is the mock recorder for MockVioData.
type MockVioDataMockRecorder struct {
	mock *MockVioData
}

// NewMockVioData creates a new mock instance.
func NewMockVioData(ctrl *gomock.Controller) *MockVioData {
	mock := &MockVioData{ctrl: ctrl}
	mock.recorder = &MockVioDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVioData) EXPECT() *MockVioDataMockRecorder {
	return m.recorder
}

// GetIPLocationByIP mocks base method.
func (m *MockVioData) GetIPLocationByIP(ctx context.Context, IP string) (*IPLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPLocationByIP", ctx, IP)
	ret0, _ := ret[0].(*IPLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPLocationByIP indicates an expected call of GetIPLocationByIP.
func (mr *MockVioDataMockRecorder) GetIPLocationByIP(ctx, IP interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPLocationByIP", reflect.TypeOf((*MockVioData)(nil).GetIPLocationByIP), ctx, IP)
}
