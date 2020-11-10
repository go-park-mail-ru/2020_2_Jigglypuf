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
	TestingStruct *TestingCookieStruct = nil
)

func setUp(t *testing.T){
	TestingStruct = new(TestingCookieStruct)
	TestingStruct.goMockController = gomock.NewController(t)

	TestingStruct.useCaseMock = mock.NewMockUseCase(TestingStruct.goMockController)
	TestingStruct.handler = NewCookieHandler(TestingStruct.useCaseMock)
}

func tearDown(){
	TestingStruct.goMockController.Finish()
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
	TestingStruct.useCaseMock.EXPECT().CheckCookie(gomock.Any()).Return(uint64(1), true)
	_, returnErr := TestingStruct.handler.CheckCookie(testReq)
	if !returnErr{
		t.Fatalf("TEST: Success cookie")
	}

	tearDown()
}


func TestCookieDeliveryFailCase(t *testing.T){
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/cookie/", nil)
	_, returnErr := TestingStruct.handler.CheckCookie(testReq)
	if returnErr{
		t.Fatalf("TEST: Failure cookie")
	}

	tearDown()
}