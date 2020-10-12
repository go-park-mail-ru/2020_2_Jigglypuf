package cookie

import "net/http"

type Repository interface {
	GetCookie(*http.Cookie) error
}
