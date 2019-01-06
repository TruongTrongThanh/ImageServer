package service

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/TruongTrongThanh/ImageServer/helper"
)

// StoreImageFile ...
func StoreImageFile(f multipart.File, fHeader *multipart.FileHeader, name string) error {
	dir := os.Getenv("StoredPath")
	num, _ := strconv.Atoi(name)
	subDir := helper.SimpleHash(num)
	dirErr := os.Mkdir(dir+"/"+subDir, os.ModeDir)
	if dirErr != nil {
		return dirErr
	}
	pathname := path.Join(dir, subDir, name+path.Ext(fHeader.Filename))
	file, createErr := os.Create(pathname)
	if createErr != nil {
		return createErr
	}
	defer file.Close()

	_, copyErr := io.Copy(file, f)
	if copyErr != nil {
		return copyErr
	}
	return nil
}

// DeleteImageFile ...
func DeleteImageFile(img *Image) (error, error) {
	dir := os.Getenv("StoredPath")
	var path, thumbPath string
	var err, thumbErr error

	path = img.GetLocalPath(false)
	err = os.Remove(dir + "/" + path)

	if img.ThumbURL != "" {
		thumbPath = img.GetLocalPath(true)
		thumbErr = os.Remove(dir + "/" + thumbPath)
	}
	return err, thumbErr
}
