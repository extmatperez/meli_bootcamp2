package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	dataSource := "products_user:password@tcp(localhost:3306)/products_db"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
}
