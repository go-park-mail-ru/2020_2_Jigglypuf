//go:generate mockgen -source repository.go -destination mock/AuthenticationRep_mock.go -package mock
package authentication

import (
	"backend/internal/pkg/models"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUser(Login string) (*models.User, error)
	GetUserByID(userID uint64) (*models.User, error)
}
