package middleware

import (
	"log"
	"net/http"

	"github.com/TruongTrongThanh/ImageServer/helper"
)

// Common ...
func Common(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Global Headers
		w.Header().Add("Content-Type", "application/json")

		// Logging
		userAgent := helper.GetOrDefault(r.Header.Get("User-Agent"), "Unknown user-agent")
		log.Println(userAgent.(string) + " - " + r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
