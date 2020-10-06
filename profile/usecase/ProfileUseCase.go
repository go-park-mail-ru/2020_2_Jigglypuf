package usecase

import (
	"models"
	"profile"
)


type ProfileUseCase struct {
	DBConn profile.ProfileRepository
}


func NewProfileUseCase( dbConn profile.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{
		DBConn: dbConn,
	}
}


func (t *ProfileUseCase) CreateProfile( profile *models.Profile ) error {
	return t.DBConn.CreateProfile(profile)
}


func (t *ProfileUseCase) DeleteProfile( profile *models.Profile ) error {
	return t.DBConn.DeleteProfile(profile)
}


func (t *ProfileUseCase) GetProfile( login *string )( *models.Profile, error ) {
	return t.DBConn.GetProfile(login)
}


func (t *ProfileUseCase) UpdateCredentials( profile *models.Profile ) error {
	return t.DBConn.UpdateCredentials(profile)
}


func (t *ProfileUseCase) UpdatedProfile( profile *models.Profile ) error {
	return t.DBConn.UpdateProfile(profile)
}
