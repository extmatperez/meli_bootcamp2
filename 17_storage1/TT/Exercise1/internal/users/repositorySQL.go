package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/Exercise1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/Exercise1/pkg/db"
)

type RepositorySQL interface {
	Store(user models.User) (models.User, error)
	GetOne(id int) models.User
	Update(user models.User) (models.User, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(user models.User) (models.User, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO users(first_name, last_name, email, age, height, active, cration_date) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate)
	if err != nil {
		return models.User{}, err
	}
	idCreated, _ := result.LastInsertId()
	user.ID = int(idCreated)

	return user, nil
}

func (r *repositorySQL) GetOne(id int) models.User {
	db := db.StorageDB
	var userRead models.User
	rows, err := db.Query("SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return userRead
	}
	for rows.Next() {
		err := rows.Scan(&userRead.ID, &userRead.FirstName, &userRead.LastName, &userRead.Email, &userRead.Age, &userRead.Height, &userRead.Active, &userRead.CrationDate)
		if err != nil {
			log.Fatal(err)
			return userRead
		}
		// // En caso de querer devolver mas de uno, por ejemplo un getFirstName
		// myUsers = append(myUsers, userRead)
	}
	return userRead
}

func (r *repositorySQL) Update(user models.User) (models.User, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE users SET first_name = ?, last_name = ?, email =?, age=?, height =?, active =?, cration_date=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate, user.ID)
	if err != nil {
		return models.User{}, err
	}
	partUpdate, _ := result.RowsAffected()
	if partUpdate == 0 {
		return models.User{}, errors.New("User not found")
		// return models.User{FirstName: user.FirstName, LastName: user.LastName, Age: user.Age, Height: user.Height, Active: user.Active, CrationDate: user.CrationDate}, nil
	}
	return user, nil
}
