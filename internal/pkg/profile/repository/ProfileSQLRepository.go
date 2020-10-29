package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"log"
)

type ProfileSQLRepository struct {
	DBConnection *sql.DB
}

func NewProfileSQLRepository(connection *sql.DB) *ProfileSQLRepository {
	return &ProfileSQLRepository{
		connection,
	}
}

func (t *ProfileSQLRepository) CreateProfile(profile *models.Profile) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("INSERT INTO profile(user_id,ProfileName,ProfileSurname,AvatarPath) VALUES ($1,$2,$3,$4)",
		profile.Login.ID, profile.Name, profile.Surname, profile.AvatarPath)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *ProfileSQLRepository) DeleteProfile(profile *models.Profile) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("DELETE FROM profile WHERE user_id = $1", profile.Login.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *ProfileSQLRepository) UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	_, DBErr := t.DBConnection.Exec("UPDATE profile SET ProfileName = $1, ProfileSurname = $2, AvatarPath = $3 WHERE user_id = $4",
		name, surname, avatarPath, profile.Login.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}
func (t *ProfileSQLRepository) GetProfileViaID(userID uint64) (*models.Profile, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	resultSQL := t.DBConnection.QueryRow("SELECT ProfileName, ProfileSurname, AvatarPath, user_id FROM profile WHERE user_id = $1", userID)
	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	reqProfile := new(models.Profile)
	reqProfile.Login = new(models.User)
	ScanErr := resultSQL.Scan(&reqProfile.Name, &reqProfile.Surname, &reqProfile.AvatarPath, &reqProfile.Login.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return nil, ScanErr
	}

	return reqProfile, nil
}

func (t *ProfileSQLRepository) GetProfile(login *string) (*models.Profile, error) {
	if t.DBConnection == nil {
		return nil, models.ErrFooNoDBConnection
	}

	return nil, nil
}
func (t *ProfileSQLRepository) UpdateCredentials(profile *models.Profile) error {
	if t.DBConnection == nil {
		return models.ErrFooNoDBConnection
	}

	return nil
}
