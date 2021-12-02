package internal

import "fmt"

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
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (repo *repository) GetAll() ([]User, error) {
	return users, nil
}
func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	p := User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	lastID = id
	users = append(users, p)
	return p, nil
}
func (repo *repository) LastID() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	usr := User{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	for i, v := range users {
		if v.ID == id {
			users[i] = usr
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

	index := 0
	for i, v := range users {
		if v.ID == id {
			index = i
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("La persona %d no existe", id)

}
