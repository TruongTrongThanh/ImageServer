package routes

import (
	"net/http"
	"strings"

	"github.com/TruongTrongThanh/ImageServer/helper"
	"github.com/TruongTrongThanh/ImageServer/service"
)

// GetImages ...
func GetImages(w http.ResponseWriter, r *http.Request) {
	var res interface{}

	id := r.URL.Query().Get("id")

	if id == "" {
		res = service.GetAllImages()
	} else {
		res = service.GetImageByID(id)
	}

	helper.WriteJSON(w, res)
}

// UploadImages ...
func UploadImages(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")

	if strings.Contains(contentType, "multipart/form-data") {
		if err := r.ParseMultipartForm(1000000000); err != nil {
			helper.WriteBadRequest(w, err.Error())
			return
		}
		file, fHeader, err := r.FormFile("image")
		if err != nil {
			helper.WriteBadRequest(w, "No uploaded image(s)")
			return
		}
		service.StoreMultipartFile(file, fHeader)

		helper.WriteJSON(w, "Supported format")
	} else {
		helper.WriteBadRequest(w, "Not supported format")
	}
}
