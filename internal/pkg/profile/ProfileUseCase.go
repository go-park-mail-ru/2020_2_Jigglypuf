//go:generate mockgen -source ProfileUseCase.go -destination mock/ProfileRep_mock.go -package mock
package profile

import (
	"backend/internal/pkg/models"
)

type UseCase interface {
	CreateProfile(profile *models.Profile) error
	DeleteProfile(profile *models.Profile) error
	GetProfile(login *string) (*models.Profile, error)
	UpdateCredentials(profile *models.Profile) error
	UpdateProfile(profileUserID uint64, name string, surname string, avatarPath string) error
	GetProfileViaID(userID uint64) (*models.Profile, error)
}
