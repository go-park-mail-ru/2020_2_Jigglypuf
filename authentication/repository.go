package authentication

import (
	"models"
	"net/http"
)

type AuthRepository interface{
	CreateUser(user *models.User) error
	GetUser(username string, password string) (*models.User,error)
	GetUserViaCookie(cookie *http.Cookie)(*models.User, error)
	SetCookie(user *models.User,cookieValue *http.Cookie) bool
}