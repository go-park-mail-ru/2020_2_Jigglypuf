package authentication

import(
	"models"
)

type AuthRepository interface{
	CreateUser(user *models.User) error
	GetUser(username string, password string) (*models.User,error)
}