package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/17_storage1/proyectoEjemploClase/internal/models"
)

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Daniela",
		Apellido: "Sagovia",
		Edad:     30,
	}
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)
	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCreada.Apellido)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Durazno",
		Apellido: "Rivero",
		Edad:     85,
	}
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	personaCreada := service.GetOne(2)
	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCreada.Apellido)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	personaUpdate := models.Persona{
		ID:       3,
		Nombre:   "DuraznA",
		Apellido: "RiverA",
		Edad:     15,
	}
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	personaCargada, _ := service.Update(context.Background(), personaUpdate)
	assert.Equal(t, personaUpdate.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaUpdate.Apellido, personaCargada.Apellido)
}

func TestDeleteServiceSQL(t *testing.T) {
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	error := service.Delete(4)
	assert.Nil(nil, error)

}

func TestGetByNameServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     85,
	}
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	personaEncontrada, _ := service.GetByName("Juan")
	assert.Equal(t, personaNueva.Apellido, personaEncontrada.Apellido)
}

func TestGetAllServiceSQL(t *testing.T) {
	repo := NewRepositorioSQL()
	service := NewServiceSQL(repo)
	personaEncontrada, _ := service.GetAll()
	assert.Equal(t, "Avellaneda", personaEncontrada[0].Domicilio.Ciudad)
}
