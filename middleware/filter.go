package middleware

import (
	"net/http"

	"github.com/TruongTrongThanh/ImageServer/helper"
)

// Filter ...
func Filter(next http.Handler, paths ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helper.StringSliceContains(paths, r.URL.Path) {
			http.NotFound(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
