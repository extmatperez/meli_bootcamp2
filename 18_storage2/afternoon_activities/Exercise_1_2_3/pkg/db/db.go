package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Storage_DB *sql.DB

func init() {
	data_source := "root:@tcp(localhost:3306)/db_users"
	var err error
	Storage_DB, err = sql.Open("mysql", data_source)

	if err != nil {
		panic(err)
	}
	if err := Storage_DB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database configured!")
}
