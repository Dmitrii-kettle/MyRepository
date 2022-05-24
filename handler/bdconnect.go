package bdconnect

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func DbConnect() {
	db, err := sql.Open("mysql", "root:password@/productdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
