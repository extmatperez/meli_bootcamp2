package internal

import (
	"database/sql"
	"log"

	db "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/database"
	models "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/models"
)

type RepositorySql interface {
	Store(user models.User) (models.User, error)
}

type repositorySql struct{}

func NewRepo() RepositorySql {
	return &repositorySql{}
}

func (s *repositorySql) Store(user models.User) (models.User, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO users(name, last_name, email, age, height, active, created) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
		return models.User{}, err
	}

	var result sql.Result
	result, err = stmt.Exec(user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created)
	if err != nil {
		return models.User{}, err
	}

	newId, _ := result.LastInsertId()
	user.ID = int(newId)
	return user, nil
}
