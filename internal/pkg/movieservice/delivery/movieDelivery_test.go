package delivery

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/movieservice/mock"
	cookieService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestingMovieStruct struct {
	handler          *MovieHandler
	useCaseMock      *mock.MockMovieUseCase
	GoMockController *gomock.Controller
}

var (
	TestingStruct *TestingMovieStruct = nil
)

func setUp(t *testing.T) {
	TestingStruct = new(TestingMovieStruct)

	TestingStruct.GoMockController = gomock.NewController(t)
	TestingStruct.useCaseMock = mock.NewMockMovieUseCase(TestingStruct.GoMockController)
	TestingStruct.handler = NewMovieHandler(TestingStruct.useCaseMock)
}

func tearDown() {
	TestingStruct.GoMockController.Finish()
}

func TestGetMovieListSuccessCase(t *testing.T) {
	setUp(t)

	inputArray := []models.MovieList{
		{Name: "First"},
		{Name: "Second"},
	}
	testReq := httptest.NewRequest(http.MethodGet, "/cinema/?limit=10&page=1", nil)
	testRecorder := httptest.NewRecorder()

	TestingStruct.useCaseMock.EXPECT().GetMovieList(gomock.Any(), gomock.Any()).Return(&inputArray, nil)
	TestingStruct.handler.GetMovieList(testRecorder, testReq)
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Get cinema list success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	resultArray := new([]models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), resultArray)
	if scanErr != nil {
		t.Fatalf("TEST: Get cinema list success")
	}

	tearDown()
}

func TestGetMovieSuccessCase(t *testing.T) {
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/movie/1/", nil)
	testReq = mux.SetURLVars(testReq, map[string]string{
		movieservice.GetMovieID: "1",
	})
	testRecorder := httptest.NewRecorder()
	resultItem := models.Movie{
		Name: "somemovie",
	}

	TestingStruct.useCaseMock.EXPECT().GetMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(&resultItem, nil)
	TestingStruct.handler.GetMovie(testRecorder, testReq)
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Get cinema success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	responseItem := new(models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), responseItem)
	if scanErr != nil {
		t.Fatalf("TEST: Get cinema success")
	}

	tearDown()
}

func TestGetMoviesInCinema(t *testing.T) {
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
	TestingStruct.useCaseMock.EXPECT().GetActualMovies(gomock.Any(), gomock.Any(), gomock.Any()).Return(&inputArr, nil)
	TestingStruct.handler.GetActualMovies(testRecorder, testReq)

	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Get cinema success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	responseItem := new([]models.Movie)
	scanErr := json.Unmarshal(testRecorder.Body.Bytes(), responseItem)
	if scanErr != nil {
		t.Fatalf("TEST: Get cinema in movie success")
	}
	tearDown()
}

func TestRateMovieSuccessCase(t *testing.T) {
	setUp(t)
	ratingModel := new(models.Movie)
	ratingBody, _ := json.Marshal(ratingModel)
	testReq := httptest.NewRequest(http.MethodPost, "/movie/rate", strings.NewReader(string(ratingBody)))
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, cookieService.ContextUserIDName, uint64(1))
	ctx = context.WithValue(ctx, cookieService.ContextIsAuthName, true)
	testRecorder := httptest.NewRecorder()
	TestingStruct.useCaseMock.EXPECT().RateMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	TestingStruct.handler.RateMovie(testRecorder, testReq.WithContext(ctx))

	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Rate movie success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}

	tearDown()
}

func TestHandlerOnInvalidMethod(t *testing.T) {
	setUp(t)
	var testCases = []struct {
		HandlerFunc func(w http.ResponseWriter, r *http.Request)
		Method      string
	}{
		{
			TestingStruct.handler.RateMovie,
			http.MethodGet,
		},
		{
			TestingStruct.handler.GetActualMovies,
			http.MethodPost,
		},
		{
			TestingStruct.handler.GetMovie,
			http.MethodPost,
		},
		{
			TestingStruct.handler.GetMovieList,
			http.MethodPost,
		},
	}

	for _, val := range testCases {
		testReq := httptest.NewRequest(val.Method, "/movie/", nil)
		testRecorder := httptest.NewRecorder()
		val.HandlerFunc(testRecorder, testReq)
		if testRecorder.Code != http.StatusMethodNotAllowed {
			t.Fatalf("TEST: Invalid method movie "+
				"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
		}
	}

	tearDown()
}

func TestIncorrectGetParametersMovie(t *testing.T) {
	setUp(t)

	var testCases = []struct {
		HandleFunc func(w http.ResponseWriter, r *http.Request)
		URL        string
	}{
		{
			TestingStruct.handler.GetMovieList,
			"/cinema/?limit=asdas&page=qwewqe",
		},
		{
			TestingStruct.handler.GetMovieList,
			"/cinema/",
		},
		{
			TestingStruct.handler.GetActualMovies,
			"/cinema/?limit=asdas&page=qwewqe",
		},
		{
			TestingStruct.handler.GetActualMovies,
			"/cinema/",
		},
	}

	for _, val := range testCases {
		testReq := httptest.NewRequest(http.MethodGet, val.URL, nil)
		testRecorder := httptest.NewRecorder()
		val.HandleFunc(testRecorder, testReq)

		if testRecorder.Code != http.StatusBadRequest {
			t.Fatalf("TEST: Invalid get parameters movie "+
				"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
		}
	}

	tearDown()
}

func TestRateMovieUnAuthCase(t *testing.T) {
	setUp(t)
	ratingModel := new(models.Movie)
	ratingBody, _ := json.Marshal(ratingModel)
	testReq := httptest.NewRequest(http.MethodPost, "/movie/rate", strings.NewReader(string(ratingBody)))
	testRecorder := httptest.NewRecorder()
	TestingStruct.handler.RateMovie(testRecorder, testReq)

	if testRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("TEST: Rate movie success "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusUnauthorized)
	}

	tearDown()
}
