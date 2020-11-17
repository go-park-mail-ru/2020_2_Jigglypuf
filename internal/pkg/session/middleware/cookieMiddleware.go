package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/delivery"
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
		ctx = context.WithValue(ctx, session.ContextIsAuthName, isAuth)
		ctx = context.WithValue(ctx, session.ContextUserIDName, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
