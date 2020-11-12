// Code generated by MockGen. DO NOT EDIT.
// Source: ScheduleUsecase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTimeTableUseCase is a mock of TimeTableUseCase interface
type MockTimeTableUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockTimeTableUseCaseMockRecorder
}

// MockTimeTableUseCaseMockRecorder is the mock recorder for MockTimeTableUseCase
type MockTimeTableUseCaseMockRecorder struct {
	mock *MockTimeTableUseCase
}

// NewMockTimeTableUseCase creates a new mock instance
func NewMockTimeTableUseCase(ctrl *gomock.Controller) *MockTimeTableUseCase {
	mock := &MockTimeTableUseCase{ctrl: ctrl}
	mock.recorder = &MockTimeTableUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimeTableUseCase) EXPECT() *MockTimeTableUseCaseMockRecorder {
	return m.recorder
}

// GetMovieSchedule mocks base method
func (m *MockTimeTableUseCase) GetMovieSchedule(MovieID, CinemaID, date string) (*[]models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieSchedule", MovieID, CinemaID, date)
	ret0, _ := ret[0].(*[]models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieSchedule indicates an expected call of GetMovieSchedule
func (mr *MockTimeTableUseCaseMockRecorder) GetMovieSchedule(MovieID, CinemaID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieSchedule", reflect.TypeOf((*MockTimeTableUseCase)(nil).GetMovieSchedule), MovieID, CinemaID, date)
}

// GetSchedule mocks base method
func (m *MockTimeTableUseCase) GetSchedule(ScheduleID string) (*models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", ScheduleID)
	ret0, _ := ret[0].(*models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockTimeTableUseCaseMockRecorder) GetSchedule(ScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockTimeTableUseCase)(nil).GetSchedule), ScheduleID)
}
