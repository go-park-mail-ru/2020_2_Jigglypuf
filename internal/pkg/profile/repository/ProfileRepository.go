package repository

import (
	"backend/internal/pkg/authentication"
	"backend/internal/pkg/models"
	"sync"
)

type ProfileRepository struct {
	Profiles       []models.Profile
	Mu             *sync.RWMutex
	UserRepository authentication.AuthRepository
}

func NewProfileRepository(mutex *sync.RWMutex, repository authentication.AuthRepository) *ProfileRepository {
	return &ProfileRepository{
		Profiles:       []models.Profile{},
		Mu:             mutex,
		UserRepository: repository,
	}
}

type ProfileNotFound struct{}

type ProfileAlreadyExists struct{}

func (t ProfileNotFound) Error() string {
	return "Profile not found!"
}

func (t ProfileAlreadyExists) Error() string {
	return "Profile already exists!"
}

// TODO уточнить тк у нас связка вместе с CreateUser
func (t *ProfileRepository) CreateProfile(profile *models.Profile) error {
	success := true

	t.Mu.RLock()
	{
		for _, value := range t.Profiles {
			if value.Login.Username == profile.Login.Username {
				success = false
			}
		}
	}
	t.Mu.RUnlock()

	if !success {
		return ProfileAlreadyExists{}
	}

	t.Mu.Lock()
	{
		t.Profiles = append(t.Profiles, *profile)
	}
	t.Mu.Unlock()

	return nil
}

// TODO DeleteProfile (profile *models.Profile) error
func (t *ProfileRepository) DeleteProfile(profile *models.Profile) error {
	return nil
}

func (t *ProfileRepository) GetProfile(login *string) (*models.Profile, error) {
	profile := new(models.Profile)
	success := false
	t.Mu.RLock()
	{
		for _, value := range t.Profiles {
			if value.Login.Username == *login {
				*profile = value
				success = true
				break
			}
		}
	}
	t.Mu.RUnlock()

	if !success {
		return profile, ProfileNotFound{}
	}

	return profile, nil
}

func (t *ProfileRepository) GetProfileViaID(userID uint64) (*models.Profile, error) {
	user, err := t.UserRepository.GetUserByID(userID)
	profile := new(models.Profile)
	if err != nil {
		return profile, err
	}

	success := false
	t.Mu.RLock()
	{
		for _, value := range t.Profiles {
			if value.Login.Username == user.Username {
				*profile = value
				success = true
				break
			}
		}
	}
	t.Mu.RUnlock()

	if !success {
		profile.Login = user
		t.Mu.Lock()
		{
			t.Profiles = append(t.Profiles, *profile)
		}
		t.Mu.Unlock()
	}

	return profile, nil
}

// TODO UpdateCredentials(profile *models.Profile) error
func (t *ProfileRepository) UpdateCredentials(profile *models.Profile) error {
	if _, err := t.GetProfile(&profile.Login.Username); err != nil {
		return ProfileNotFound{}
	}

	return nil
}

func (t *ProfileRepository) UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error {
	if _, err := t.GetProfile(&profile.Login.Username); err != nil {
		return ProfileNotFound{}
	}

	if name != "" {
		profile.Name = name
	}

	if surname != "" {
		profile.Surname = surname
	}

	if avatarPath != "" {
		profile.AvatarPath = avatarPath
	}

	profileIndex := 0
	t.Mu.RLock()
	{
		for index, val := range t.Profiles {
			if val.Login.Username == profile.Login.Username {
				profileIndex = index
			}
		}
	}
	t.Mu.RUnlock()

	t.Mu.Lock()
	{
		t.Profiles[profileIndex] = *profile
	}
	t.Mu.Unlock()

	return nil
}
