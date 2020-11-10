//go:generate mockgen -source CookieRepository.go -destination mock/CookieRep_mock.go -package mock
package cookie

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/models"
	"net/http"
)

type Repository interface {
	GetCookie(*http.Cookie) (*models.DBResponse, error)
	SetCookie(cookie *http.Cookie, userID uint64) error
	RemoveCookie(cookie *http.Cookie) error
}
