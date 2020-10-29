package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type AuthSQLRepository struct {
	DBConn *sql.DB
}

func NewAuthSQLRepository(connection *sql.DB) *AuthSQLRepository {
	log.Println("creating new auth postgresql repository")
	return &AuthSQLRepository{
		connection,
	}
}

func (t *AuthSQLRepository) CreateUser(user *models.User) error {
	if t.DBConn == nil {
		return models.ErrFooNoDBConnection
	}
	ScanErr := t.DBConn.QueryRow("INSERT INTO users (username, password) VALUES ($1,$2) RETURNING ID", user.Username, user.Password).Scan(&user.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return errors.New("service not available")
	}
	return nil
}

func (t *AuthSQLRepository) GetUser(username string, hashPassword string) (*models.User, error) {
	if t.DBConn == nil {
		return nil, models.ErrFooNoDBConnection
	}

	requiredUser := new(models.User)
	result := t.DBConn.QueryRow("SELECT id, username, password FROM users WHERE username = $1 AND password = $2", username, hashPassword)
	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultErr := result.Scan(&requiredUser.ID, &requiredUser.Username, &requiredUser.Password)
	if resultErr != nil {
		log.Println(resultErr)
		return nil, resultErr
	}

	return requiredUser, nil
}

func (t *AuthSQLRepository) GetUserByID(userID uint64) (*models.User, error) {
	if t.DBConn == nil {
		return nil, models.ErrFooNoDBConnection
	}

	requiredUser := new(models.User)
	result := t.DBConn.QueryRow("SELECT id, username, password FROM users WHERE id = $1", userID)

	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultErr := result.Scan(&requiredUser.ID, &requiredUser.Username, &requiredUser.Password)
	if resultErr != nil {
		log.Println(resultErr)
		return nil, resultErr
	}

	return requiredUser, nil
}
