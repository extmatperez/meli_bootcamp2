package db

import (
	"database/sql"
	"log"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/productos"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err := StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("conexion creada")

}
