package models

import "net/http"

type DBResponse struct{
	ID uint64
	CookieValue string
	UserID uint64
	Cookie http.Cookie
}
