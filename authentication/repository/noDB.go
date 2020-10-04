package repository

import (
	"models"
	"sync"
)

type AuthRepository struct{
	Users []models.User
	Mu *sync.RWMutex
}

func NewUserRepository(mutex *sync.RWMutex) *AuthRepository{
	return &AuthRepository{
		Users: []models.User{},
		Mu: mutex,
	}
}

type UserNotFound struct{}
type UserAlreadyExists struct{}

func (t UserNotFound) Error() string{
	return "User not found!"
}

func (t UserAlreadyExists) Error() string{
	return "User already exists!"
}


func (t *AuthRepository) CreateUser(user *models.User) error{
	success := true

	t.Mu.RLock()
	{
		for _, val := range t.Users{
			if val.Username == user.Username{
				success = false
			}
		}
	}
	t.Mu.RUnlock()

	if !success{
		return UserAlreadyExists{}
	}

	t.Mu.Lock()
	{
		t.Users = append(t.Users, *user)
		//t.Cookies[*user] = *cookieValue
	}
	t.Mu.Unlock()

	return nil
}


func (t *AuthRepository) GetUser(username, password string) (*models.User, error){
	user := new(models.User)
	success := false

	for _,val := range t.Users{
		if val.Username == username && val.Password == password{
			*user = val
			success = true
			break
		}
	}

	if !success{
		return user, UserNotFound{}
	}

	return user, nil
}