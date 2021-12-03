package internal

import (
	"fmt"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
)
type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []Persona

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, edad int) (Persona, error)
	Update(id int, nombre string, apellido string, edad int) (Persona, error)
	UpdateNombre(id int, nombre string) (Persona, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Persona, error) {

	//para que persista
	err := repo.db.Read(&personas)
	if err != nil {
		return nil, err
	}
	return personas, nil

}

func (repo *repository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	//para que persista
	repo.db.Read(&personas)

	per := Persona{id, nombre, apellido, edad}

	personas = append(personas, per)
	err := repo.db.Write(personas)

	if err != nil {
		return Persona{}, err
	}

	return per, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&personas)

	if err != nil {
		return 0, err
	}

	if len(personas) == 0 {
		return 0, nil
	}

	return personas[len(personas)-1].ID, nil
}

func (repo *repository) Update(id int, nombre, apellido string, edad int) (Persona, error) {
	err := repo.db.Read(&personas)
	if err != nil {
		return Persona{}, err
	}

	per := Persona{id, nombre, apellido, edad}

	for i := range personas {
		if personas[i].ID == id {
			personas[i] = per
			err := repo.db.Write(personas)
			if err != nil {
				return Persona{}, err
			}
			return per, nil
		}
	}
	return Persona{}, fmt.Errorf("el ID: %d no existe", id)
}

func (repo *repository) UpdateNombre(id int, nombre string) (Persona, error) {
	err := repo.db.Read(&personas)
	if err != nil {
		return Persona{}, err
	}
	for i := range personas {
		if personas[i].ID == id {
			personas[i].Nombre = nombre
			err := repo.db.Write(personas)
			if err != nil {
				return Persona{}, err
			}
			return personas[i], nil
		}
	}
	return Persona{}, fmt.Errorf("el ID: %d no existe", id)
}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&personas)
	if err != nil {
		return err
	}

	index := 0
	for i := range personas {
		if personas[i].ID == id {
			index = i
			personas = append(personas[:index], personas[index+1:]...)
			err := repo.db.Write(personas)
			return err
		}
	}

	return fmt.Errorf("el ID: %d no existe", id)
}