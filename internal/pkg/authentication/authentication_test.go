package authentication

import (
	"backend/internal/pkg/authentication/delivery"
	"backend/internal/pkg/authentication/repository"
	"backend/internal/pkg/authentication/usecase"
	cookieMock "backend/internal/pkg/middleware/cookie/mock"
	profileMock "backend/internal/pkg/profile/mock"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestingAuthenticationStruct struct {
	Handler          *delivery.UserHandler
	UseCase          UserUseCase
	Repository       AuthRepository
	DBConn           *sql.DB
	DBMock           sqlmock.Sqlmock
	GoMockController *gomock.Controller
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
}

func tearDown() {
	TestingStruct.GoMockController.Finish()
	_ = TestingStruct.DBConn.Close()
}

func TestSignUp(t *testing.T) {
	setUp(t)
	incorrectAuthenticationBody := "{\"Login\": \"Aydar\", \"Password\": \"aydar\"}"
	incorrectRegistrationBody := "{\"Login\": \"Aydar\", \"Password\": \"aydar\"}"
	authenticationBody := "{\"Login\": \"Aydar@mail.ru\", \"Password\": \"aydar\"}"
	RegistrationBody := "{\"Login\": \"SomeOne@mail.ru\", \"Password\": \"aydar\"}"
	var testCases = []struct {
		TestName           string
		TestRequest        *http.Request
		TestResponse       http.Response
		TestResponseWriter *httptest.ResponseRecorder
		TestHandler        func(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	}{
		{
			"Проверяем что register работает и регистрирует пользователя",
				httptest.NewRequest("POST", "/signup/", strings.NewReader(RegistrationBody)),
				http.Response{StatusCode: http.StatusOK},
				httptest.NewRecorder(),
				TestingStruct.Handler.RegisterHandler,
		},
		{
			"Проверяем что signin работает и отдает куку",
				httptest.NewRequest("POST", "/signin/", strings.NewReader(authenticationBody)),
				http.Response{StatusCode: http.StatusOK},
				httptest.NewRecorder(),
				TestingStruct.Handler.AuthHandler,
		},
	}

	tearDown()
}
