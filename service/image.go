package service

import (
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/TruongTrongThanh/ImageServer/helper"
	"github.com/TruongTrongThanh/ImageServer/repository"
)

// Image struct for image instance
type Image struct {
	ID                       int64
	Name, Ext, URL, ThumbURL string
	CreatedAt, ModifiedAt    JSONTime
}

// JSONTime ...
type JSONTime time.Time

// MarshalJSON ...
func (j JSONTime) MarshalJSON() ([]byte, error) {
	s := time.Time(j).UTC().Format(`2006-01-02T15:04:05UTC`)
	return []byte(strconv.Quote(s)), nil
}

// NewImage ...
func NewImage(id int64, name, ext string, hasThumb bool, created, modified JSONTime) *Image {
	img := new(Image)
	img.ID = id
	img.Name = name
	img.Ext = ext
	img.CreatedAt = created
	img.ModifiedAt = modified
	img.URL = fmt.Sprintf(
		"%s%s",
		helper.GetFileServerHost(),
		img.GetLocalPath(false))
	if hasThumb {
		img.ThumbURL = fmt.Sprintf(
			"%s%s",
			helper.GetFileServerHost(),
			img.GetLocalPath(true))
	}
	return img
}

// Delete ...
func (img *Image) Delete() bool {
	repository.BeginTransaction()
	repository.DeleteImage(img.ID)
	err, thumbErr := DeleteImageFile(img)
	if err != nil || thumbErr != nil {
		repository.Rollback()
		log.Printf("Some image(s) can't delete: %d\n", img.ID)
		return false
	}
	repository.Commit()
	return true
}

// GetLocalPath ...
func (img *Image) GetLocalPath(isThumb bool) string {
	hash := helper.SimpleHash(int(img.ID))
	var name string
	if isThumb {
		name = fmt.Sprintf("t_%v", img.ID)
	} else {
		name = fmt.Sprint(img.ID)
	}
	return fmt.Sprintf("%s/%v.%s", hash, name, img.Ext)
}

// GetAllImages get image list
func GetAllImages() []Image {
	return []Image{}
}

// GetImageByID get image by using ID
func GetImageByID(_id int64) (*Image, error) {
	row := repository.GetOneImage(_id)
	var id int64
	var name, ext string
	var hasThumb bool
	var created, modified time.Time

	err := row.Scan(&id, &name, &ext, &hasThumb, &created, &modified)
	if err != nil {
		return nil, err
	}
	return NewImage(id, name, ext, hasThumb, JSONTime(created), JSONTime(modified)), nil
}

// UploadImage handle uploaded image
func UploadImage(imgFile multipart.File, header *multipart.FileHeader) (*Image, error) {
	repository.BeginTransaction()
	base, ext := helper.SplitFilename(header.Filename)
	id, created := repository.InsertImage(base, ext, false)
	err := StoreImageFile(imgFile, header, fmt.Sprint(id))
	if err != nil {
		repository.Rollback()
		return nil, err
	}
	repository.Commit()
	return NewImage(id, base, ext, false, JSONTime(created), JSONTime(created)), nil
}
