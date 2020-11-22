package authentication

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/delivery"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/interfaces"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/repository"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/authentication/usecase"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	profileMock "github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile/mock"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestingAuthenticationStruct struct {
	Handler          *delivery.UserHandler
	UseCase          interfaces.UserUseCase
	Repository       interfaces.AuthRepository
	ProfileRepMock   *profileMock.MockRepository
	DBConn           *sql.DB
	DBMock           sqlmock.Sqlmock
	GoMockController *gomock.Controller
}

type TestCase struct {
	TestName           string
	TestRequest        *http.Request
	TestResponse       http.Response
	TestResponseWriter *httptest.ResponseRecorder
	TestHandler        func(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

var (
	testingStruct *TestingAuthenticationStruct = nil
)

func setUp(t *testing.T) {
	testingStruct = new(TestingAuthenticationStruct)
	testingStruct.GoMockController = gomock.NewController(t)

	DBConnect, DBMock, DBErr := sqlmock.New()
	if DBErr != nil {
		t.Fatal(DBErr)
	}
	testingStruct.DBConn = DBConnect
	testingStruct.DBMock = DBMock
	profileRepFoo := profileMock.NewMockRepository(testingStruct.GoMockController)
	authRep := repository.NewAuthSQLRepository(DBConnect)
	authUC := usecase.NewUserUseCase(authRep, profileRepFoo, "testing_salt")
	authHandler := delivery.NewUserHandler(authUC)
	testingStruct.Repository = authRep
	testingStruct.UseCase = authUC
	testingStruct.Handler = authHandler
	testingStruct.ProfileRepMock = profileRepFoo
}

func tearDown() {
	testingStruct.GoMockController.Finish()
	_ = testingStruct.DBConn.Close()
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
	var value = TestCase{
		"Проверяем что register работает и регистрирует пользователя",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusOK},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}

	// проверка работы регистрации

	resultRows := []string{"ID"}
	testingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1))
	testingStruct.DBMock.ExpectCommit()

	testingStruct.ProfileRepMock.EXPECT().CreateProfile(gomock.Any()).Return(nil)

	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}

	tearDown()
}

func TestSignUpValidation(t *testing.T) {
	setUp(t)
	incorrectRegistrationModel := models.RegistrationInput{
		Login:    "Aydar",
		Password: "aydar",
		Name:     "kek",
		Surname:  "kekov",
	}

	incorrectRegistrationBody, _ := json.Marshal(incorrectRegistrationModel)
	var value = TestCase{
		"Проверяем что register валидирует почту",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(incorrectRegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}

	// проверка работы валидатора
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
	tearDown()
}

func TestUserIdentityCase(t *testing.T) {
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)
	var testCase = TestCase{
		"Проверяем что нельзя создать юзера с одинаковым логином",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}

	testingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnError(errors.New("user already exists"))
	testingStruct.DBMock.ExpectCommit()

	testCase.TestHandler(testCase.TestResponseWriter, testCase.TestRequest, httprouter.Params{})
	if testCase.TestResponseWriter.Code != testCase.TestResponse.StatusCode {
		t.Fatalf("TEST: "+testCase.TestName+" "+
			"handler returned wrong status code: got %v want %v", testCase.TestResponseWriter.Code, testCase.TestResponse.StatusCode)
	}
	tearDown()
}

func TestCheckSignUpErrors(t *testing.T) {
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)
	var value = TestCase{
		"Проверяем что ошибки обрабатываются",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}

	// check if database returned error
	testingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnError(errors.New("some error"))
	testingStruct.DBMock.ExpectCommit()
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
	tearDown()
}
func TestSignUpProfileError(t *testing.T) {
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)
	var value = TestCase{
		"Проверяем что ошибки обрабатываются",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}

	// check if profile returned error
	resultRows := []string{"ID"}
	testingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1))
	testingStruct.DBMock.ExpectCommit()
	testingStruct.ProfileRepMock.EXPECT().CreateProfile(gomock.Any()).Return(errors.New("some error"))
	value.TestResponseWriter = httptest.NewRecorder()
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
	tearDown()
}

func TestSignUpDeliveryError(t *testing.T) {
	setUp(t)
	correctRegistrationModel := models.RegistrationInput{
		Login:    "Aydar@mail.ru",
		Password: "aisndoandoqw",
		Name:     "kek",
		Surname:  "kekov",
	}
	RegistrationBody, _ := json.Marshal(correctRegistrationModel)
	var value = TestCase{
		"Проверяем что ошибки обрабатываются",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.RegisterHandler,
	}
	// check if method not allowed
	value.TestRequest = httptest.NewRequest(http.MethodGet, "/signup/", strings.NewReader("sinaoifnoisd"))
	value.TestResponseWriter = httptest.NewRecorder()
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, http.StatusMethodNotAllowed)
	}

	// check if input body is trash
	value.TestRequest = httptest.NewRequest("POST", "/signup/", strings.NewReader("sinaoifnoisd"))
	value.TestResponseWriter = httptest.NewRecorder()
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
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
	var testCase = TestCase{
		"Проверяем что signin работает и отдает куку",
		httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody))),
		http.Response{StatusCode: http.StatusOK},
		httptest.NewRecorder(),
		testingStruct.Handler.AuthHandler,
	}

	resultRows := []string{"ID", "Login", "Password"}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(correctAuthenticationModel.Password+"testing_salt"), 7)
	testingStruct.DBMock.ExpectQuery("SELECT .* FROM users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1, correctAuthenticationModel.Login,
		string(hashPassword)))
	testingStruct.DBMock.ExpectCommit()

	testCase.TestHandler(testCase.TestResponseWriter, testCase.TestRequest, httprouter.Params{})
	if testCase.TestResponseWriter.Code != testCase.TestResponse.StatusCode {
		t.Fatalf("TEST: "+testCase.TestName+" "+
			"handler returned wrong status code: got %v want %v", testCase.TestResponseWriter.Code, testCase.TestResponse.StatusCode)
	}

	tearDown()
}

func TestSignInValidator(t *testing.T) {
	setUp(t)
	incorrectAuthenticationModel := models.AuthInput{
		Login:    "someone",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(incorrectAuthenticationModel)
	var testCase = TestCase{
		"Проверяем что валидатор валидирует правильно",
		httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		testingStruct.Handler.AuthHandler,
	}

	testCase.TestHandler(testCase.TestResponseWriter, testCase.TestRequest, httprouter.Params{})
	if testCase.TestResponseWriter.Code != testCase.TestResponse.StatusCode {
		t.Fatalf("TEST: "+testCase.TestName+" "+
			"handler returned wrong status code: got %v want %v", testCase.TestResponseWriter.Code, testCase.TestResponse.StatusCode)
	}
	tearDown()
}

func TestSignInErrors(t *testing.T) {
	setUp(t)
	authenticationModel := models.AuthInput{
		Login:    "iasndia@ansoia.ru",
		Password: "voasndoiasndqw",
	}
	authenticationBody, _ := json.Marshal(authenticationModel)

	// check if method is incorrect
	testReq := httptest.NewRequest("GET", "/signin/", strings.NewReader(string(authenticationBody)))
	testRecorder := httptest.NewRecorder()
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: method not allowed "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	// check if input is incorrect
	testReq = httptest.NewRequest("POST", "/signin/", strings.NewReader("incorrect input"))
	testRecorder = httptest.NewRecorder()
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: incorrect input format "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	// check if input is empty
	emptyModel := models.AuthInput{}
	emptyBody, _ := json.Marshal(emptyModel)
	testReq = httptest.NewRequest("POST", "/signin/", strings.NewReader(string(emptyBody)))
	testRecorder = httptest.NewRecorder()
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: empty input model "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestSignInInternalDBError(t *testing.T) {
	setUp(t)
	authenticationModel := models.AuthInput{
		Login:    "iasndia@ansoia.ru",
		Password: "voasndoiasndqw",
	}
	authenticationBody, _ := json.Marshal(authenticationModel)
	// check if db returned error
	testingStruct.DBMock.ExpectQuery("SELECT .* FROM users .*").WillReturnError(errors.New("some error"))
	testingStruct.DBMock.ExpectCommit()
	testReq := httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody)))
	testRecorder := httptest.NewRecorder()
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: handling db error "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestSignInPasswordIncorrect(t *testing.T) {
	setUp(t)
	authenticationModel := models.AuthInput{
		Login:    "iasndia@ansoia.ru",
		Password: "voasndoiasndqw",
	}
	authenticationBody, _ := json.Marshal(authenticationModel)
	// check if password is incorrect
	resultRows := []string{"ID", "Login", "Password"}
	testingStruct.DBMock.ExpectQuery("SELECT .* FROM users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1,
		authenticationModel.Login, "another_password"))
	testingStruct.DBMock.ExpectCommit()
	testReq := httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody)))
	testRecorder := httptest.NewRecorder()
	testingStruct.Handler.AuthHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusBadRequest {
		t.Fatalf("TEST: handling incorrct pass "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}
	tearDown()
}

func TestLogOutSuccessCase(t *testing.T) {
	setUp(t)
	testReq := httptest.NewRequest(http.MethodPost, "/logout/", nil)
	testResponseRecorder := httptest.NewRecorder()
	testCookie := http.Cookie{
		Name:     session.SessionCookieName,
		Value:    models.RandStringRunes(32),
		Expires:  time.Now().Add(96 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	}
	ctx := testReq.Context()
	ctx = context.WithValue(ctx, session.ContextIsAuthName, true)
	ctx = context.WithValue(ctx, session.ContextUserIDName, 1)
	testReq.AddCookie(&testCookie)

	testingStruct.Handler.SignOutHandler(testResponseRecorder, testReq.WithContext(ctx), httprouter.Params{})
	if testResponseRecorder.Code != http.StatusOK {
		t.Fatalf("TEST: success logout status code "+
			"handler returned wrong status code: got %v want %v", testResponseRecorder.Code, http.StatusOK)
	}
	tearDown()
}

func TestLogOutErrors(t *testing.T) {
	setUp(t)

	// check if method is invalid
	testReq := httptest.NewRequest(http.MethodGet, "/logout", nil)
	testRecorder := httptest.NewRecorder()
	testingStruct.Handler.SignOutHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("TEST: method not allowed logout status code "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusMethodNotAllowed)
	}

	// check if no authorization
	testReq = httptest.NewRequest(http.MethodPost, "/logout", nil)
	testRecorder = httptest.NewRecorder()
	testingStruct.Handler.SignOutHandler(testRecorder, testReq, httprouter.Params{})
	if testRecorder.Code != http.StatusUnauthorized {
		t.Fatalf("TEST: unauthorized logout status code "+
			"handler returned wrong status code: got %v want %v", testRecorder.Code, http.StatusUnauthorized)
	}
	tearDown()
}
