package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/17_storage1/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCreada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCargada := service.GetOne(2)

	assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	personaUpdate := models.Persona{
		ID:       2,
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCargada, _ := service.Update(personaUpdate)

	assert.Equal(t, personaUpdate.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaUpdate.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestUpdateServiceSQL_Failed(t *testing.T) {
	//Arrange
	personaUpdate := models.Persona{
		ID:       15,
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	_, err := service.Update(personaUpdate)

	assert.Equal(t, "No se encontro la persona", err.Error())
	// assert.Nil(t, misPersonas)
}
