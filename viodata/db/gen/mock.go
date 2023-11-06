// Code generated by MockGen. DO NOT EDIT.
// Source: db/gen/querier.go

// Package gen is a generated GoMock package.
package gen

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// BulkInsertIPLocations mocks base method.
func (m *MockQuerier) BulkInsertIPLocations(ctx context.Context, arg []BulkInsertIPLocationsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkInsertIPLocations", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkInsertIPLocations indicates an expected call of BulkInsertIPLocations.
func (mr *MockQuerierMockRecorder) BulkInsertIPLocations(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInsertIPLocations", reflect.TypeOf((*MockQuerier)(nil).BulkInsertIPLocations), ctx, arg)
}

// GetCountIPLocationsByIPAddress mocks base method.
func (m *MockQuerier) GetCountIPLocationsByIPAddress(ctx context.Context, ipAddress string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountIPLocationsByIPAddress", ctx, ipAddress)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountIPLocationsByIPAddress indicates an expected call of GetCountIPLocationsByIPAddress.
func (mr *MockQuerierMockRecorder) GetCountIPLocationsByIPAddress(ctx, ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountIPLocationsByIPAddress", reflect.TypeOf((*MockQuerier)(nil).GetCountIPLocationsByIPAddress), ctx, ipAddress)
}

// GetCountryByID mocks base method.
func (m *MockQuerier) GetCountryByID(ctx context.Context, id string) (Country, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountryByID", ctx, id)
	ret0, _ := ret[0].(Country)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountryByID indicates an expected call of GetCountryByID.
func (mr *MockQuerierMockRecorder) GetCountryByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountryByID", reflect.TypeOf((*MockQuerier)(nil).GetCountryByID), ctx, id)
}

// GetIPLocationsByIPAddress mocks base method.
func (m *MockQuerier) GetIPLocationsByIPAddress(ctx context.Context, ipAddress string) (GetIPLocationsByIPAddressRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPLocationsByIPAddress", ctx, ipAddress)
	ret0, _ := ret[0].(GetIPLocationsByIPAddressRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPLocationsByIPAddress indicates an expected call of GetIPLocationsByIPAddress.
func (mr *MockQuerierMockRecorder) GetIPLocationsByIPAddress(ctx, ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPLocationsByIPAddress", reflect.TypeOf((*MockQuerier)(nil).GetIPLocationsByIPAddress), ctx, ipAddress)
}

// InsertCountry mocks base method.
func (m *MockQuerier) InsertCountry(ctx context.Context, arg InsertCountryParams) (InsertCountryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCountry", ctx, arg)
	ret0, _ := ret[0].(InsertCountryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCountry indicates an expected call of InsertCountry.
func (mr *MockQuerierMockRecorder) InsertCountry(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCountry", reflect.TypeOf((*MockQuerier)(nil).InsertCountry), ctx, arg)
}

// InsertIPLocationWIP mocks base method.
func (m *MockQuerier) InsertIPLocationWIP(ctx context.Context, arg InsertIPLocationWIPParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertIPLocationWIP", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertIPLocationWIP indicates an expected call of InsertIPLocationWIP.
func (mr *MockQuerierMockRecorder) InsertIPLocationWIP(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertIPLocationWIP", reflect.TypeOf((*MockQuerier)(nil).InsertIPLocationWIP), ctx, arg)
}