package cookie

import "net/http"

type Repository interface {
	GetCookie(*http.Cookie) (uint64, error)
	SetCookie(cookie *http.Cookie, userID uint64) error
	RemoveCookie(cookie *http.Cookie) error
}
