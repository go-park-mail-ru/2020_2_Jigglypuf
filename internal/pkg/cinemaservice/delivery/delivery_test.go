package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/cinemaservice/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestingCinemaStruct struct {
	Handler *CinemaHandler
	UseCaseMock *mock.MockUseCase
	GoMockController *gomock.Controller
}

var(
	testingStruct *TestingCinemaStruct = nil
)

func setUp(t *testing.T){
	testingStruct = new(TestingCinemaStruct)
	testingStruct.GoMockController = gomock.NewController(t)
	CinemaUCMock := mock.NewMockUseCase(testingStruct.GoMockController)
	MainHandler := NewCinemaHandler(CinemaUCMock)
	testingStruct.Handler = MainHandler
	testingStruct.UseCaseMock = CinemaUCMock
}

func tearDown(){
	testingStruct.GoMockController.Finish()
}


func TestGetCinemaSuccessCase(t *testing.T){
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet,"/cinema/1",nil)
	testRecorder := httptest.NewRecorder()
	testParams := httprouter.Params{
		httprouter.Param{Key: cinemaservice.CinemaIDParam, Value: "1"},
	}
	testCinema := models.Cinema{
		Name: "testCinema",
	}
	testingStruct.UseCaseMock.EXPECT().GetCinema(gomock.Any()).Return(&testCinema, nil)
	testingStruct.Handler.GetCinema(testRecorder, testReq, testParams)
	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: Get single cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	resultCinema := new(models.Cinema)
	testBodyDecoder := json.NewDecoder(testRecorder.Body)
	decodeErr := testBodyDecoder.Decode(resultCinema)
	if decodeErr != nil{
		t.Fatalf("TEST: Get single cinema "+
			"return cinema model decode err \n")
	}
	tearDown()
}

func TestGetCinemaIncorrectMethod(t *testing.T){
	setUp(t)
	testReq := httptest.NewRequest(http.MethodPost, "/cinema/1", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.GetCinema(testRecorder,testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusMethodNotAllowed{
		t.Fatalf("TEST: Incorrect method in get cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestGetCinemaIncorrectParam(t *testing.T){
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet,"/cinema/1", nil)
	testRecorder := httptest.NewRecorder()
	testParams := httprouter.Params{
		httprouter.Param{Key: cinemaservice.CinemaIDParam, Value: "incorrectParam"},
	}
	testingStruct.Handler.GetCinema(testRecorder, testReq, testParams)
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: incorrect param single cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}
	tearDown()
}

func TestGetCinemaHandlingUCErr(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cinema/1", nil)
	testRecorder := httptest.NewRecorder()
	testParams := httprouter.Params{
		httprouter.Param{Key: cinemaservice.CinemaIDParam, Value: "1"},
	}

	testingStruct.UseCaseMock.EXPECT().GetCinema(gomock.Any()).Return(nil,errors.New("test error"))
	testingStruct.Handler.GetCinema(testRecorder, testReq, testParams)
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: handling uc error single cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}
	tearDown()
}

func TestGetCinemaListSuccessCase(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cinema/?limit=10&page=1", nil)
	testRecorder := httptest.NewRecorder()

	returnArray := []models.Cinema{
		models.Cinema{Name: "first"},
		models.Cinema{Name: "Second"},
	}
	testingStruct.UseCaseMock.EXPECT().GetCinemaList(gomock.Any(), gomock.Any()).Return(&returnArray, nil)
	testingStruct.Handler.GetCinemaList(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: success list cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	var resultList = new([]models.Cinema)
	decodeErr := json.Unmarshal(testRecorder.Body.Bytes(), resultList)
	if decodeErr != nil{
		t.Fatalf("TEST: Get list cinema "+
			"return cinema model list decode err \n")
	}
	if !reflect.DeepEqual(*resultList, returnArray){
		t.Fatalf("TEST: Get list cinema arrays not equal")
	}

	tearDown()
}

func TestGetCinemaListIncorrectMethod(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodPost, "/cinema/", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.GetCinemaList(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusMethodNotAllowed{
		t.Fatalf("TEST: incorrect method handle " +
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	tearDown()
}

func TestGetCinemaListIncorrectQueryParams(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cinema/?limit=abc&page=ebc", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.GetCinemaList(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: incorrect query handle " +
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}

	tearDown()
}

func TestGetCinemaListNoQueryParams(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cinema/", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.GetCinemaList(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: No query handle " +
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}

	tearDown()
}

func TestGetCinemaListUCErrorHandle(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cinema/?limit=10&page=1", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.UseCaseMock.EXPECT().GetCinemaList(gomock.Any(), gomock.Any()).Return(nil,errors.New("test err"))
	testingStruct.Handler.GetCinemaList(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: success list cinema "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}

	tearDown()
}