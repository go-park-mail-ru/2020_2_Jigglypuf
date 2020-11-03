package repository

import (
	"backend/internal/pkg/models"
	"database/sql"
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
	ScanErr := t.DBConn.QueryRow("INSERT INTO users (Login, password) VALUES ($1,$2) RETURNING ID", user.Login, user.Password).Scan(&user.ID)
	if ScanErr != nil {
		log.Println(ScanErr)
		return models.ErrFooInternalDBErr
	}
	return nil
}

func (t *AuthSQLRepository) GetUser(Login string) (*models.User, error) {
	if t.DBConn == nil {
		return nil, models.ErrFooNoDBConnection
	}

	requiredUser := new(models.User)
	result := t.DBConn.QueryRow("SELECT id, Login, password FROM users WHERE Login = $1", Login)
	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultErr := result.Scan(&requiredUser.ID, &requiredUser.Login, &requiredUser.Password)
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
	result := t.DBConn.QueryRow("SELECT id, Login, password FROM users WHERE id = $1", userID)

	rowsErr := result.Err()
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	resultErr := result.Scan(&requiredUser.ID, &requiredUser.Login, &requiredUser.Password)
	if resultErr != nil {
		log.Println(resultErr)
		return nil, resultErr
	}

	return requiredUser, nil
}
