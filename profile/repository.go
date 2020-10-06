package profile

import (
	"models"
)

type ProfileRepository interface {
	CreateProfile( profile *models.Profile ) error
	DeleteProfile( profile *models.Profile ) error
	GetProfile( login *string ) ( *models.Profile, error )
	UpdateCredentials( profile *models.Profile ) error
	UpdateProfile( profile *models.Profile ) error
}
