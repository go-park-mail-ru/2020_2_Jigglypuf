package usecase

import (
	"backend/internal/pkg/models"
	"backend/internal/pkg/profile"
)

type ProfileUseCase struct {
	DBConn profile.Repository
}

func NewProfileUseCase(dbConn profile.Repository) *ProfileUseCase {
	return &ProfileUseCase{
		DBConn: dbConn,
	}
}

func (t *ProfileUseCase) CreateProfile(profile *models.Profile) error {
	return t.DBConn.CreateProfile(profile)
}

func (t *ProfileUseCase) DeleteProfile(profile *models.Profile) error {
	return t.DBConn.DeleteProfile(profile)
}

func (t *ProfileUseCase) GetProfile(login *string) (*models.Profile, error) {
	return t.DBConn.GetProfile(login)
}

func (t *ProfileUseCase) GetProfileViaID(userID uint64) (*models.Profile, error) {
	return t.DBConn.GetProfileViaID(userID)
}

func (t *ProfileUseCase) UpdateCredentials(profile *models.Profile) error {
	return t.DBConn.UpdateCredentials(profile)
}

func (t *ProfileUseCase) UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error {
	return t.DBConn.UpdateProfile(profile, name, surname, avatarPath)
}
