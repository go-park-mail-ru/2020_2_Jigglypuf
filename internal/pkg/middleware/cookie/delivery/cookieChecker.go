package delivery

import (
	"backend/internal/pkg/middleware/cookie"
	"net/http"
	"time"
)

type CookieHandler struct {
	dbConn cookie.Repository
}

func NewCookieHandler(rep cookie.Repository) *CookieHandler {
	return &CookieHandler{
		dbConn: rep,
	}
}

func (t *CookieHandler) CheckCookie(r *http.Request) (uint64, bool) {
	cookieValue, err := r.Cookie("session_id")
	if err != nil {
		return 0, false
	}

	value, cookieErr := t.dbConn.GetCookie(cookieValue)
	if cookieErr == nil{
		if time.Now().Before(value.Cookie.Expires){
			return 0, false
		}
		return value.UserID, true
	}
	return 0, false
}
