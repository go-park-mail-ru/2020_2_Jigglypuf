package monitoring

import (
	"github.com/go-park-mail-ru/2020_2_Jigglypuf/internal/pkg/promconfig"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"time"
)


var URLResponseLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_url_duration",
		Help:    "Latency of response duration",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	},
	[]string{"url","handler","method","status"},
)


func init(){
	prometheus.MustRegister(URLResponseLatency)
}


func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := promconfig.StatusErr
		handler := "no_handler"
		start := time.Now()
		next.ServeHTTP(w, r)
		handlerName := w.Header().Get(promconfig.HandlerNameID)
		if handlerName != "" {
			log.Println("Handler name was found")
			handler = handlerName
			handlerStatus := w.Header().Get(promconfig.StatusNameID)
			if handlerStatus != "" {
				log.Println("Status name was found", handlerStatus)
				status = handlerStatus
				w.Header().Del(promconfig.StatusNameID)
			}
			w.Header().Del(promconfig.HandlerNameID)
		}
		log.Println(r.URL.Path,handler, r.Method, status)
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			URLResponseLatency.WithLabelValues(r.URL.Path,handler, r.Method, status).Observe(v)
		}))
		defer func(){
			timer.ObserveDuration()
		}()

		log.Printf("LOG [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}
