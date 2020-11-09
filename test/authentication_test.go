package test

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/authentication/delivery"
	"backend/internal/pkg/authentication/repository"
	"backend/internal/pkg/authentication/usecase"
	cookieMock "backend/internal/pkg/middleware/cookie/mock"
	"backend/internal/pkg/models"
	profileMock "backend/internal/pkg/profile/mock"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestingAuthenticationStruct struct {
	Handler          *delivery.UserHandler
	UseCase          authentication.UserUseCase
	Repository       authentication.AuthRepository
	ProfileRepMock   *profileMock.MockRepository
	CookieRepMock    *cookieMock.MockRepository
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
	TestingStruct *TestingAuthenticationStruct = nil
)

func setUp(t *testing.T) {
	TestingStruct = new(TestingAuthenticationStruct)
	TestingStruct.GoMockController = gomock.NewController(t)

	DBConnect, DBMock, DBErr := sqlmock.New()
	if DBErr != nil {
		t.Fatal(DBErr)
	}
	TestingStruct.DBConn = DBConnect
	TestingStruct.DBMock = DBMock
	profileRepFoo := profileMock.NewMockRepository(TestingStruct.GoMockController)
	cookieRepFoo := cookieMock.NewMockRepository(TestingStruct.GoMockController)
	authRep := repository.NewAuthSQLRepository(DBConnect)
	authUC := usecase.NewUserUseCase(authRep, profileRepFoo, cookieRepFoo, "testing_salt")
	authHandler := delivery.NewUserHandler(authUC)
	TestingStruct.Repository = authRep
	TestingStruct.UseCase = authUC
	TestingStruct.Handler = authHandler
	TestingStruct.ProfileRepMock = profileRepFoo
	TestingStruct.CookieRepMock = cookieRepFoo
}

func tearDown() {
	TestingStruct.GoMockController.Finish()
	_ = TestingStruct.DBConn.Close()
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
		TestingStruct.Handler.RegisterHandler,
	}

	// проверка работы регистрации

	resultRows := []string{"ID"}
	TestingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1))
	TestingStruct.DBMock.ExpectCommit()

	TestingStruct.ProfileRepMock.EXPECT().CreateProfile(gomock.Any()).Return(nil)
	TestingStruct.CookieRepMock.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(nil)

	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
	if value.TestResponse.StatusCode == http.StatusOK && value.TestResponseWriter.Header()["Set-Cookie"][0] == "" {
		t.Errorf("handler doesn`t returned cookie")
	}

	tearDown()
}

func TestSignUpValidation(t *testing.T) {
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
		TestingStruct.Handler.RegisterHandler,
	}

	// проверка работы валидатора
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
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
		TestingStruct.Handler.RegisterHandler,
	}

	TestingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnError(errors.New("user already exists"))
	TestingStruct.DBMock.ExpectCommit()

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
		"Проверяем что register работает и регистрирует пользователя",
		httptest.NewRequest("POST", "/signup/", strings.NewReader(string(RegistrationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		TestingStruct.Handler.RegisterHandler,
	}

	// check if database returned error
	TestingStruct.DBMock.ExpectQuery("INSERT INTO users .*").WillReturnError(errors.New("some error"))
	TestingStruct.DBMock.ExpectCommit()
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}

	// check if profile returned error
	TestingStruct.ProfileRepMock.EXPECT().CreateProfile(gomock.Any()).Return(errors.New("some error"))
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}

	// check if cookie service returned error
	TestingStruct.CookieRepMock.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
	value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
	if value.TestResponseWriter.Code != http.StatusBadRequest {
		t.Fatalf("TEST: "+value.TestName+" "+
			"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
	}
}

func TestSignInSuccessCase(t *testing.T) {
	setUp(t)
	correctAuthenticationModel := models.AuthInput{
		Login: "someone@somene.ru",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(correctAuthenticationModel)
	var testCase = TestCase{
		"Проверяем что signin работает и отдает куку",
		httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody))),
		http.Response{StatusCode: http.StatusOK},
		httptest.NewRecorder(),
		TestingStruct.Handler.AuthHandler,
	}


	resultRows := []string{"ID","Login","Password"}
	hashPassword,_ := bcrypt.GenerateFromPassword([]byte(correctAuthenticationModel.Password + "testing_salt"),7)
	TestingStruct.DBMock.ExpectQuery("SELECT .* FROM users .*").WillReturnRows(sqlmock.NewRows(resultRows).AddRow(1,correctAuthenticationModel.Login, 
		string(hashPassword)))
	TestingStruct.DBMock.ExpectCommit()
	
	TestingStruct.CookieRepMock.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(nil)

	testCase.TestHandler(testCase.TestResponseWriter, testCase.TestRequest, httprouter.Params{})
	if testCase.TestResponseWriter.Code != testCase.TestResponse.StatusCode {
		t.Fatalf("TEST: "+testCase.TestName+" "+
			"handler returned wrong status code: got %v want %v", testCase.TestResponseWriter.Code, testCase.TestResponse.StatusCode)
	}
	if testCase.TestResponse.StatusCode == http.StatusOK && testCase.TestResponseWriter.Header()["Set-Cookie"][0] == "" {
		t.Errorf("handler doesn`t returned cookie")
	}
	tearDown()
}

func TestSignInValidator(t *testing.T){
	setUp(t)
	incorrectAuthenticationModel := models.AuthInput{
		Login: "someone",
		Password: "voasndoiasndqw",
	}

	authenticationBody, _ := json.Marshal(incorrectAuthenticationModel)
	var testCase = TestCase{
		"Проверяем что валидатор валидирует",
		httptest.NewRequest("POST", "/signin/", strings.NewReader(string(authenticationBody))),
		http.Response{StatusCode: http.StatusBadRequest},
		httptest.NewRecorder(),
		TestingStruct.Handler.AuthHandler,
	}

	testCase.TestHandler(testCase.TestResponseWriter, testCase.TestRequest, httprouter.Params{})
	if testCase.TestResponseWriter.Code != testCase.TestResponse.StatusCode {
		t.Fatalf("TEST: "+testCase.TestName+" "+
			"handler returned wrong status code: got %v want %v", testCase.TestResponseWriter.Code, testCase.TestResponse.StatusCode)
	}
	tearDown()
}

//func TestSignUp(t *testing.T) {
//	setUp(t)
//	RegistrationBody := "{\"Login\": \"SomeOne@mail.ru\", \"Password\": \"aydar\"}"
//	JailRegistrationBody := "iasndoiansdoi"
//	var testCases = []struct {
//		TestName           string
//		TestRequest        *http.Request
//		TestResponse       http.Response
//		TestResponseWriter *httptest.ResponseRecorder
//		TestHandler        func(w http.ResponseWriter, r *http.Request, params httprouter.Params)
//	}{
//		{
//			"Проверяем translationError при отправке неправильного тела на регистрацию",
//			httptest.NewRequest("POST", "/signup/", strings.NewReader(JailRegistrationBody)),
//			http.Response{StatusCode: http.StatusBadRequest},
//			httptest.NewRecorder(),
//			TestingStruct.Handler.RegisterHandler,
//		},
//		{
//			"Проверяем MethodNotAllowed при отправке запроса с неправильным методом на регистрацию",
//			httptest.NewRequest("GET", "/signup/", strings.NewReader(RegistrationBody)),
//			http.Response{StatusCode: http.StatusMethodNotAllowed},
//			httptest.NewRecorder(),
//			TestingStruct.Handler.RegisterHandler,
//		},
//	}
//
//	for _, value := range testCases {
//		value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
//		if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
//			t.Fatalf("TEST: "+value.TestName+" "+
//				"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
//		}
//		if value.TestResponse.StatusCode == http.StatusOK && value.TestResponseWriter.Header()["Set-Cookie"][0] == "" {
//			t.Errorf("handler doesn`t returned cookie")
//		}
//	}
//
//	tearDown()
//}
//
//func TestSignIn(t *testing.T) {
//	setUp(t)
//	incorrectAuthenticationBody := "{\"Login\": \"Aydar\", \"Password\": \"aydar\"}"
//	authenticationBody := "{\"Login\": \"Aydar@mail.ru\", \"Password\": \"aydar\"}"
//	var testCases = []struct {
//		TestName           string
//		TestRequest        *http.Request
//		TestResponse       http.Response
//		TestResponseWriter *httptest.ResponseRecorder
//		TestHandler        func(w http.ResponseWriter, r *http.Request, params httprouter.Params)
//	}{
//		{
//			"Проверяем что signin работает и отдает куку",
//			httptest.NewRequest("POST", "/signin/", strings.NewReader(authenticationBody)),
//			http.Response{StatusCode: http.StatusOK},
//			httptest.NewRecorder(),
//			TestingStruct.Handler.AuthHandler,
//		},
//		{
//			"Проверяем что signin валидирует почту",
//			httptest.NewRequest("POST", "/signin/", strings.NewReader(incorrectAuthenticationBody)),
//			http.Response{StatusCode: http.StatusOK},
//			httptest.NewRecorder(),
//			TestingStruct.Handler.AuthHandler,
//		},
//	}
//
//	for _, value := range testCases {
//		value.TestHandler(value.TestResponseWriter, value.TestRequest, httprouter.Params{})
//		if value.TestResponseWriter.Code != value.TestResponse.StatusCode {
//			t.Fatalf("TEST: "+value.TestName+" "+
//				"handler returned wrong status code: got %v want %v", value.TestResponseWriter.Code, value.TestResponse.StatusCode)
//		}
//		if value.TestResponse.StatusCode == http.StatusOK && value.TestResponseWriter.Header()["Set-Cookie"][0] == "" {
//			t.Errorf("handler doesn`t returned cookie")
//		}
//	}
//
//	tearDown()
//}
