package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/internal/models"
)

type ServiceSQL interface {
	Store(nombre, apellido string, edad int) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(ctx context.Context, persona models.Persona) (models.Persona, error)
	Delete(id int) error
	GetByName(nombre string) (models.Persona, error)
	GetAll() ([]models.PersonaGetAllDTO, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(nombre, apellido string, edad int) (models.Persona, error) {

	newPersona := models.Persona{Nombre: nombre, Apellido: apellido, Edad: edad}
	personaCreada, err := ser.repository.Store(newPersona)

	if err != nil {
		return models.Persona{}, err
	}
	return personaCreada, nil
}

func (ser *serviceSQL) GetOne(id int) models.Persona {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(ctx context.Context, persona models.Persona) (models.Persona, error) {
	return ser.repository.Update(ctx, persona)
}

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *serviceSQL) GetByName(nombre string) (models.Persona, error) {
	return ser.repository.GetByName(nombre)
}

func (ser *serviceSQL) GetAll() ([]models.PersonaGetAllDTO, error) {
	return ser.repository.GetAll()
}
