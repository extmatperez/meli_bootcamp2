package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/8_goweb3/tarde/ejercicio1/pkg/store"
)

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        string `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
	UpdateNombre(id int, nombre string) (User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository() Repository {
	return &repository{}
}
func (repo *repository) GetAll() ([]User, error) {
	return users, nil
}
func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	repo.db.Read(&users)

	usr := User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}

	users = append(users, usr)
	err := repo.db.Write(users)

	if err != nil {
		return User{}, err
	}
	return usr, nil
}
func (repo *repository) LastID() (int, error) {
	err := repo.db.Read(&users)

	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, nil
	}
	return lastID, nil
}

func (repo *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	err := repo.db.Read(&users)
	if err != nil {
		return User{}, err
	}
	usr := User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	for i, v := range users {
		if v.ID == id {
			users[i] = usr
			err := repo.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return usr, nil
		}
	}
	return User{}, fmt.Errorf("la persona %d no existe", id)

}

func (repo *repository) UpdateNombre(id int, nombre string) (User, error) {

	for i, v := range users {
		if v.ID == id {
			users[i].Nombre = nombre
			return users[i], nil
		}
	}
	return User{}, fmt.Errorf("la persona %d no existe", id)

}

func (repo *repository) Delete(id int) error {

	err := repo.db.Read(&users)
	if err != nil {
		return err
	}

	index := 0
	for i, v := range users {
		if v.ID == id {
			index = i
			users = append(users[:index], users[index+1:]...)
			err := repo.db.Write(users)

			return err
		}
	}
	return fmt.Errorf("la persona %d no existe", id)

}
