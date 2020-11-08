// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	models "backend/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CheckAvailability mocks base method
func (m *MockRepository) CheckAvailability(hallID uint64, place *models.TicketPlace) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAvailability", hallID, place)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAvailability indicates an expected call of CheckAvailability
func (mr *MockRepositoryMockRecorder) CheckAvailability(hallID, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAvailability", reflect.TypeOf((*MockRepository)(nil).CheckAvailability), hallID, place)
}

// GetHallStructure mocks base method
func (m *MockRepository) GetHallStructure(hallID uint64) (*models.CinemaHall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHallStructure", hallID)
	ret0, _ := ret[0].(*models.CinemaHall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHallStructure indicates an expected call of GetHallStructure
func (mr *MockRepositoryMockRecorder) GetHallStructure(hallID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHallStructure", reflect.TypeOf((*MockRepository)(nil).GetHallStructure), hallID)
}
