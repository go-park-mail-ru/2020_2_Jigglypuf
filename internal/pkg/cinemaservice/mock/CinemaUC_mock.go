// Code generated by MockGen. DO NOT EDIT.
// Source: CinemaUseCase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "backend/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateCinema mocks base method
func (m *MockUseCase) CreateCinema(arg0 *models.Cinema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCinema", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCinema indicates an expected call of CreateCinema
func (mr *MockUseCaseMockRecorder) CreateCinema(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCinema", reflect.TypeOf((*MockUseCase)(nil).CreateCinema), arg0)
}

// GetCinema mocks base method
func (m *MockUseCase) GetCinema(id uint64) (*models.Cinema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCinema", id)
	ret0, _ := ret[0].(*models.Cinema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCinema indicates an expected call of GetCinema
func (mr *MockUseCaseMockRecorder) GetCinema(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCinema", reflect.TypeOf((*MockUseCase)(nil).GetCinema), id)
}

// GetCinemaList mocks base method
func (m *MockUseCase) GetCinemaList(limit, page int) (*[]models.Cinema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCinemaList", limit, page)
	ret0, _ := ret[0].(*[]models.Cinema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCinemaList indicates an expected call of GetCinemaList
func (mr *MockUseCaseMockRecorder) GetCinemaList(limit, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCinemaList", reflect.TypeOf((*MockUseCase)(nil).GetCinemaList), limit, page)
}
