package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/Exercise1/internal/models"
)

type RepositorySQLMock interface {
	Store(user models.User) (models.User, error)
	GetOne(id int) models.User
	Update(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(id int) error
	GetOneWithContext(ctx context.Context, id int) (models.User, error)
	GetFullData() ([]models.User, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(user models.User) (models.User, error) {
	myQuery := "INSERT INTO users(first_name, last_name, email, age, height, active, cration_date) VALUES(?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(myQuery)
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

func (r *repositorySQLMock) GetOne(id int) models.User {

	var userRead models.User
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users WHERE id = ?"
	rows, err := r.db.Query(myQuery, id)

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
func (r *repositorySQLMock) GetAll() ([]models.User, error) {
	var misUsers []models.User
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users"
	var userRead models.User
	rows, err := r.db.Query(myQuery)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&userRead.ID, &userRead.FirstName, &userRead.LastName, &userRead.Email, &userRead.Age, &userRead.Height, &userRead.Active, &userRead.CrationDate)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misUsers = append(misUsers, userRead)
	}
	return misUsers, nil
}

func (r *repositorySQLMock) Update(user models.User) (models.User, error) {
	myQuery := "UPDATE users SET first_name = ?, last_name = ?, email =?, age=?, height =?, active =?, cration_date=? WHERE id=?"
	stmt, err := r.db.Prepare(myQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate, user.ID)
	if err != nil {
		return models.User{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.User{}, errors.New("No se encontro la user")
	}

	return user, nil
}

func (r *repositorySQLMock) Delete(id int) error {
	myQuery := "DELETE FROM users WHERE id = ?"
	stmt, err := r.db.Prepare(myQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("No se encontro la user")
	}
	return nil
}

func (r *repositorySQLMock) GetFullData() ([]models.User, error) {
	var misUsers []models.User
	myQuery := "select u.id,u.first_name,u.last_name,u.email,u.age,u.height,u.active,u.cration_date,c.country_name,c.name from users u inner join city c on u.address = c.id"
	var userRead models.User
	rows, err := r.db.Query(myQuery)

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
		misUsers = append(misUsers, userRead)
	}
	return misUsers, nil
}

func (r *repositorySQLMock) GetOneWithContext(ctx context.Context, id int) (models.User, error) {

	var userRead models.User
	myQuery := "SELECT id,first_name, last_name, email, age, height, active, cration_date FROM users WHERE id = ?"
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := r.db.QueryContext(ctx, myQuery, id)

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
