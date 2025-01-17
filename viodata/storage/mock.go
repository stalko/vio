// Code generated by MockGen. DO NOT EDIT.
// Source: storage/storage.go

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// BulkInsertIPLocation mocks base method.
func (m *MockStorage) BulkInsertIPLocation(ctx context.Context, IPLocations []InsertIPLocation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkInsertIPLocation", ctx, IPLocations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkInsertIPLocation indicates an expected call of BulkInsertIPLocation.
func (mr *MockStorageMockRecorder) BulkInsertIPLocation(ctx, IPLocations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInsertIPLocation", reflect.TypeOf((*MockStorage)(nil).BulkInsertIPLocation), ctx, IPLocations)
}

// GetIPLocationsByIPAddress mocks base method.
func (m *MockStorage) GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (*IPLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPLocationsByIPAddress", ctx, ipAddress)
	ret0, _ := ret[0].(*IPLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPLocationsByIPAddress indicates an expected call of GetIPLocationsByIPAddress.
func (mr *MockStorageMockRecorder) GetIPLocationsByIPAddress(ctx, ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPLocationsByIPAddress", reflect.TypeOf((*MockStorage)(nil).GetIPLocationsByIPAddress), ctx, ipAddress)
}
