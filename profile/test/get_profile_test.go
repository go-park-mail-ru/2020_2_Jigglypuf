package test

import (
	authDelivery "authentication/delivery"
	authRepository "authentication/repository"
	authUseCase "authentication/usecase"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	profileDelivery "profile/delivery"
	profileRepository "profile/repository"
	profileUseCase "profile/usecase"
	"strconv"
	"strings"
	"sync"
	"testing"
)


const (
	getProfileUrl = "/getprofile/"
	signUpUrl = "/signup/"
	salt = "oisndoiqwe123"
)


var bodies = map[string]string {
	"authorized": "{\"Name\":\"\",\"Surname\":\"\",\"AvatarPath\":\"\"}",
	"unauthorized": "{\"StatusCode\":401,\"Response\":\"TWV0aG9kTm90QWxsb3dlZA==\"}",
}

var data = map[string][]struct {
	expectedRequestBody string
	response int
}{
	"successRequest": {
		{bodies["authorized"], http.StatusOK},
	},
	"failureRequest": {
		{bodies["unauthorized"], http.StatusUnauthorized},
	},
}


func TestGetProfileCases(t *testing.T) {
	mutex := sync.RWMutex{}
	requestProfile := httptest.NewRequest("GET", getProfileUrl, nil)
	authR := authRepository.NewUserRepository(&mutex)
	authUC := authUseCase.NewUserUseCase(authR, salt)
	authHandler := authDelivery.NewUserHandler(authUC)
	profileUC := profileUseCase.NewProfileUseCase(profileRepository.NewProfileRepository(&mutex, authR))
	for testName, requestsSlice := range data {
		t.Run(testName, func(t *testing.T) {
			for _, request := range requestsSlice {


				if testName == "successRequest" {
					signUpInfo := strings.NewReader("{\"Login\": \"Pro11\", \"Password\": \"1234\"}")
					writerAuth := httptest.NewRecorder()
					requestAuth := httptest.NewRequest("POST", signUpUrl, signUpInfo)
					authHandler.RegisterHandler(writerAuth, requestAuth)
					requestProfile.Header.Set("Cookie", writerAuth.Header()["Set-Cookie"][0])
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

				require.Equal(t, request.expectedRequestBody, string(body), "Response must be OK")
				err = resp.Body.Close()
				if err != nil {
					t.Fatal(err)
				}
				requestProfile.Header.Del("Cookie")
			}
		})
	}

}