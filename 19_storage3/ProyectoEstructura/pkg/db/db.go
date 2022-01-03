package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/dbproductos"

	var err error

	StorageDB, err = sql.Open("mysql", dataSource)

	if err != nil {
		fmt.Println("aqa")
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		fmt.Println("aqa ping")

		panic(err)
	}

	log.Println("Database Configured")
}
