package repository

import (
	"models"
	"net/http"
	"sync"
	"time"
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

	t.Mu.RLock()
	for _,val := range t.Users{
		if val.Username == username && val.Password == password{
			*user = val
			success = true
			break
		}
	}
	t.Mu.RUnlock()
	if !success{
		return user, UserNotFound{}
	}

	return user, nil
}

func (t *AuthRepository) GetUserViaCookie(cookie *http.Cookie)(*models.User, error){
	user := new(models.User)
	success := false

	t.Mu.RLock()
	for _,val := range t.Users{
		if cookie != nil && cookie.Value == val.Cookie.Value && time.Now().Before(val.Cookie.Expires){
			*user = val
			success = true
		}
	}
	t.Mu.RUnlock()
	if !success{
		return user, UserNotFound{}
	}

	return user, nil
}

func (t *AuthRepository) SetCookie(user *models.User,cookieValue *http.Cookie) bool{
	userIndex := 0
	t.Mu.RLock()
	{
		for index, val := range t.Users{
			if val.Username == user.Username{
				userIndex = index
			}
		}
	}
	t.Mu.RUnlock()

	t.Mu.RLock()
	{
		t.Users[userIndex].Cookie = *cookieValue
	}
	t.Mu.RUnlock()

	return true
}