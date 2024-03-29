// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/weBEE9/opt-auth-backend/service (interfaces: OTPService)
//
// Generated by this command:
//
//	mockgen -destination=../mock/service_mock.go -package=mock github.com/weBEE9/opt-auth-backend/service OTPService
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockOTPService is a mock of OTPService interface.
type MockOTPService struct {
	ctrl     *gomock.Controller
	recorder *MockOTPServiceMockRecorder
}

// MockOTPServiceMockRecorder is the mock recorder for MockOTPService.
type MockOTPServiceMockRecorder struct {
	mock *MockOTPService
}

// NewMockOTPService creates a new mock instance.
func NewMockOTPService(ctrl *gomock.Controller) *MockOTPService {
	mock := &MockOTPService{ctrl: ctrl}
	mock.recorder = &MockOTPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOTPService) EXPECT() *MockOTPServiceMockRecorder {
	return m.recorder
}

// GenOTP mocks base method.
func (m *MockOTPService) GenOTP(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenOTP", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenOTP indicates an expected call of GenOTP.
func (mr *MockOTPServiceMockRecorder) GenOTP(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenOTP", reflect.TypeOf((*MockOTPService)(nil).GenOTP), arg0, arg1)
}

// TTL mocks base method.
func (m *MockOTPService) TTL(arg0 context.Context, arg1 string) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", arg0, arg1)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TTL indicates an expected call of TTL.
func (mr *MockOTPServiceMockRecorder) TTL(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockOTPService)(nil).TTL), arg0, arg1)
}

// VerifyOTP mocks base method.
func (m *MockOTPService) VerifyOTP(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOTP", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOTP indicates an expected call of VerifyOTP.
func (mr *MockOTPServiceMockRecorder) VerifyOTP(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOTP", reflect.TypeOf((*MockOTPService)(nil).VerifyOTP), arg0, arg1, arg2)
}
