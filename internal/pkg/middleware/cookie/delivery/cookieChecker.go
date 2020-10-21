package delivery

import (
	"backend/internal/pkg/middleware/cookie"
	"net/http"
)

type CookieHandler struct {
	dbConn cookie.Repository
}

func NewCookieHandler(rep cookie.Repository) *CookieHandler {
	return &CookieHandler{
		dbConn: rep,
	}
}

func (t *CookieHandler) CheckCookie(r *http.Request) bool {
	cookieValue, err := r.Cookie("session_id")
	if err != nil {
		return false
	}

	cookieErr := t.dbConn.GetCookie(cookieValue)
	return cookieErr == nil
}
