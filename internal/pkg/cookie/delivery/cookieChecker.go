package delivery

import (
	"backend/internal/pkg/cookie"
	"net/http"
)

type CookieService struct{
	dbConn cookie.CookieRepository
}

func NewCookieService(rep cookie.CookieRepository) *CookieService {
	return &CookieService{
		dbConn: rep,
	}
}


func (t *CookieService) CheckCookie(r *http.Request) bool{
	cookieValue, err := r.Cookie("session_id")
	if err != nil{
		return false
	}

	cookieErr := t.dbConn.GetCookie(cookieValue)
	if cookieErr != nil{
		return false
	}

	return true
}
