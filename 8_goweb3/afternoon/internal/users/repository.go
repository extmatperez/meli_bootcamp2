package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/afternoon/pkg/store"
)

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name""`
	LastName string  `json:"last_name""`
	Email    string  `json:"email""`
	Age      int     `json:"age""`
	Height   float64 `json:"height""`
	Active   bool    `json:"active""`
	Created  string  `json:"created""`
}

var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	LastID() (int, error)
	Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	UpdateLastNameAge(id int, lastName string, age int) (User, error)
	Delete(id int) (bool, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	r.db.Read(&users)

	return users, nil
}

func (r *repository) LastID() (int, error) {
	var users []User
	err := r.db.Read(&users)

	if len(users) == 0 || err != nil {
		return 0, nil
	}
	fmt.Println(users[len(users)-1])
	return users[len(users)-1].ID, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	var users []User
	r.db.Read(&users)

	u := User{id, name, lastName, email, age, height, active, created}
	users = append(users, u)

	err := r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	var users []User
	r.db.Read(&users)

	u := User{id, name, lastName, email, age, height, active, created}
	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return User{}, nil
	}

	users[i] = u

	err := r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) UpdateLastNameAge(id int, lastName string, age int) (User, error) {
	var users []User
	r.db.Read(&users)

	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return User{}, nil
	}

	users[i].LastName = lastName
	users[i].Age = age

	err := r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return users[i], nil
}

func (r *repository) Delete(id int) (bool, error) {
	var users []User
	r.db.Read(&users)

	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return false, nil
	}

	users = append(users[:i], users[i+1:]...)
	err := r.db.Write(users)

	if err != nil {
		return false, err
	}

	return true, nil
}
