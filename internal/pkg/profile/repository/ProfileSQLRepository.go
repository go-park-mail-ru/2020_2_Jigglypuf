package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"errors"
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
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("INSERT INTO profile(`user_id`,`ProfileName`,`ProfileSurname`,`AvatarPath`)",
		profile.Login.ID, profile.Name, profile.Surname, profile.AvatarPath)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *ProfileSQLRepository) DeleteProfile(profile *models.Profile) error {
	if t.DBConnection == nil {
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("DELETE FROM profile WHERE `user_id` = ?", profile.Login.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}

func (t *ProfileSQLRepository) UpdateProfile(profile *models.Profile, name, surname, avatarPath string) error {
	if t.DBConnection == nil {
		return errors.New("no database connection")
	}

	_, DBErr := t.DBConnection.Exec("UPDATE profile SET `ProfileName` = ?, `ProfileSurname` = ?, `AvatarPath` = ? WHERE `user_id` = ?",
		name, surname, avatarPath, profile.Login.ID)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	return nil
}
func (t *ProfileSQLRepository) GetProfileViaID(userID uint64) (*models.Profile, error) {
	if t.DBConnection == nil {
		return nil, errors.New("no database connection")
	}

	resultSQL, DBErr := t.DBConnection.Query("SELECT `ProfileName`, `ProfileSurname`, `AvatarPath`, `user_id` WHERE `user_id` = ?", userID)
	if DBErr != nil {
		log.Println(DBErr)
		return nil, DBErr
	}
	rowsErr := resultSQL.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	reqProfile := new(models.Profile)
	ScanErr := resultSQL.Scan(&reqProfile.Name, &reqProfile.Surname, &reqProfile.AvatarPath, &reqProfile.Login.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return nil, ScanErr
	}

	return reqProfile, nil
}

func (t *ProfileSQLRepository) GetProfile(login *string) (*models.Profile, error) {
	if t.DBConnection == nil {
		return nil, errors.New("no database connection")
	}

	return nil, nil
}
func (t *ProfileSQLRepository) UpdateCredentials(profile *models.Profile) error {
	if t.DBConnection == nil {
		return errors.New("no database connection")
	}

	return nil
}