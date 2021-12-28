package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
)

// Esto hace que por ejemplo luego de cada insert se tome como una base de datos transaccional, y se realice rollback.

func init() {
	dataSource := "root@tcp(localhost:3306)/dbpayments"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "identificador")
	if err != nil {
		return db, err
	}
	return db, nil
}
