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
		return errors.New("no database connection")
	}
	result, DBErr := t.DBConn.Exec("INSERT INTO users (username, password) VALUES ($1,$2)", user.Username, user.Password)
	if DBErr != nil {
		log.Println(DBErr)
		return DBErr
	}

	ID, IDErr := result.LastInsertId()
	if IDErr != nil {
		log.Println(DBErr)
		return IDErr
	}

	user.ID = uint64(ID)
	return nil
}

func (t *AuthSQLRepository) GetUser(username string, hashPassword string) (*models.User, error) {
	if t.DBConn == nil {
		return nil, errors.New("no database connection")
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
		return nil, errors.New("no database connection")
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
