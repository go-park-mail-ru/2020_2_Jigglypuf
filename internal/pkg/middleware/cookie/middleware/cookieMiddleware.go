package middleware

import (
	CookieService "backend/internal/app/cookieserver"
	"backend/internal/pkg/middleware/cookie"
	"context"
	"net/http"
)

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isAuth := true
		if !CookieService.CookieManager.CookieDelivery.CheckCookie(r) {
			isAuth = false
		}
		ctx = context.WithValue(ctx, cookie.ContextIsAuthName, isAuth)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
