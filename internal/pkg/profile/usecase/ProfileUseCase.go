package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/profile"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
)

type ProfileUseCase struct {
	sanitizer  *bluemonday.Policy
	repository profile.Repository
}

func NewProfileUseCase(rep profile.Repository) *ProfileUseCase {
	return &ProfileUseCase{
		sanitizer:  bluemonday.UGCPolicy(),
		repository: rep,
	}
}

func (t *ProfileUseCase) CreateProfile(reqProfile *models.Profile) error {
	if reqProfile.Surname == "" || reqProfile.Name == "" {
		return models.ErrFooIncorrectInputInfo
	}
	if reqProfile.AvatarPath == "" {
		reqProfile.AvatarPath = profile.NoAvatarImage
	}
	utils.SanitizeInput(t.sanitizer, &reqProfile.Name, &reqProfile.Surname, &reqProfile.AvatarPath)
	return t.repository.CreateProfile(reqProfile)
}

func (t *ProfileUseCase) DeleteProfile(profile *models.Profile) error {
	return t.repository.DeleteProfile(profile)
}

func (t *ProfileUseCase) GetProfile(login *string) (*models.Profile, error) {
	return t.repository.GetProfile(login)
}

func (t *ProfileUseCase) GetProfileViaID(userID uint64) (*models.Profile, error) {
	return t.repository.GetProfileViaID(userID)
}

func (t *ProfileUseCase) UpdateCredentials(profile *models.Profile) error {
	return t.repository.UpdateCredentials(profile)
}

func (t *ProfileUseCase) UpdateProfile(profileUserID uint64, name, surname, avatarPath string) error {
	if name == "" && surname == "" && avatarPath == "" {
		return models.ErrFooIncorrectInputInfo
	}
	profileInput, profileError := t.GetProfileViaID(profileUserID)
	if profileError != nil {
		return models.ErrFooNoAuthorization
	}

	utils.SanitizeInput(t.sanitizer, &name, &surname, &avatarPath)
	if name == "" {
		name = profileInput.Name
	}
	if surname == "" {
		surname = profileInput.Surname
	}
	if avatarPath == "" {
		avatarPath = profileInput.AvatarPath
	}
	return t.repository.UpdateProfile(profileInput, name, surname, avatarPath)
}
