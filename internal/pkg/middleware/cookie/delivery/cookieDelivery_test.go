package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestingCookieStruct struct{
	handler *CookieHandler
	useCaseMock *mock.MockUseCase
	goMockController *gomock.Controller
}

var(
	testingStruct *TestingCookieStruct = nil
)

func setUp(t *testing.T){
	testingStruct = new(TestingCookieStruct)
	testingStruct.goMockController = gomock.NewController(t)

	testingStruct.useCaseMock = mock.NewMockUseCase(testingStruct.goMockController)
	testingStruct.handler = NewCookieHandler(testingStruct.useCaseMock)
}

func tearDown(){
	testingStruct.goMockController.Finish()
}

func TestCookieSuccessDelivery(t *testing.T){
	setUp(t)
	cookieValue := http.Cookie{
		Name:     cookie.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}

	testReq := httptest.NewRequest("GET", "/cookie/", nil)
	testReq.AddCookie(&cookieValue)
	testingStruct.useCaseMock.EXPECT().CheckCookie(gomock.Any()).Return(uint64(1), true)
	_, returnErr := testingStruct.handler.CheckCookie(testReq)
	if !returnErr{
		t.Fatalf("TEST: Success cookie")
	}

	tearDown()
}


func TestCookieDeliveryFailCase(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cookie/", nil)
	_, returnErr := testingStruct.handler.CheckCookie(testReq)
	if returnErr{
		t.Fatalf("TEST: Failure cookie")
	}

	tearDown()
}