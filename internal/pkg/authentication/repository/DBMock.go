package repository

import (
	"backend/internal/pkg/models"
	"sync"
)

type AuthRepository struct {
	Users []models.User
	Mu    *sync.RWMutex
}

func NewUserRepository(mutex *sync.RWMutex) *AuthRepository {
	return &AuthRepository{
		Users: []models.User{},
		Mu:    mutex,
	}
}

type UserNotFound struct{}
type UserAlreadyExists struct{}

func (t UserNotFound) Error() string {
	return "User not found!"
}

func (t UserAlreadyExists) Error() string {
	return "User already exists!"
}

func (t *AuthRepository) CreateUser(user *models.User) error {
	success := true
	var lastID uint64 = 0
	t.Mu.RLock()
	{
		for _, val := range t.Users {
			if val.Username == user.Username {
				success = false
				break
			}
			if val.ID >= lastID {
				lastID = val.ID + 1
			}
		}
	}
	t.Mu.RUnlock()

	if !success {
		return UserAlreadyExists{}
	}

	user.ID = lastID

	t.Mu.Lock()
	{
		t.Users = append(t.Users, *user)
	}
	t.Mu.Unlock()

	return nil
}

func (t *AuthRepository) GetUser(username, password string) (*models.User, error) {
	user := new(models.User)
	success := false

	t.Mu.RLock()
	for _, val := range t.Users {
		if val.Username == username && val.Password == password {
			*user = val
			success = true
			break
		}
	}
	t.Mu.RUnlock()
	if !success {
		return user, UserNotFound{}
	}

	return user, nil
}

func (t *AuthRepository) GetUserByID(userID uint64) (*models.User, error) {
	user := new(models.User)
	success := false

	t.Mu.RLock()
	for _, val := range t.Users {
		if val.ID == userID {
			*user = val
			success = true
		}
	}
	t.Mu.RUnlock()
	if !success {
		return user, UserNotFound{}
	}

	return user, nil
}
