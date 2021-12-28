package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/12_testing3/ProyectoEstructura/pkg/store"
	"github.com/stretchr/testify/assert"
)

var prodServiceData string = `[
	{
	 "id": 1,
	 "nombre": "Espejo",
	 "color": "Azul",
	 "precio": 40,
	 "stock": 10,
	 "codigo": "156ujma8ssssDA",
	 "publicado": true,
	 "fecha_creacion": "25/10/2020"
	},
	{
	 "id": 3,
	 "nombre": "Auricular",
	 "color": "Verde",
	 "precio": 25.3,
	 "stock": 10,
	 "codigo": "156ujma8ssssDA",
	 "publicado": true,
	 "fecha_creacion": "25/10/2020"
	}
   ]`

type storeServiceMock struct {
	Data    interface{}
	SpyRead bool
}

func (s *storeServiceMock) Read(data interface{}) error {
	s.SpyRead = true
	json.Unmarshal([]byte(prodData), &data)

	return nil
}
func (s *storeServiceMock) Write(data interface{}) error {
	return nil
}

func TestUpdateService(t *testing.T) {
	//arr
	dataByte := []byte(prodData)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	// store := storeServiceMock{}
	repo := NewRepository(&storeStub)
	ser := NewService(repo)

	miProd, err := ser.Update(1, "EspejoUpdated", "Azul", 40, 10, "156ujma8ssssDA", true, "25/10/2020")

	var expected []Producto
	json.Unmarshal([]byte(prodServiceData), &expected)

	assert.Equal(t, expected[0].ID, miProd.ID)
	assert.Nil(t, err)
	// assert.True(t, store.SpyRead)
}

/*
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
}*/
