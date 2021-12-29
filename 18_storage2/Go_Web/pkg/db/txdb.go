package pkg

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/storeUsers"

	txdb.Register("txdb", "mysql", dataSource)

}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("txdb", "identificar")
	if err != nil {
		return db, db.Ping()
	}

	return db, nil
}
