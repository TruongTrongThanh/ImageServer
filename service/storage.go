package service

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

// StoreMultipartFile ...
func StoreMultipartFile(f multipart.File, fHeader *multipart.FileHeader) error {
	dir := os.Getenv("StoredPath")
	pathname := path.Join(dir, fHeader.Filename)
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
