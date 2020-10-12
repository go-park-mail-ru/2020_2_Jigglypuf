package tests

import (
	authDelivery "backend/internal/pkg/authentication/delivery"
	authRepository "backend/internal/pkg/authentication/repository"
	authUseCase "backend/internal/pkg/authentication/usecase"
	profileDelivery "backend/internal/pkg/profile/delivery"
	profileRepository "backend/internal/pkg/profile/repository"
	profileUseCase "backend/internal/pkg/profile/usecase"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
)

const (
	getProfileURL = "/getprofile/"
	signUpURL     = "/signup/"
	salt          = "oisndoiqwe123"
)

var bodiesResponse = map[string]string{
	"authorized":   `{"Name":"","Surname":"","AvatarPath":""}`,
	"unauthorized": `{"StatusCode":401,"Response":"You not authorized!"}`,
}

var bodiesRequest = map[string]string{
	"authorized":   `{"Login": "Pro100", "Password": "1234"}`,
	"unauthorized": "",
}

var data = map[string][]struct {
	expectedRequestBody  string
	expectedResponseBody string
	response             int
}{
	"successRequest": {
		{bodiesRequest["authorized"], bodiesResponse["authorized"], http.StatusOK},
	},
	"failureRequest": {
		{bodiesRequest["unauthorized"], bodiesResponse["unauthorized"], http.StatusUnauthorized},
	},
}

func TestGetProfileCases(t *testing.T) {
	mutex := sync.RWMutex{}
	requestProfile := httptest.NewRequest("GET", getProfileURL, nil)
	authR := authRepository.NewUserRepository(&mutex)
	authUC := authUseCase.NewUserUseCase(authR, salt)
	authHandler := authDelivery.NewUserHandler(authUC)
	profileUC := profileUseCase.NewProfileUseCase(profileRepository.NewProfileRepository(&mutex, authR))
	for testName, requestsSlice := range data {
		requestsSlice := requestsSlice
		t.Run(testName, func(t *testing.T) {
			for _, request := range requestsSlice {
				signUpInfo := strings.NewReader(request.expectedRequestBody)
				writerAuth := httptest.NewRecorder()
				requestAuth := httptest.NewRequest("POST", signUpURL, signUpInfo)
				authHandler.RegisterHandler(writerAuth, requestAuth)
				if cookie := writerAuth.Header()["Set-Cookie"]; cookie != nil {
					requestProfile.Header.Set("Cookie", cookie[0])
				}

				profileHandler := profileDelivery.NewProfileHandler(profileUC)
				writerProfile := httptest.NewRecorder()
				profileHandler.GetProfile(writerProfile, requestProfile)
				resp := writerProfile.Result()
				if resp.StatusCode != request.response {
					t.Errorf(strconv.Itoa(resp.StatusCode))
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}

				require.Equal(t, request.expectedResponseBody, string(body), "Response must be OK")
				requestProfile.Header.Del("Cookie")
				err = resp.Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
