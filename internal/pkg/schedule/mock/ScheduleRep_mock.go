// Code generated by MockGen. DO NOT EDIT.
// Source: ScheduleRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTimeTableRepository is a mock of TimeTableRepository interface
type MockTimeTableRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTimeTableRepositoryMockRecorder
}

// MockTimeTableRepositoryMockRecorder is the mock recorder for MockTimeTableRepository
type MockTimeTableRepositoryMockRecorder struct {
	mock *MockTimeTableRepository
}

// NewMockTimeTableRepository creates a new mock instance
func NewMockTimeTableRepository(ctrl *gomock.Controller) *MockTimeTableRepository {
	mock := &MockTimeTableRepository{ctrl: ctrl}
	mock.recorder = &MockTimeTableRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimeTableRepository) EXPECT() *MockTimeTableRepositoryMockRecorder {
	return m.recorder
}

// GetMovieSchedule mocks base method
func (m *MockTimeTableRepository) GetMovieSchedule(MovieID uint64, date string) (*[]models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieSchedule", MovieID, date)
	ret0, _ := ret[0].(*[]models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieSchedule indicates an expected call of GetMovieSchedule
func (mr *MockTimeTableRepositoryMockRecorder) GetMovieSchedule(MovieID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieSchedule", reflect.TypeOf((*MockTimeTableRepository)(nil).GetMovieSchedule), MovieID, date)
}

// GetMovieCinemaSchedule mocks base method
func (m *MockTimeTableRepository) GetMovieCinemaSchedule(MovieID, CinemaID uint64, date string) (*[]models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieCinemaSchedule", MovieID, CinemaID, date)
	ret0, _ := ret[0].(*[]models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieCinemaSchedule indicates an expected call of GetMovieCinemaSchedule
func (mr *MockTimeTableRepositoryMockRecorder) GetMovieCinemaSchedule(MovieID, CinemaID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieCinemaSchedule", reflect.TypeOf((*MockTimeTableRepository)(nil).GetMovieCinemaSchedule), MovieID, CinemaID, date)
}

// GetScheduleHallID mocks base method
func (m *MockTimeTableRepository) GetScheduleHallID(scheduleID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduleHallID", scheduleID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleHallID indicates an expected call of GetScheduleHallID
func (mr *MockTimeTableRepositoryMockRecorder) GetScheduleHallID(scheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleHallID", reflect.TypeOf((*MockTimeTableRepository)(nil).GetScheduleHallID), scheduleID)
}
