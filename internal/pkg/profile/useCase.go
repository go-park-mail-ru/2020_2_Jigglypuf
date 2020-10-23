package profile

import (
	"backend/internal/pkg/models"
)

type UseCase interface {
	CreateProfile(profile *models.Profile) error
	DeleteProfile(profile *models.Profile) error
	GetProfile(login *string) (*models.Profile, error)
	UpdateCredentials(profile *models.Profile) error
	UpdateProfile(profile *models.Profile, name string, surname string, avatarPath string) error
	GetProfileViaID(userID uint64) (*models.Profile, error)
}
