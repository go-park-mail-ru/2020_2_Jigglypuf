package middleware

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type CookieMiddlewareTesting struct {
	delivery         *delivery.CookieHandler
	useCase          *mock.MockUseCase
	goMockController *gomock.Controller
}

var (
	middlewareTesting *CookieMiddlewareTesting = nil
)

func setUp(t *testing.T) {
	middlewareTesting = new(CookieMiddlewareTesting)
	middlewareTesting.goMockController = gomock.NewController(t)
	middlewareTesting.useCase = mock.NewMockUseCase(middlewareTesting.goMockController)
	middlewareTesting.delivery = delivery.NewCookieHandler(middlewareTesting.useCase)
}

func tearDown() {
	middlewareTesting.goMockController.Finish()
}

func TestCookieMiddlewareSuccess(t *testing.T) {
	setUp(t)
	cookieValue := http.Cookie{
		Name:     session.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}
	isAuth := false
	var userID uint64 = 0
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuth = r.Context().Value(session.ContextIsAuthName).(bool)
		userID = r.Context().Value(session.ContextUserIDName).(uint64)
	})
	testFunc := CookieMiddleware(testHandler, middlewareTesting.delivery)
	testReq := httptest.NewRequest("GET", "/session/", nil)
	testReq.AddCookie(&cookieValue)
	middlewareTesting.useCase.EXPECT().CheckCookie(gomock.Any()).Return(uint64(1), true)
	testRecorder := httptest.NewRecorder()
	testFunc.ServeHTTP(testRecorder, testReq)
	assert.Equal(t, http.StatusOK, testRecorder.Code)
	assert.Equal(t, true, isAuth)
	assert.Equal(t, uint64(1), userID)
	tearDown()
}
