package repository

import (
	"backend/internal/pkg/authentication"
	"net/http"
)

type CookieRepository struct{
	userRep authentication.AuthRepository
}

func NewCookieRepository(rep authentication.AuthRepository) *CookieRepository{
	return &CookieRepository{
		userRep: rep,
	}
}

func (t *CookieRepository) GetCookie(cookie *http.Cookie) error{
	_, getErr := t.userRep.GetUserViaCookie(cookie)
	return getErr
}
