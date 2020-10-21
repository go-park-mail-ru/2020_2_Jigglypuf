package cookie

import "net/http"

type Repository interface {
	GetCookie(*http.Cookie) error
	SetCookie(cookie *http.Cookie) error
	RemoveCookie(cookie *http.Cookie) error
}
