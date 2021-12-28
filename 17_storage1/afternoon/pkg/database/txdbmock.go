package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	txdb.Register("txdb", "mysql", "test_db_user:Test_DB#123@/bootcamp")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "prueba")

	if err != nil {
		return db, db.Ping()
	}

	return db, err
}
