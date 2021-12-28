package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

func init() {
	err := godotenv.Load("../../cmd/server/.env")

	if err != nil {
		log.Fatal(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbServer := os.Getenv("DB_SERVER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dataSource := dbUser + ":" + dbPassword + "@tcp(" + dbServer + ":" + dbHost + ")/" + dbName

	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB ready")
}
