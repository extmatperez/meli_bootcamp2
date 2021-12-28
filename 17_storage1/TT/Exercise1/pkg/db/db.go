package db

import (
	"database/sql"
	"log"

	// Ojo con esta importacion y el _, es necesario para que no falle.
	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root@tcp(localhost:3306)/dbUser"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	if err := StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("DataBase configured")
}
