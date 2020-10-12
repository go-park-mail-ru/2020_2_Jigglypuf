package profile

import (
	"backend/internal/pkg/models"
	"net/http"
)

type Repository interface {
	CreateProfile(profile *models.Profile) error
	DeleteProfile(profile *models.Profile) error
	GetProfile(login *string) (*models.Profile, error)
	UpdateCredentials(profile *models.Profile) error
	UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error
	GetProfileViaCookie(cookie *http.Cookie) (*models.Profile, error)
}
