package middlewares

import (
	"log"
	"net/http"
)

// Filter ...
func Filter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Filter handle...")
		next.ServeHTTP(w, r)
	})
}
