// Code generated by MockGen. DO NOT EDIT.
// Source: Repository.go

// Package mock is a generated GoMock package.
package mock

import (
	mapset "github.com/deckarep/golang-set"
	models "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
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

// GetMovieRatingsDataset mocks base method
func (m *MockRepository) GetMovieRatingsDataset() (*[]models.RecommendationDataFrame, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieRatingsDataset")
	ret0, _ := ret[0].(*[]models.RecommendationDataFrame)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieRatingsDataset indicates an expected call of GetMovieRatingsDataset
func (mr *MockRepositoryMockRecorder) GetMovieRatingsDataset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieRatingsDataset", reflect.TypeOf((*MockRepository)(nil).GetMovieRatingsDataset))
}

// GetRecommendedMovieList mocks base method
func (m *MockRepository) GetRecommendedMovieList(set *mapset.Set) (*[]models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendedMovieList", set)
	ret0, _ := ret[0].(*[]models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendedMovieList indicates an expected call of GetRecommendedMovieList
func (mr *MockRepositoryMockRecorder) GetRecommendedMovieList(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendedMovieList", reflect.TypeOf((*MockRepository)(nil).GetRecommendedMovieList), set)
}

// GetPopularMovies mocks base method
func (m *MockRepository) GetPopularMovies() (*[]models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularMovies")
	ret0, _ := ret[0].(*[]models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularMovies indicates an expected call of GetPopularMovies
func (mr *MockRepositoryMockRecorder) GetPopularMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularMovies", reflect.TypeOf((*MockRepository)(nil).GetPopularMovies))
}
