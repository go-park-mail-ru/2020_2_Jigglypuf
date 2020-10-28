package cookie

import (
	"backend/internal/pkg/models"
	"net/http"
)

type Repository interface {
	GetCookie(*http.Cookie) (*models.TarantoolResponse, error)
	SetCookie(cookie *http.Cookie, userID uint64) error
	RemoveCookie(cookie *http.Cookie) error
}
