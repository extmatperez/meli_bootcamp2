package internal

import (
	"context"
	"database/sql"
	"log"

	models "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/models"
)

const (
	getAllQuery = "SELECT id, name, last_name, email, age, height, active, created FROM users"
	getOneQuery = "SELECT id, name, last_name, email, age, height, active, created FROM users WHERE id = ?"
	insertQuery = "INSERT INTO users(name, last_name, email, age, height, active, created) VALUES (?,?,?,?,?,?,?)"
	updateQuery = "UPDATE users SET name = ?, last_name = ?, email = ?, age = ?, height = ?, active = ?, created = ? WHERE id = ?"
)

type RepositorySql interface {
	GetAll() ([]models.User, error)
	GetOne(id int) (models.User, error)
	Store(user models.User) (models.User, error)
	Update(user models.User, ctx context.Context) (models.User, error)
}

type repositorySql struct {
	db *sql.DB
}

func NewRepositorySql(db *sql.DB) RepositorySql {
	return &repositorySql{db: db}
}

func (s *repositorySql) GetAll() ([]models.User, error) {
	var user models.User
	var users []models.User

	rows, err := s.db.Query(getAllQuery)

	if err != nil {
		log.Println(err)
		return []models.User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.Created)
		if err != nil {
			log.Println(err)
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *repositorySql) GetOne(id int) (models.User, error) {
	var user models.User

	rows, err := s.db.Query(getOneQuery, id)

	if err != nil {
		log.Println(err)
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.Created)
		if err != nil {
			log.Println(err)
			return user, err
		}
	}

	return user, nil
}

func (s *repositorySql) Store(user models.User) (models.User, error) {
	stmt, err := s.db.Prepare(insertQuery)

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

func (s *repositorySql) Update(user models.User, ctx context.Context) (models.User, error) {
	stmt, err := s.db.Prepare(updateQuery)

	if err != nil {
		log.Fatal(err)
		return models.User{}, err
	}

	_, err = stmt.ExecContext(ctx, user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
