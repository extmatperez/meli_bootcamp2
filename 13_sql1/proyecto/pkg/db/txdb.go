package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/dbtransactions"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())
	if err == nil {
		return db, db.Ping()
	}
	return db, err
}
