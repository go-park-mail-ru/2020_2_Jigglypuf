// Code generated by MockGen. DO NOT EDIT.
// Source: AuthUsecase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockUserUseCase is a mock of UserUseCase interface
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// SignUp mocks base method
func (m *MockUserUseCase) SignUp(input *models.RegistrationInput) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", input)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp
func (mr *MockUserUseCaseMockRecorder) SignUp(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserUseCase)(nil).SignUp), input)
}

// SignIn mocks base method
func (m *MockUserUseCase) SignIn(input *models.AuthInput) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", input)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn
func (mr *MockUserUseCaseMockRecorder) SignIn(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserUseCase)(nil).SignIn), input)
}

// SignOut mocks base method
func (m *MockUserUseCase) SignOut(cookie *http.Cookie) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOut", cookie)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignOut indicates an expected call of SignOut
func (mr *MockUserUseCaseMockRecorder) SignOut(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOut", reflect.TypeOf((*MockUserUseCase)(nil).SignOut), cookie)
}
