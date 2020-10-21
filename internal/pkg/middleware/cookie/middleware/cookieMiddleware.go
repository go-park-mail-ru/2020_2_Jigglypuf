package middleware

import (
	"backend/internal/pkg/middleware/cookie/delivery"
	"backend/internal/pkg/models"
	"net/http"
)

func CookieMiddleware(t *delivery.CookieHandler, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if !t.CheckCookie(r){
			models.UnauthorizedHTTPResponse(&w)
			return
		}

		next.ServeHTTP(w,r)
	})
}
