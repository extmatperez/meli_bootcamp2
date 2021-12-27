package internal

import (
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/store"
)

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Active   bool    `json:"active"`
	Created  string  `json:"created"`
}

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
	err := r.db.Read(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) LastID() (int, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return 0, err
	}

	if len(users) == 0 || err != nil {
		return 0, nil
	}
	return users[len(users)-1].ID, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return User{}, err
	}

	u := User{id, name, lastName, email, age, height, active, created}
	users = append(users, u)

	err = r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return User{}, err
	}

	if err != nil {
		return User{}, err
	}

	u := User{id, name, lastName, email, age, height, active, created}
	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return User{}, nil
	}

	users[i] = u

	err = r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) UpdateLastNameAge(id int, lastName string, age int) (User, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return User{}, err
	}

	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return User{}, nil
	}

	users[i].LastName = lastName
	users[i].Age = age

	err = r.db.Write(users)
	if err != nil {
		return User{}, err
	}

	return users[i], nil
}

func (r *repository) Delete(id int) (bool, error) {
	var users []User
	err := r.db.Read(&users)

	if err != nil {
		return false, err
	}

	i := 0
	for i < len(users) && users[i].ID != id {
		i++
	}

	if i == len(users) {
		return false, nil
	}

	users = append(users[:i], users[i+1:]...)
	err = r.db.Write(users)

	if err != nil {
		return false, err
	}

	return true, nil
}
