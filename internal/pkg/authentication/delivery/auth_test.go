package delivery

import (
	"backend/internal/pkg/authentication/mock"
	"backend/internal/pkg/middleware/cookie"
	"backend/internal/pkg/models"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestingAuthStruct struct{
	Handler *UserHandler
	UseCaseMock *mock.MockUserUseCase
	GoMockController *gomock.Controller
}

var(
	TestingStruct *TestingAuthStruct = nil
)

func setUp(t *testing.T){
	TestingStruct = new(TestingAuthStruct)
	TestingStruct.GoMockController = gomock.NewController(t)

	UseCaseMock := mock.NewMockUserUseCase(TestingStruct.GoMockController)
	TestingStruct.UseCaseMock = UseCaseMock
	Handler := NewUserHandler(UseCaseMock)
	TestingStruct.Handler = Handler
}

func tearDown(){
	TestingStruct.GoMockController.Finish()
}

func TestSignUpSuccessCase(t *testing.T){
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)
	// проверка работы регистрации
	cookieValue := http.Cookie{
		Name:     cookie.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}

	testReq := httptest.NewRequest(http.MethodPost, "/auth/",strings.NewReader(string(RegistrationBody)))
	testRecorder := httptest.NewRecorder()

	TestingStruct.UseCaseMock.EXPECT().SignUp(gomock.Any()).Return(&cookieValue, nil)
	TestingStruct.Handler.RegisterHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: Success sign up "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusOK)
	}
	if testRecorder.Code == http.StatusOK && testRecorder.Header()["Set-Cookie"][0] == "" {
		t.Errorf("handler doesn`t returned cookie")
	}
	tearDown()
}

func TestSignUpInvalidMethodErrorHandling(t *testing.T){
	setUp(t)
	TestRequest :=  httptest.NewRequest(http.MethodGet,"/signup/",strings.NewReader("sinaoifnoisd"))
	TestResponseWriter := httptest.NewRecorder()
	TestingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: invalid method error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestSignUpInvalidBodyErrorHandling(t *testing.T){
	setUp(t)
	TestRequest :=  httptest.NewRequest(http.MethodPost,"/signup/",strings.NewReader("sinaoifnoisd"))
	TestResponseWriter := httptest.NewRecorder()
	TestingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: invalid body error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusBadRequest)
	}
	tearDown()
}

func TestSignUpUseCaseErrorHandling(t *testing.T){
	setUp(t)

	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)

	TestRequest :=  httptest.NewRequest(http.MethodPost,"/signup/",strings.NewReader(string(RegistrationBody)))
	TestResponseWriter := httptest.NewRecorder()
	TestingStruct.UseCaseMock.EXPECT().SignUp(gomock.Any()).Return(nil,errors.New("test err"))
	TestingStruct.Handler.RegisterHandler(TestResponseWriter, TestRequest, httprouter.Params{})
	if TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: invalid body error handling "+
			"handler returned wrong status code: got %v want %v", TestResponseWriter.Code, http.StatusBadRequest)
	}

	tearDown()
}

//func TestSignInSuccessCase(t *testing.T)