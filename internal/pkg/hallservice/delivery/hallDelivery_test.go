package delivery

import (
	"backend/internal/pkg/hallservice"
	"backend/internal/pkg/hallservice/mock"
	"backend/internal/pkg/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestingHallStruct struct{
	handler *HallDelivery
	useCaseMock *mock.MockUseCase
	GoMockController *gomock.Controller
}

var(
	TestingStruct *TestingHallStruct = nil
)

func setUp(t *testing.T){
	TestingStruct = new(TestingHallStruct)
	TestingStruct.GoMockController = gomock.NewController(t)

	TestingStruct.useCaseMock = mock.NewMockUseCase(TestingStruct.GoMockController)
	TestingStruct.handler = NewHallDelivery(TestingStruct.useCaseMock)
}

func tearDown(){
	TestingStruct.GoMockController.Finish()
}

func TestGetHallStructureSuccessCase(t *testing.T){
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet, "/hall/1/", nil)
	testRecorder := httptest.NewRecorder()

	hallItem := new(models.CinemaHall)
	varsMap := map[string]string{
		hallservice.HallIDPathName: "1",
	}
	TestingStruct.useCaseMock.EXPECT().GetHallStructure(gomock.Any()).Return(hallItem, nil)
	testReq = mux.SetURLVars(testReq, varsMap)

	TestingStruct.handler.GetHallStructure(testRecorder, testReq)

	if testRecorder.Code != http.StatusOK{
		t.Fatalf("TEST: Success get hall "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}

	tearDown()
}

func TestGetHallStructureFailureCases(t *testing.T){
	setUp(t)

	testCases := []struct{
		Request *http.Request
		Recorder *httptest.ResponseRecorder
		StatusCode int
	}{
		{
			httptest.NewRequest(http.MethodPost, "/hall/", nil),
			httptest.NewRecorder(),
			405,
		},
		{
			httptest.NewRequest(http.MethodGet, "/hall/", nil),
			httptest.NewRecorder(),
			400,
		},
	}

	for _, val := range testCases{
		TestingStruct.handler.GetHallStructure(val.Recorder, val.Request)
		if val.Recorder.Code != val.StatusCode{
			t.Fatalf("TEST: Failure get hall "+
				"handler returned wrong status code: got %v want %v", val.Recorder.Code, val.StatusCode)
		}
	}

	tearDown()
}

func TestGetHallUCErrorHandling(t *testing.T){
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet, "/hall/", nil)
	testRecorder := httptest.NewRecorder()

	varsMap := map[string]string{
		hallservice.HallIDPathName: "1",
	}
	TestingStruct.useCaseMock.EXPECT().GetHallStructure(gomock.Any()).Return(nil, errors.New("test error"))
	testReq = mux.SetURLVars(testReq, varsMap)

	TestingStruct.handler.GetHallStructure(testRecorder, testReq)
	if testRecorder.Code != http.StatusBadRequest{
		t.Fatalf("TEST: Failure UC get hall "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}
	tearDown()
}