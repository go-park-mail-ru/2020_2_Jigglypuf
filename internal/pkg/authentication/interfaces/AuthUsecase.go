//go:generate mockgen -source AuthUsecase.go -destination mock/AuthenticationUC_mock.go -package mock
package interfaces

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"net/http"
)

type UserUseCase interface {
	SignUp(input *models.RegistrationInput) (uint64, error)
	SignIn(input *models.AuthInput) (uint64, error)
	SignOut(cookie *http.Cookie) (*http.Cookie, error)
}
