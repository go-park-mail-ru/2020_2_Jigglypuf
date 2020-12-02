package csrf

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type MiddlewareCSRFTesting struct {
	mainStruct *HashCSRFToken
}

var (
	middlewareTesting *MiddlewareCSRFTesting = nil
)

func setUp() {
	middlewareTesting = new(MiddlewareCSRFTesting)
	middlewareTesting.mainStruct, _ = NewHashCSRFToken("some secret", time.Hour)
}

func TestGenerateCSRFHandler(t *testing.T) {
	setUp()

	testReq := httptest.NewRequest(http.MethodGet, "/csrf/generate", nil)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, session.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, session.ContextUserIDName, uint64(1))
	testRecorder := httptest.NewRecorder()

	middlewareTesting.mainStruct.GenerateCSRFToken(testRecorder, testReq.WithContext(ctx))
	assert.Equal(t, http.StatusOK, testRecorder.Code)

	resp := new(Response)
	castErr := json.Unmarshal(testRecorder.Body.Bytes(), &resp)
	if castErr != nil {
		t.Fatalf("TEST: Success csrf generate: incorrect response body format")
	}
	if resp.Token == "" {
		t.Fatalf("TEST: Success csrf generate: incorrect response body")
	}
}

func TestFailGenerateCSRF(t *testing.T) {
	setUp()
	testReq := httptest.NewRequest(http.MethodGet, "/csrf/generate", nil)
	testRecorder := httptest.NewRecorder()

	middlewareTesting.mainStruct.GenerateCSRFToken(testRecorder, testReq)
	assert.Equal(t, http.StatusUnauthorized, testRecorder.Code)
}

func TestIncorrectMethodGenerateCSRF(t *testing.T) {
	testReq := httptest.NewRequest(http.MethodPost, "/csrf/generate", nil)
	testRecorder := httptest.NewRecorder()

	middlewareTesting.mainStruct.GenerateCSRFToken(testRecorder, testReq)
	assert.Equal(t, http.StatusMethodNotAllowed, testRecorder.Code)
}

func TestCheckCSRFSuccess(t *testing.T) {
	setUp()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	csrfReq := httptest.NewRequest(http.MethodGet, "/somepath/", nil)
	ctx := csrfReq.Context()
	ctx = context.WithValue(ctx, session.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, session.ContextUserIDName, uint64(1))
	csrfRec := httptest.NewRecorder()
	middlewareTesting.mainStruct.GenerateCSRFToken(csrfRec, csrfReq.WithContext(ctx))
	assert.Equal(t, csrfRec.Code, http.StatusOK)
	csrfToken := new(Response)
	castErr := json.Unmarshal(csrfRec.Body.Bytes(), csrfToken)
	if castErr != nil {
		t.Fatalf("Incorrect generate token body")
	}

	testFunc := middlewareTesting.mainStruct.CSRFMiddleware(handler)
	testReq := httptest.NewRequest(http.MethodPost, "/somepath/", nil)
	ctxHandler := testReq.Context()
	ctxHandler = context.WithValue(ctxHandler, session.ContextIsAuthName, true)
	ctxHandler = context.WithValue(ctxHandler, session.ContextUserIDName, uint64(1))
	testReq.Header.Set("X-CSRF-Token", csrfToken.Token)
	newRec := httptest.NewRecorder()
	testFunc.ServeHTTP(newRec, testReq.WithContext(ctxHandler))
	assert.Equal(t, http.StatusOK, newRec.Code)
}

func TestCheckCSRFTokenError(t *testing.T) {
	setUp()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	testFunc := middlewareTesting.mainStruct.CSRFMiddleware(testHandler)
	var testCases = []struct {
		request    *http.Request
		statusCode int
	}{
		{
			httptest.NewRequest(http.MethodGet, "/somepath/", nil),
			200,
		},
		{
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
			200,
		},
	}

	for _, val := range testCases {
		recorder := httptest.NewRecorder()
		testFunc.ServeHTTP(recorder, val.request)
		assert.Equal(t, val.statusCode, recorder.Code)
	}
}

func TestIncorrectToken(t *testing.T) {
	setUp()
	pastTime := time.Now().Add(-time.Hour).Unix()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	testFunc := middlewareTesting.mainStruct.CSRFMiddleware(testHandler)
	var testCases = []struct {
		token   string
		request *http.Request
	}{
		{
			"asidnoaisnd:aoinoisadnoi:oiansidoansd",
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			"naiosdnaoidna:aoisdnoaisnd",
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			"ansiodsandsa:" + strconv.FormatInt(pastTime, 10),
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
		{
			"ansiodsandsa:" + strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10),
			httptest.NewRequest(http.MethodPost, "/somepath/", nil),
		},
	}

	for _, val := range testCases {
		ctx := val.request.Context()
		ctx = context.WithValue(ctx, session.ContextIsAuthName, true)
		ctx = context.WithValue(ctx, session.ContextUserIDName, uint64(1))
		val.request.Header.Set("X-CSRF-Token", val.token)
		Rec := httptest.NewRecorder()
		testFunc.ServeHTTP(Rec, val.request.WithContext(ctx))
		assert.Equal(t, http.StatusForbidden, Rec.Code)
	}
}
