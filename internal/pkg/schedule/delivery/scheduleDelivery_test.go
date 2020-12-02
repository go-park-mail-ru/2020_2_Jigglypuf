package delivery

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/schedule/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ScheduleTesting struct {
	handler          *ScheduleDelivery
	UseCaseMock      *mock.MockTimeTableUseCase
	GoMockController *gomock.Controller
}

var (
	testingStruct *ScheduleTesting = nil
)

func setup(t *testing.T) {
	testingStruct = new(ScheduleTesting)
	testingStruct.GoMockController = gomock.NewController(t)
	testingStruct.UseCaseMock = mock.NewMockTimeTableUseCase(testingStruct.GoMockController)
	testingStruct.handler = NewScheduleDelivery(testingStruct.UseCaseMock)
}

func teardown() {
	testingStruct.GoMockController.Finish()
}

func TestGetMovieScheduleSuccess(t *testing.T) {
	setup(t)
	returnArr := []models.Schedule{
		{
			PremierTime: time.Now(),
		},
	}
	testReq := httptest.NewRequest(http.MethodGet, "/somepath/?movie_id=1&cinema_id=2&date=2020-02-12",
		nil)
	testingStruct.UseCaseMock.EXPECT().GetMovieSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(&returnArr, nil)
	testRec := httptest.NewRecorder()

	testingStruct.handler.GetMovieSchedule(testRec, testReq)
	assert.Equal(t, http.StatusOK, testRec.Code)

	resultArr := new([]models.Schedule)
	castErr := json.Unmarshal(testRec.Body.Bytes(), resultArr)
	if castErr != nil {
		t.Fatalf("TEST: Success get movie schedule: incorrect body format")
	}
	teardown()
}

func TestGetMovieScheduleErrorsHandling(t *testing.T) {
	setup(t)

	var testCases = []struct {
		request    *http.Request
		call       *gomock.Call
		statusCode int
	}{
		{
			httptest.NewRequest(http.MethodPost, "/somePath/", nil),
			nil,
			405,
		},
		{
			httptest.NewRequest(http.MethodGet, "/somePath/", nil),
			testingStruct.UseCaseMock.EXPECT().GetMovieSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("some error")),
			400,
		},
	}

	for _, val := range testCases {
		testRec := httptest.NewRecorder()
		testingStruct.handler.GetMovieSchedule(testRec, val.request)
		assert.Equal(t, val.statusCode, testRec.Code)
	}

	teardown()
}

func TestGetScheduleSuccess(t *testing.T) {
	setup(t)

	testReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)

	mux.SetURLVars(testReq, map[string]string{
		schedule.ScheduleID: "1",
	})

	testRec := httptest.NewRecorder()
	returnItem := new(models.Schedule)
	testingStruct.UseCaseMock.EXPECT().GetSchedule(gomock.Any()).Return(returnItem, nil)

	testingStruct.handler.GetSchedule(testRec, testReq)
	assert.Equal(t, http.StatusOK, testRec.Code)

	teardown()
}

func TestGetScheduleErrorHandling(t *testing.T) {
	setup(t)

	var testCases = []struct {
		request    *http.Request
		statusCode int
		call       *gomock.Call
	}{
		{
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
			405,
			nil,
		},
		{
			httptest.NewRequest(http.MethodGet, "/somepath/", nil),
			400,
			testingStruct.UseCaseMock.EXPECT().GetSchedule(gomock.Any()).Return(nil, errors.New("some error")),
		},
	}

	for _, val := range testCases {
		mux.SetURLVars(val.request, map[string]string{
			schedule.ScheduleID: "1",
		})
		testRec := httptest.NewRecorder()

		testingStruct.handler.GetSchedule(testRec, val.request)
		assert.Equal(t, val.statusCode, testRec.Code)
	}
}
