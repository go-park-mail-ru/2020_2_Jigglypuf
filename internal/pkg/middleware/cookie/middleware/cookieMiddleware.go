package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/middleware/cookie/delivery"
	"net/http"
)

func CookieMiddleware(next http.Handler, cookieManager *delivery.CookieHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isAuth := true
		val, ok := cookieManager.CheckCookie(r)
		if !ok {
			isAuth = false
		}
		ctx = context.WithValue(ctx, cookie.ContextIsAuthName, isAuth)
		ctx = context.WithValue(ctx, cookie.ContextUserIDName, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
