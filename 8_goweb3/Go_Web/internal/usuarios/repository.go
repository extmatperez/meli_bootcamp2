package internal

import "fmt"

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
	Delete(id int) error
	EditarNombreEdad(id int, nombre string, edad int) (Usuario, error)
}

type repository struct{}

var usuarios []Usuario
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	lastID = id
	nuevoUsuario := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	usuarios = append(usuarios, nuevoUsuario)
	return nuevoUsuario, nil
}

func (repo *repository) LastID() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	updateUser := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	for i, user := range usuarios {
		if user.ID == id {
			usuarios[i] = updateUser
			return updateUser, nil
		}
	}
	return Usuario{}, fmt.Errorf("el usuario con ID %d no existe", id)
}

func (repo *repository) Delete(id int) error {

	for i, user := range usuarios {
		if user.ID == id {
			// usuarios = append(usuarios[:i], usuarios[i+1:]...)
			usuarios[i].Activo = false
			return nil
		}
	}
	return fmt.Errorf("el usuario %d no existe", id)
}

func (repo *repository) EditarNombreEdad(id int, nombre string, edad int) (Usuario, error) {
	for i, user := range usuarios {
		if user.ID == id {
			fmt.Println(id, " ++ ", user.ID)
			fmt.Println(user)
			usuarios[i].Nombre = nombre
			usuarios[i].Edad = edad
			fmt.Println(user)
			return usuarios[i], nil
		}
	}
	return Usuario{}, fmt.Errorf("usuario %d no existe", id)
}
