// Code generated by MockGen. DO NOT EDIT.
// Source: gateway.go

// Package country is a generated GoMock package.
package country

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCountryGateway is a mock of CountryGateway interface.
type MockCountryGateway struct {
	ctrl     *gomock.Controller
	recorder *MockCountryGatewayMockRecorder
}

// MockCountryGatewayMockRecorder is the mock recorder for MockCountryGateway.
type MockCountryGatewayMockRecorder struct {
	mock *MockCountryGateway
}

// NewMockCountryGateway creates a new mock instance.
func NewMockCountryGateway(ctrl *gomock.Controller) *MockCountryGateway {
	mock := &MockCountryGateway{ctrl: ctrl}
	mock.recorder = &MockCountryGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountryGateway) EXPECT() *MockCountryGatewayMockRecorder {
	return m.recorder
}

// CountIPsByCountryCode mocks base method.
func (m *MockCountryGateway) CountIPsByCountryCode(ctx context.Context, countryCode string) (IPCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountIPsByCountryCode", ctx, countryCode)
	ret0, _ := ret[0].(IPCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountIPsByCountryCode indicates an expected call of CountIPsByCountryCode.
func (mr *MockCountryGatewayMockRecorder) CountIPsByCountryCode(ctx, countryCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountIPsByCountryCode", reflect.TypeOf((*MockCountryGateway)(nil).CountIPsByCountryCode), ctx, countryCode)
}

// GetIPInfo mocks base method.
func (m *MockCountryGateway) GetIPInfo(ctx context.Context, ip string) (IPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPInfo", ctx, ip)
	ret0, _ := ret[0].(IPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPInfo indicates an expected call of GetIPInfo.
func (mr *MockCountryGatewayMockRecorder) GetIPInfo(ctx, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPInfo", reflect.TypeOf((*MockCountryGateway)(nil).GetIPInfo), ctx, ip)
}

// TopTenISPByCountryCode mocks base method.
func (m *MockCountryGateway) TopTenISPByCountryCode(ctx context.Context, countryCode string) ([]ISPCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopTenISPByCountryCode", ctx, countryCode)
	ret0, _ := ret[0].([]ISPCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopTenISPByCountryCode indicates an expected call of TopTenISPByCountryCode.
func (mr *MockCountryGatewayMockRecorder) TopTenISPByCountryCode(ctx, countryCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopTenISPByCountryCode", reflect.TypeOf((*MockCountryGateway)(nil).TopTenISPByCountryCode), ctx, countryCode)
}
