package authentication

import (
	"backend/internal/pkg/models"
	"net/http"
)

type UserUseCase interface{
	SignUp(input *models.RegistrationInput)(*http.Cookie, error)
	SignIn(input *models.AuthInput) (*http.Cookie, error)
	SignOut()error
}