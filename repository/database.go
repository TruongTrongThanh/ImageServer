package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	// Mysql ext
	_ "github.com/go-sql-driver/mysql"
)

//ImageDB is Global DB instance
var ImageDB *sql.DB

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
