package authentication

import "models"

type UserUseCase interface{
	SignUp(input *models.RegistrationInput) error
	SignIn(input *models.AuthInput) error
	SignOut()error
}