package repository

import (
	"database/sql"
	"time"
)

const createImageTableCmd = "CREATE TABLE Image (ID INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, Name VARCHAR(50), Ext VARCHAR(5), HasThumbnail BOOLEAN, CreatedAt TIMESTAMP, ModifiedAt TIMESTAMP)"
const insertImageCmd = "INSERT INTO Image (Name, Ext, HasThumbnail, CreatedAt, ModifiedAt) VALUES (?, ?, ?, ?, ?)"
const getOneImageCmd = "SELECT * FROM Image WHERE ID = ?"
const deleteImageCmd = "DELETE FROM Image WHERE ID = ?"

// CreateImageTable ...
func CreateImageTable() {
	createIfNotExists(createImageTableCmd, "Image")
}

// InsertImage ...
func InsertImage(name string, ext string, hasThumb bool) (int64, time.Time) {
	created := time.Now()
	res, err := ImageTx.Exec(insertImageCmd, name, ext, hasThumb, created, created)
	if err != nil {
		panic(err)
	}
	id, _ := res.LastInsertId()
	return id, created
}

// DeleteImage ...
func DeleteImage(id int64) {
	_, err := ImageTx.Exec(deleteImageCmd, id)
	if err != nil {
		panic(err)
	}
}

// GetOneImage ...
func GetOneImage(id int64) *sql.Row {
	return ImageDB.QueryRow(getOneImageCmd, id)
}
