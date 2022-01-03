package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/transactiondb"
	txdb.Register("txdb", "mysql", dataSource)
}
func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())
	if err == nil {
		return db, db.Ping()
	}
	return db, err
}
