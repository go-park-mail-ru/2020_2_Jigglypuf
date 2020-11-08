//go:generate mockgen -source CookieRepository.go -destination mock/CookieRep_mock.go -package mock
package cookie

import (
	"backend/internal/pkg/models"
	"net/http"
)

type Repository interface {
	GetCookie(*http.Cookie) (*models.DBResponse, error)
	SetCookie(cookie *http.Cookie, userID uint64) error
	RemoveCookie(cookie *http.Cookie) error
}
