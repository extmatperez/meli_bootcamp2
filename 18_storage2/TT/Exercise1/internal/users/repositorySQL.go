package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/Exercise1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/Exercise1/pkg/db"
)

type RepositorySQL interface {
	Store(user models.User) (models.User, error)
	GetOne(id int) models.User
	Update(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(id int) error
	GetOneWithContext(ctx context.Context, id int) (models.User, error)
	GetFullData() ([]models.User, error)
	UpdateWithContext(ctx context.Context, user models.User) (models.User, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(user models.User) (models.User, error) {
	db := db.StorageDB

	myQuery := "INSERT INTO users(first_name, last_name, email, age, height, active, cration_date) VALUES(?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(myQuery)
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
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users WHERE id = ?"
	rows, err := db.Query(myQuery, id)
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
	}
	return userRead
}

func (r *repositorySQL) Update(user models.User) (models.User, error) {
	db := db.StorageDB
	myQuery := "UPDATE users SET first_name = ?, last_name = ?, email =?, age=?, height =?, active =?, cration_date=? WHERE id=?"
	stmt, err := db.Prepare(myQuery)
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

func (r *repositorySQL) GetAll() ([]models.User, error) {
	db := db.StorageDB
	var myUsers []models.User
	var userRead models.User
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users"
	rows, err := db.Query(myQuery)
	if err != nil {
		log.Fatal(err)
		return myUsers, err
	}
	for rows.Next() {
		err := rows.Scan(&userRead.ID, &userRead.FirstName, &userRead.LastName, &userRead.Email, &userRead.Age, &userRead.Height, &userRead.Active, &userRead.CrationDate)
		if err != nil {
			log.Fatal(err)
			return myUsers, err
		}
		// En caso de querer devolver mas de uno, por ejemplo un getFirstName
		myUsers = append(myUsers, userRead)
	}
	return myUsers, nil
}

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB
	myQuery := "DELETE FROM users WHERE id = ?"
	stmt, err := db.Prepare(myQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	partUpdate, _ := result.RowsAffected()
	if partUpdate == 0 {
		return errors.New("User not found")
	}
	return nil
}

func (r *repositorySQL) GetFullData() ([]models.User, error) {
	db := db.StorageDB
	var myUsers []models.User
	var userRead models.User
	myQuery := "select u.id,u.first_name,u.last_name,u.email,u.age,u.height,u.active,u.cration_date,c.country_name,c.name from users u inner join city c on u.address = c.id"
	rows, err := db.Query(myQuery)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&userRead.ID, &userRead.FirstName, &userRead.LastName, &userRead.Email, &userRead.Age, &userRead.Height, &userRead.Active, &userRead.CrationDate, &userRead.Address.CountryName, &userRead.Address.Name)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		myUsers = append(myUsers, userRead)
	}
	return myUsers, nil
}

func (r *repositorySQL) GetOneWithContext(ctx context.Context, id int) (models.User, error) {
	db := db.StorageDB
	var userRead models.User
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users WHERE id = ?"
	rows, err := db.QueryContext(context.Background(), myQuery, id)
	if err != nil {
		log.Fatal(err)
		return userRead, err
	}
	for rows.Next() {
		err := rows.Scan(&userRead.ID, &userRead.FirstName, &userRead.LastName, &userRead.Email, &userRead.Age, &userRead.Height, &userRead.Active, &userRead.CrationDate)
		if err != nil {
			log.Fatal(err)
			return userRead, err
		}
	}
	return userRead, nil
}

func (r *repositorySQL) UpdateWithContext(ctx context.Context, user models.User) (models.User, error) {
	db := db.StorageDB
	myQuery := "UPDATE users SET first_name = ?, last_name = ?, email =?, age=?, height =?, active =?, cration_date=? WHERE id=?"
	stmt, err := db.PrepareContext(ctx, myQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate, user.ID)
	if err != nil {
		return models.User{}, err
	}
	partUpdate, _ := result.RowsAffected()
	if partUpdate == 0 {
		return models.User{}, errors.New("User not found")
	}
	return user, nil
}
