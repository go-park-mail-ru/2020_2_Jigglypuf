package repository

import (
	"backend/internal/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type CookieArrayTable struct {
	cookie http.Cookie
	userID uint64
}

type CookieRepository struct {
	cookieArr []CookieArrayTable
	mu        *sync.RWMutex
}

func NewCookieRepository(mutex *sync.RWMutex) *CookieRepository {
	return &CookieRepository{
		make([]CookieArrayTable, 3),
		mutex,
	}
}

func (t *CookieRepository) GetCookie(cookie *http.Cookie) (*models.TarantoolResponse, error) {
	success := false
	var userID uint64 = 0
	t.mu.RLock()
	{
		for _, val := range t.cookieArr {
			if val.cookie.Value == cookie.Value {
				success = true
				userID = val.userID
			}
		}
	}
	t.mu.RUnlock()

	if !success {
		return nil, errors.New("cookie doesnt exist")
	}
	fmt.Print(userID)
	return nil, nil
}

func (t *CookieRepository) SetCookie(cookie *http.Cookie, userID uint64) error {
	success := false
	t.mu.Lock()
	{
		t.cookieArr = append(t.cookieArr, CookieArrayTable{*cookie, userID})
		success = true
	}
	t.mu.Unlock()

	if !success {
		return errors.New("cookie doesnt saved")
	}
	return nil
}

func (t *CookieRepository) RemoveCookie(cookie *http.Cookie) error {
	success := false
	t.mu.Lock()
	{
		for index, val := range t.cookieArr {
			if val.cookie.Value == cookie.Value {
				t.cookieArr[index] = t.cookieArr[len(t.cookieArr)-1]
				t.cookieArr = t.cookieArr[:len(t.cookieArr)-1]
				success = true
			}
		}
	}
	t.mu.Unlock()

	if !success {
		return errors.New("cookie not found")
	}
	return nil
}
