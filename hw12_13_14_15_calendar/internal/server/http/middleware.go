package internalhttp

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("%s [%v] %s %s HTTP/1.1 200 %v \"%s\"", r.RemoteAddr, end, r.Method, time.Since(start), r.URL, r.UserAgent())
	})
}
