package authentication

import (
	"backend/internal/pkg/models"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUser(username string, password string) (*models.User, error)
	GetUserByID(userID uint64) (*models.User, error)
}
