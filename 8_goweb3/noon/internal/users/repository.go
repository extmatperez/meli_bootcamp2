package internal

import "fmt"

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

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	LastID() (int, error)
	Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	u := User{id, name, lastName, email, age, height, active, created}
	users = append(users, u)
	lastID = u.ID
	return u, nil
}

func (r *repository) Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	u := User{id, name, lastName, email, age, height, active, created}
	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return User{}, fmt.Errorf("User %d not found", id)
	}

	users[i] = u
	return u, nil
}
