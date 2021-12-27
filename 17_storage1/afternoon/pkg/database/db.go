package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init() {
	datasource := "test_db_user:Test_DB#123@/bootcamp"

	var err error

	StorageDB, err = sql.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

}
