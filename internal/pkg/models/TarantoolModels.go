package models

import "net/http"

type TarantoolResponse struct{
	ID uint64
	CookieValue string
	UserID uint64
	Cookie http.Cookie
}
