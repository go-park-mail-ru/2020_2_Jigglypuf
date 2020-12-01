package delivery

import (
	"context"
	"encoding/json"
	"errors"
	mock "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces/mock"
	authService "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/proto/codegen"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestingAuthStruct struct {
	Handler          *UserHandler
	UseCaseMock      *mock.MockAuthManager
	GoMockController *gomock.Controller
}

var (
	testingStruct *TestingAuthStruct = nil
)

func setUp(t *testing.T) {
	testingStruct = new(TestingAuthStruct)
	testingStruct.GoMockController = gomock.NewController(t)
	UseCaseMock := mock.NewMockAuthManager(testingStruct.GoMockController)
	testingStruct.UseCaseMock = UseCaseMock
	Handler := NewUserHandler(UseCaseMock)
	testingStruct.Handler = Handler
}

func tearDown() {
	testingStruct.GoMockController.Finish()
}

func TestSignUpSuccessCase(t *testing.T) {
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)

	testReq := httptest.NewRequest(http.MethodPost, "/auth/", strings.NewReader(string(RegistrationBody)))
	testRecorder := httptest.NewRecorder()

	testingStruct.UseCaseMock.EXPECT().SignUp(gomock.Any(),gomock.Any()).Return(&authService.Response{UserID: uint64(1)}, nil)
	testingStruct.Handler.RegisterHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Success sign up "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	tearDown()
}

func TestSignUpInvalidMethodErrorHandling(t *testing.T) {
	setUp(t)
	TestRequest := httptest.NewRequest(http.MethodGet, "/signup/", strings.NewReader("sinaoifnoisd"))
	TestResponseWriter := httptest.NewRecorder()
	testingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: invalid method error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestSignUpInvalidBodyErrorHandling(t *testing.T) {
	setUp(t)
	TestRequest := httptest.NewRequest(http.MethodPost, "/signup/", strings.NewReader("sinaoifnoisd"))
	TestResponseWriter := httptest.NewRecorder()
	testingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: invalid body error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusBadRequest)
	}
	tearDown()
}

func TestSignUpUseCaseErrorHandling(t *testing.T) {
	setUp(t)

	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)

	TestRequest := httptest.NewRequest(http.MethodPost, "/signup/", strings.NewReader(string(RegistrationBody)))
	TestResponseWriter := httptest.NewRecorder()
	testingStruct.UseCaseMock.EXPECT().SignUp(gomock.Any(),gomock.Any()).Return(nil, errors.New("test err"))
	testingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: invalid body error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusBadRequest)
	}

	tearDown()
}

func TestSignInSuccessCase(t *testing.T) {
	setUp(t)

	correctAuthenticationModel := models.AuthInput{
		Login:    "someone@somene.ru",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(correctAuthenticationModel)

	testReq := httptest.NewRequest(http.MethodPost, "/login/", strings.NewReader(string(authenticationBody)))
	testRecorder := httptest.NewRecorder()

	testingStruct.UseCaseMock.EXPECT().SignIn(gomock.Any(),gomock.Any()).Return(&authService.Response{UserID: uint64(1)}, nil)
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Success log in test case "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}

	tearDown()
}

func TestLogInInvalidMethod(t *testing.T) {
	setUp(t)

	testReq := httptest.NewRequest(http.MethodGet, "/login/", nil)
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: Invalid method log in test case "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	tearDown()
}

func TestLogInIncorrectInput(t *testing.T) {
	setUp(t)
	testReq := httptest.NewRequest(http.MethodPost, "/login/", strings.NewReader("isndfosnfnosdf"))
	testRecorder := httptest.NewRecorder()

	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: Invalid input log in test case "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	tearDown()
}

func TestLogInUCErrorHandling(t *testing.T) {
	setUp(t)
	correctAuthenticationModel := models.AuthInput{
		Login:    "someone@somene.ru",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(correctAuthenticationModel)

	testReq := httptest.NewRequest(http.MethodPost, "/login/", strings.NewReader(string(authenticationBody)))
	testRecorder := httptest.NewRecorder()

	testingStruct.UseCaseMock.EXPECT().SignIn(gomock.Any(),gomock.Any()).Return(nil, errors.New("test error"))
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: Use case error log in test case "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusBadRequest)
	}
	tearDown()
}

func TestLogOutSuccessCase(t *testing.T) {
	setUp(t)

	testReq := httptest.NewRequest(http.MethodPost, "/logout/", nil)
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, session.ContextIsAuthName, true)
	testRecorder := httptest.NewRecorder()
	cookieValue := http.Cookie{
		Name: "session_id",
		Value: "New_value",
	}

	testReq.AddCookie(&cookieValue)
	//testingStruct.UseCaseMock.EXPECT().S(gomock.Any()).Return(&cookieValue, nil)
	testingStruct.Handler.SignOutHandler(testRecorder, testReq.WithContext(ctx), httprouter.Params{})
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Success log out test case "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}

	tearDown()
}

func TestLogOutErrorsHandling(t *testing.T) {
	setUp(t)

	testCases := []struct {
		TestRequest  *http.Request
		TestRecorder *httptest.ResponseRecorder
		ResponseCode int
	}{
		{
			httptest.NewRequest(http.MethodGet, "/logout/", nil),
			httptest.NewRecorder(),
			405,
		},
		{
			httptest.NewRequest(http.MethodPost, "/logout/", strings.NewReader("aisndoansd")),
			httptest.NewRecorder(),
			401,
		},
	}

	for _, val := range testCases {
		testingStruct.Handler.SignOutHandler(val.TestRecorder, val.TestRequest, httprouter.Params{})
		if val.TestRecorder.Code != val.ResponseCode {
			t.Fatalf("TEST: Success log out test case "+
				"handler returned wrong status code: got %v want %v", val.TestRecorder.Code, val.ResponseCode)
		}
	}

	tearDown()
}

func TestLogOutUCErrorHandling(t *testing.T) {
	setUp(t)
	correctAuthenticationModel := models.AuthInput{
		Login:    "someone@somene.ru",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(correctAuthenticationModel)

	testReq := httptest.NewRequest(http.MethodPost, "/logout", strings.NewReader(string(authenticationBody)))
	testRec := httptest.NewRecorder()
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, session.ContextIsAuthName, true)

	//testingStruct.UseCaseMock.EXPECT().SignOut(gomock.Any()).Return(nil, errors.New("some error"))
	testingStruct.Handler.SignOutHandler(testRec, testReq.WithContext(ctx), httprouter.Params{})

	if testRec.Code != http.StatusUnauthorized {
		t.Fatalf("TEST: Success log out test case "+
			"handler returned wrong status code: got %v want %v", testRec.Code, http.StatusUnauthorized)
	}
	tearDown()
}
