//go:generate mockgen -source AuthRepository.go -destination mock/AuthenticationRep_mock.go -package mock
package interfaces

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUser(Login string) (*models.User, error)
	GetUserByID(userID uint64) (*models.User, error)
}
