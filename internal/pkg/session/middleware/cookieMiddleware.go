package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session"
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/session/delivery"
	"log"
	"net/http"
	"time"
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
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		if cookieValue, ok := r.Context().Value(session.ContextCookieName).(http.Cookie); ok{
			log.Println("setting cookie", cookieValue)
			http.SetCookie(w, &cookieValue)
			if time.Now().After(cookieValue.Expires){
				cookieManager.RemoveCookie(&cookieValue)
				return
			}
			userID, okUser := r.Context().Value(session.ContextUserIDName).(uint64)
			if okUser{
				cookieManager.SetCookie(&cookieValue, userID)
			}
		}
	})
}
