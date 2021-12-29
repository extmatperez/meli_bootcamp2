package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {
	txdb.Register("txdb", "mysql", "products_user:password@tcp(localhost:3306)/products_db")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
