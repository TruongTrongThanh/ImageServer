package middlewares

import (
	"log"
	"net/http"
)

// Auth ...
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth handle...")
		next.ServeHTTP(w, r)
	})
}
