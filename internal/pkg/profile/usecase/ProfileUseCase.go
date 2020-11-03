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

func (t *ProfileUseCase) CreateProfile(reqProfile *models.Profile) error {
	if reqProfile.Surname == "" || reqProfile.Name == ""{
		return models.ErrFooIncorrectInputInfo
	}
	if reqProfile.AvatarPath == ""{
		reqProfile.AvatarPath = profile.NoAvatarImage
	}
	return t.DBConn.CreateProfile(reqProfile)
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
	if name == "" && surname == "" && avatarPath == ""{
		return models.ErrFooIncorrectInputInfo
	}

	if name == "" {
		name = profile.Name
	}
	if surname == "" {
		surname = profile.Surname
	}
	if avatarPath == "" {
		avatarPath = profile.AvatarPath
	}
	return t.DBConn.UpdateProfile(profile, name, surname, avatarPath)
}
