package tests

import (
	"authentication/delivery"
	"authentication/repository"
	"authentication/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

const (
	serverPort = "8080"
	salt = "someSaltTest"
)



func TestAuthenticationAPISuccessCases(t *testing.T){
	mutex := sync.RWMutex{}
	authrep := repository.NewUserRepository(&mutex)
	authUseCase := usecase.NewUserUseCase(authrep, salt)
	authHandler := delivery.NewUserHandler(authUseCase)

	//authenticationBody := "{\"Login\": \"Aydar\", \"Password\": \"aydar\"}"
	authenticationBody := "{\"Login\": \"Pro100\", \"Password\": \"1234\"}"

	var testCases = []struct{
		TestName string
		TestRequest *http.Request
		TestResponse http.Response
		TestResponseWriter *httptest.ResponseRecorder
		TestHandler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			"Проверяем что register работает и регистрирует пользователя",
			httptest.NewRequest("POST","/signup/", strings.NewReader(authenticationBody)),
			http.Response{StatusCode: http.StatusOK},
			httptest.NewRecorder(),
			authHandler.RegisterHandler,
		},
		{
			"Проверяем что signin работает и отдает куку",
			httptest.NewRequest("POST","/signin/", strings.NewReader(authenticationBody)),
			http.Response{StatusCode: http.StatusOK},
			httptest.NewRecorder(),
			authHandler.AuthHandler,
		},
	}

	for _, val := range testCases{
		val.TestHandler(val.TestResponseWriter, val.TestRequest)
		if val.TestResponseWriter.Code != val.TestResponse.StatusCode{
			t.Fatalf("TEST: " + val.TestName + " " +
				"handler returned wrong status code: got %v want %v",val.TestResponseWriter.Code, val.TestResponse.StatusCode)
		}
		if val.TestResponseWriter.Header()["Set-Cookie"][0] == ""{
			t.Errorf("handler doesn`t returned cookie")
		}
	}

}

func TestAuthenticationAPIFAILCases(t *testing.T) {
	mutex := sync.RWMutex{}
	authrep := repository.NewUserRepository(&mutex)
	authUseCase := usecase.NewUserUseCase(authrep, salt)
	authHandler := delivery.NewUserHandler(authUseCase)

	incorrectAuthenticationBody := "{ \"inasd\": \"Aydar\",\"oainsd\":\"aydar\""
	authenticationBody := "{ \"Login\": \"Aydar\", \"Password\": \"aydar\" }"
	fakeAuthenticationBody := "{ \"Login\": \"Bulat\",\"Password\":\"aydar\"}"

	//создадим пользователя, чтобы убедиться что нельзя повторно создать пользователя
	authHandler.RegisterHandler(httptest.NewRecorder(), httptest.NewRequest("POST", "/signup",strings.NewReader(authenticationBody)))

	var testCases = []struct{
		TestName string
		TestRequest *http.Request
		TestResponse http.Response
		TestResponseWriter *httptest.ResponseRecorder
		TestHandler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			"Проверяем translationError при отправке неправильного тела на регистрацию",
			httptest.NewRequest("POST","/signup/", strings.NewReader(incorrectAuthenticationBody)),
			http.Response{StatusCode: http.StatusBadRequest},
			httptest.NewRecorder(),
			authHandler.RegisterHandler,
		},
		{
			"Проверяем translationError при отправке неправильного тела на авторизацию",
			httptest.NewRequest("POST","/signin/", strings.NewReader(incorrectAuthenticationBody)),
			http.Response{StatusCode: http.StatusBadRequest},
			httptest.NewRecorder(),
			authHandler.AuthHandler,
		},
		{
			"Проверяем MethodNotAllowed при отправке запроса с неправильным методом на регистрацию",
			httptest.NewRequest("GET","/signup/", strings.NewReader(incorrectAuthenticationBody)),
			http.Response{StatusCode: http.StatusMethodNotAllowed},
			httptest.NewRecorder(),
			authHandler.RegisterHandler,
		},
		{
			"Проверяем MethodNotAllowed при отправке запроса с неправильным методом на авторизацию",
			httptest.NewRequest("GET","/signin/", strings.NewReader(incorrectAuthenticationBody)),
			http.Response{StatusCode: http.StatusMethodNotAllowed},
			httptest.NewRecorder(),
			authHandler.AuthHandler,
		},
		{
			"Проверяем что нельзя создать юзера с одинаковым логином",
			httptest.NewRequest("POST","/signup/", strings.NewReader(authenticationBody)),
			http.Response{StatusCode: http.StatusBadRequest},
			httptest.NewRecorder(),
			authHandler.RegisterHandler,
		},
		{
			"Проверяем что нельзя зайти за  юзера которого не существует",
			httptest.NewRequest("POST","/signin/", strings.NewReader(fakeAuthenticationBody)),
			http.Response{StatusCode: http.StatusBadRequest},
			httptest.NewRecorder(),
			authHandler.AuthHandler,
		},
	}

	for _, val := range testCases{
		val.TestHandler(val.TestResponseWriter, val.TestRequest)
		if val.TestResponseWriter.Code != val.TestResponse.StatusCode{
			t.Fatalf("TEST: " + val.TestName + " " +"handler returned wrong status code: got %v want %v",val.TestResponseWriter.Code, val.TestResponse.StatusCode)
		}
	}
}
