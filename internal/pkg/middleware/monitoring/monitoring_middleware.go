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
		handlerName := r.Context().Value(promconfig.HandlerNameID)
		if handlerName != nil {
			log.Println("Handler name was found")
			handler = handlerName.(string)
			handlerStatus := r.Context().Value(promconfig.StatusNameID)
			if handlerStatus != nil{
				status = handlerStatus.(string)
			}
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
