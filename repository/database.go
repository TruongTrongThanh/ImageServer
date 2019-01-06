package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	// Mysql ext
	_ "github.com/go-sql-driver/mysql"
)

// ImageDB is Global DB instance
var ImageDB *sql.DB

// ImageTx is Global DB transcation instance
var ImageTx *sql.Tx

// Connect ...
func Connect() {
	db, err := sql.Open("mysql", os.Getenv("DatabaseString"))
	if err != nil {
		panic(err)
	}
	ImageDB = db
}

func createIfNotExists(query string, tableName string) {
	_, err := ImageDB.Exec(query)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1050") {
			fmt.Printf("%s table already exists\n", tableName)
		} else {
			panic(err)
		}
	}
}

// BeginTransaction ...
func BeginTransaction() {
	var err error
	ImageTx, err = ImageDB.Begin()
	if err != nil {
		panic(err)
	}
}

// Commit ...
func Commit() {
	err := ImageTx.Commit()
	if err != nil {
		panic(err)
	}
}

// Rollback ...
func Rollback() {
	err := ImageTx.Rollback()
	if err != nil {
		panic(err)
	}
}
