package cookie

import "net/http"

type CookieRepository interface{
	GetCookie(*http.Cookie) error
}
