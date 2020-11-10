//go:generate mockgen -source ProfileUseCase.go -destination mock/ProfileUC_mock.go -package mock
package profile

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type UseCase interface {
	CreateProfile(profile *models.Profile) error
	DeleteProfile(profile *models.Profile) error
	GetProfile(login *string) (*models.Profile, error)
	UpdateCredentials(profile *models.Profile) error
	UpdateProfile(profileUserID uint64, name string, surname string, avatarPath string) error
	GetProfileViaID(userID uint64) (*models.Profile, error)
}
