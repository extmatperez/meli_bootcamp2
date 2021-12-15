package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/pkg/store"
)

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

type repository struct {
	db store.Store
}

var usuarios []Usuario

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	err := repo.db.Read(&usuarios)
	return usuarios, err
}

func (repo *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {

	err := repo.db.Read(&usuarios)
	if err != nil {
		return Usuario{}, err
	}
	nuevoUsuario := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	usuarios = append(usuarios, nuevoUsuario)
	err = repo.db.Write(usuarios)

	if err != nil {
		return Usuario{}, err
	}
	return nuevoUsuario, nil
}

func (repo *repository) LastID() (int, error) {
	err := repo.db.Read(&usuarios)

	if err != nil {
		return 0, err
	}

	if len(usuarios) == 0 {
		return 0, nil
	}
	return usuarios[len(usuarios)-1].ID, nil
}

func (repo *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return Usuario{}, err
	}

	updateUser := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	for i, user := range usuarios {
		if user.ID == id {
			usuarios[i] = updateUser
			err := repo.db.Write(&usuarios)
			if err != nil {
				return Usuario{}, err
			}
			return updateUser, nil
		}
	}
	return Usuario{}, fmt.Errorf("el usuario con ID %d no existe", id)
}

func (repo *repository) Delete(id int) error {

	err := repo.db.Read(&usuarios)
	if err != nil {
		return err
	}

	for i, user := range usuarios {
		if user.ID == id {
			// usuarios = append(usuarios[:i], usuarios[i+1:]...)
			usuarios[i].Activo = false
			err := repo.db.Write(&usuarios)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("el usuario %d no existe", id)
}

func (repo *repository) EditarNombreEdad(id int, nombre string, edad int) (Usuario, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return Usuario{}, err
	}

	for i, user := range usuarios {
		if user.ID == id {
			fmt.Println(id, " ++ ", user.ID)
			usuarios[i].Nombre = nombre
			usuarios[i].Edad = edad

			err := repo.db.Write(&usuarios)
			if err != nil {
				return Usuario{}, err
			}
			return usuarios[i], nil
		}
	}
	return Usuario{}, fmt.Errorf("usuario %d no existe", id)
}
