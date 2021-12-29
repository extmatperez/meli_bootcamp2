package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
)

func init() {

	dataSource := "root:@tcp(localhost:3306)/productos_db"

	txdb.Register("txdb", "mysql", dataSource)
}

func InitDb() (*sql.DB, error) {

	db, err := sql.Open("txdb", "identificar")

	if err != nil {
		return db, err
	}

	return db, db.Ping()
}
