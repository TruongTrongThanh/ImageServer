package routes

import (
	"net/http"

	"github.com/TruongTrongThanh/ImageServer/helper"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Owner   string
		Version string
	}{
		"Thanh",
		"0.1.0",
	}
	helper.WriteJSON(w, info)
}

// Debug ...
func Debug(w http.ResponseWriter, r *http.Request) {
	helper.WriteJSON(w, "debug path")
}
