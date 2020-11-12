//go:generate mockgen -source ProfileRepository.go -destination mock/ProfileRep_mock.go -package mock
package profile

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
)

type Repository interface {
	CreateProfile(profile *models.Profile) error
	DeleteProfile(profile *models.Profile) error
	GetProfile(login *string) (*models.Profile, error)
	UpdateCredentials(profile *models.Profile) error
	UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error
	GetProfileViaID(userID uint64) (*models.Profile, error)
}
