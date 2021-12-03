/*
Repositorio, debe tener el acceso a la variable guardada en memoria.
- Se debe crear el archivo repository.go
- Se debe crear la estructura de la entidad
- Se deben crear las variables globales donde guardar las entidades
- Se debe generar la interface Repository con todos sus métodos
- Se debe generar la estructura repository
- Se debe generar una función que devuelva el Repositorio
- Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/
package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM2/pkg/store"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error)
	LastId() (int, error)
	LoadUser() error
	Update(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error)
	Delete(id int) error
	UpdateLastName(id int, last_name string) (User, error)
	UpdateAge(id int, age int) (User, error)
}

type repository struct {
	db store.Store
}

func (repo *repository) GetAll() ([]User, error) {
	err := repo.db.Read(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *repository) Store(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error) {
	repo.db.Read(&users)

	us := User{id, first_name, last_name, email, age, height, active, create_date}

	users = append(users, us)

	err := repo.db.Write(users)

	if err != nil {
		return User{}, err
	}

	return us, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&users)
	if err != nil {
		return 0, nil
	}
	if len(users) == 0 {
		return 0, nil
	}
	return users[len(users)-1].ID, nil
}

func (r *repository) LoadUser() error {
	// bytes, err := os.ReadFile("../../Exercise1/internal/users/users.json")
	bytes, err := os.ReadFile("/Users/joserios/Desktop/bootcamp/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM2/internal/users/users.json")
	if err != nil {
		return err
	}

	var allUsers []User

	errUnmarshall := json.Unmarshal(bytes, &allUsers)
	if errUnmarshall != nil {
		return err
	}

	users = allUsers
	return nil
}

func (repo *repository) Update(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error) {
	us := User{id, first_name, last_name, email, age, height, active, create_date}
	update := false
	for i := range users {
		if users[i].ID == id {
			us.ID = id
			users[i] = us
			update = true
		}
	}
	if !update {
		return User{}, fmt.Errorf("User %d not found", id)
	}
	return us, nil
}

func (repo *repository) Delete(id int) error {
	deleted := false
	var index int

	for i := range users {
		if users[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("User %d not found", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}

func (repo *repository) UpdateLastName(id int, last_name string) (User, error) {
	for i, v := range users {
		if v.ID == id {
			users[i].LastName = last_name
			return users[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d not found", id)
}

func (repo *repository) UpdateAge(id int, age int) (User, error) {
	for i, v := range users {
		if v.ID == id {
			users[i].Age = age
			return users[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d not found", id)
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}
