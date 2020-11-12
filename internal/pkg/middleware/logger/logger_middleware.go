package logger

import (
	"log"
	"net/http"
	"time"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("LOG [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}