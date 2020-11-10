package delivery

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/movieservice"
	"backend/internal/pkg/movieservice/mock"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestingMovieStruct struct{
	handler *MovieHandler
	useCaseMock *mock.MockMovieUseCase
	GoMockController *gomock.Controller
}

var(
	TestingStruct *TestingMovieStruct = nil
)

func setUp(t *testing.T){
	TestingStruct = new(TestingMovieStruct)

	TestingStruct.GoMockController = gomock.NewController(t)
	TestingStruct.useCaseMock = mock.NewMockMovieUseCase(TestingStruct.GoMockController)
	TestingStruct.handler = NewMovieHandler(TestingStruct.useCaseMock)
}

func tearDown(){
	TestingStruct.GoMockController.Finish()
}


func TestGetMovieListSuccessCase(t *testing.T){
	setUp(t)

	inputArray := []models.MovieList{
		models.MovieList{Name: "First"},
		models.MovieList{Name: "Second"},
	}
	testReq := httptest.NewRequest(http.MethodGet, "/cinema/?limit=10&page=1", nil)
	testRecorder := httptest.NewRecorder()

	TestingStruct.useCaseMock.EXPECT().GetMovieList(gomock.Any(), gomock.Any()).Return(&inputArray, nil)
	TestingStruct.handler.GetMovieList(testRecorder, testReq)
	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: Get cinema list success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	resultArray := new([]models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), resultArray)
	if scanErr != nil{
		t.Fatalf("TEST: Get cinema list success")
	}

	tearDown()
}

func TestGetMovieSuccessCase(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/movie/1/", nil)
	testReq = mux.SetURLVars(testReq, map[string]string{
		movieservice.GetMovieID:"1",
	})
	testRecorder := httptest.NewRecorder()
	resultItem := models.Movie{
		Name: "somemovie",
	}

	TestingStruct.useCaseMock.EXPECT().GetMovie(gomock.Any(), gomock.Any(),gomock.Any()).Return(&resultItem, nil)
	TestingStruct.handler.GetMovie(testRecorder, testReq)
	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: Get cinema success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	responseItem := new(models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), responseItem)
	if scanErr != nil{
		t.Fatalf("TEST: Get cinema success")
	}

	tearDown()
}

func TestGetMoviesInCinema(t *testing.T){
	setUp(t)

	inputArr := []models.MovieList{
		{
			Name: "NoMake",
		},
		{
			Name: "Second",
		},
	}
	testReq := httptest.NewRequest(http.MethodGet, "/movie/actual/?limit=10&page=1", nil)
	testRecorder := httptest.NewRecorder()
	TestingStruct.useCaseMock.EXPECT().GetMoviesInCinema(gomock.Any(), gomock.Any()).Return(&inputArr, nil)
	TestingStruct.handler.GetMoviesInCinema(testRecorder, testReq)

	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: Get cinema success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	responseItem := new([]models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), responseItem)
	if scanErr != nil{
		t.Fatalf("TEST: Get cinema in movie success")
	}
	tearDown()
}

//func TestRateMovieSuccessCase(t *testing.T){
//	setUp(t)
//	ratingModel := i
//	testReq := httptest.NewRequest(http.MethodPost, "/movie/rate", nil)
//
//
//
//	tearDown()
//}