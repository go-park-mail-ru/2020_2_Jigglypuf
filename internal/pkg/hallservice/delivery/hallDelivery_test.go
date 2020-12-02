package delivery

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/hallservice/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestingHallStruct struct {
	handler          *HallDelivery
	useCaseMock      *mock.MockUseCase
	GoMockController *gomock.Controller
}

var (
	testingStruct *TestingHallStruct = nil
)

func setUp(t *testing.T) {
	testingStruct = new(TestingHallStruct)
	testingStruct.GoMockController = gomock.NewController(t)

	testingStruct.useCaseMock = mock.NewMockUseCase(testingStruct.GoMockController)
	testingStruct.handler = NewHallDelivery(testingStruct.useCaseMock)
}

func tearDown() {
	testingStruct.GoMockController.Finish()
}

func TestGetHallStructureSuccessCase(t *testing.T) {
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet, "/hall/1/", nil)
	testRecorder := httptest.NewRecorder()

	hallItem := new(models.CinemaHall)
	varsMap := map[string]string{
		hallservice.HallIDPathName: "1",
	}
	testingStruct.useCaseMock.EXPECT().GetHallStructure(gomock.Any()).Return(hallItem, nil)
	testReq = mux.SetURLVars(testReq, varsMap)

	testingStruct.handler.GetHallStructure(testRecorder, testReq)

	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Success get hall "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}

	tearDown()
}

func TestGetHallStructureFailureCases(t *testing.T) {
	setUp(t)

	testCases := []struct {
		Request    *http.Request
		Recorder   *httptest.ResponseRecorder
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

	for _, val := range testCases {
		testingStruct.handler.GetHallStructure(val.Recorder, val.Request)
		if val.Recorder.Code != val.StatusCode {
			t.Fatalf("TEST: Failure get hall "+
				"handler returned wrong status code: got %v want %v", val.Recorder.Code, val.StatusCode)
		}
	}

	tearDown()
}

func TestGetHallUCErrorHandling(t *testing.T) {
	setUp(t)
	testReq := httptest.NewRequest(http.MethodGet, "/hall/", nil)
	testRecorder := httptest.NewRecorder()

	varsMap := map[string]string{
		hallservice.HallIDPathName: "1",
	}
	testingStruct.useCaseMock.EXPECT().GetHallStructure(gomock.Any()).Return(nil, errors.New("test error"))
	testReq = mux.SetURLVars(testReq, varsMap)

	testingStruct.handler.GetHallStructure(testRecorder, testReq)
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: Failure UC get hall "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}
	tearDown()
}
