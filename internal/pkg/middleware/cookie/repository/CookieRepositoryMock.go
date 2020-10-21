package repository

import (
	"net/http"
	"errors"
	"sync"
)

type CookieRepository struct {
	cookieArr []http.Cookie
	mu *sync.RWMutex
}


func NewCookieRepository(mutex *sync.RWMutex) *CookieRepository {
	return &CookieRepository{
		make([]http.Cookie,3),
		mutex,
	}
}

func (t *CookieRepository) GetCookie(cookie *http.Cookie) error {
	success := false
	t.mu.RLock()
	{
		for _,val := range t.cookieArr{
			if val.Value == cookie.Value{
				success = true
			}
		}
	}
	t.mu.RUnlock()

	if !success{
		return errors.New("Cookie doesnt exist")
	}
	return nil
}

func (t *CookieRepository) SetCookie(cookie *http.Cookie)error{
	success := false
	t.mu.Lock()
	{
		t.cookieArr = append(t.cookieArr, *cookie)
		success = true
	}
	t.mu.Unlock()

	if !success{
		return errors.New("Cookie doesnt saved")
	}
	return nil

}

func (t *CookieRepository) RemoveCookie(cookie *http.Cookie)error{
	success := false
	t.mu.Lock()
	{
		for index, val := range t.cookieArr{
			if val.Value == cookie.Value{
				t.cookieArr[index] = t.cookieArr[len(t.cookieArr) - 1]
				t.cookieArr = t.cookieArr[:len(t.cookieArr) - 1]
				success = true
			}
		}
	}
	t.mu.Unlock()

	if !success{
		return errors.New("Cookie not found")
	}
	return nil
}