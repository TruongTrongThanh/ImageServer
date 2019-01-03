package repository

import (
	"database/sql"
	"time"

	"github.com/TruongTrongThanh/ImageServer/service"
)

const createImageTableCmd = "CREATE TABLE Image (ID INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, Name VARCHAR(50), Ext VARCHAR(5), HasThumbnail BOOLEAN, CreatedAt TIMESTAMP, ModifiedAt TIMESTAMP)"
const insertImageCmd = "INSERT INTO Image (ID, Name, Ext, HasThumbnail, CreatedAt, ModifiedAt) VALUES (?, ?, ?, ?, ?, ?)"

// CreateImageTable ...
func CreateImageTable() {
	createIfNotExists(createImageTableCmd, "Image")
}

// InsertImage ...
func InsertImage(img service.Image) sql.Result {
	res, err := ImageDB.Exec(insertImageCmd, img.ID, img.Name, img.Ext, img.HasThumbnail, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
	return res
}
