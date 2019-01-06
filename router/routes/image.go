package routes

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TruongTrongThanh/ImageServer/helper"
	"github.com/TruongTrongThanh/ImageServer/service"
)

// GetImages ...
func GetImages(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var err error

	id := r.URL.Query().Get("id")

	if id == "" {
		res = service.GetAllImages()
	} else {
		numID, parseErr := strconv.Atoi(id)
		if parseErr != nil {
			helper.WriteBadRequest(w, "ID format is invalid")
			return
		}
		res, err = service.GetImageByID(int64(numID))
	}
	if err != nil {
		if err == sql.ErrNoRows {
			helper.WriteNotFound(w, "Image not found")
		} else {
			helper.WriteError(w, "")
			log.Println(err)
		}
	} else {
		helper.WriteOK(w, res)
	}
}

// UploadImages ...
func UploadImages(w http.ResponseWriter, r *http.Request) {
	if helper.IsMultipartFormData(*r) {
		// Get max size
		maxSize, envErr := strconv.Atoi(os.Getenv("MaxMultipartFormDataSize"))
		if envErr != nil {
			helper.WriteError(w, "")
			log.Println(envErr)
			return
		}
		// Parse form
		r.ParseMultipartForm(int64(maxSize))
		file, fHeader, err := r.FormFile("image")
		if err != nil {
			helper.WriteBadRequest(w, "No uploaded image(s)")
			return
		}
		// Handle upload
		img, upErr := service.UploadImage(file, fHeader)
		if upErr != nil {
			helper.WriteError(w, "")
			log.Println(upErr)
			return
		}

		helper.WriteOK(w, img)
	} else {
		helper.WriteBadRequest(w, "Not supported format")
	}
}

// DeleteImage ...
func DeleteImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		helper.WriteBadRequest(w, "ID is empty")
	} else {
		numID, parseErr := strconv.Atoi(id)
		if parseErr != nil {
			helper.WriteBadRequest(w, "ID format is invalid")
		}
		img, getErr := service.GetImageByID(int64(numID))
		if getErr != nil {
			if getErr == sql.ErrNoRows {
				helper.WriteNotFound(w, "Image not found")
			} else {
				helper.WriteError(w, "")
				log.Println(getErr)
				return
			}
		}

		if img.Delete() {
			helper.WriteOK(w, "Delete success")
		} else {
			helper.WriteError(w, "")
		}
	}
}
