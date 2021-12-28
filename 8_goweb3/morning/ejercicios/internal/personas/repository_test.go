package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
)

type StubStore struct {}

var perso string = `[
{
	"id": 1,
	"nombre": "donald",
	"apellido": "trump",
	"edad": 27
   },
   {
	"id": 2,
	"nombre": "jair",
	"apellido": "bolsonaro",
	"edad": 55
   }
]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(perso), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

/* func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//Act
	misPersonas, _ := repo.GetAll()
	var expectedResult []Persona
	json.Unmarshal([]byte(perso), &expectedResult)

	assert.Equal(t, expectedResult, misPersonas)
}  */

func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Pedro"

	//Act
	_, err := repo.Update(2, nombreExpected, "Perez", 45)

	assert.Nil(t, err)
}

func TestGetAllRepositoryMock(t *testing.T) {
	// Initializing input/output
	/* input := []Persona{
		{
			ID: 1,
			Nombre: "Boris",
			Apellido: "Jhonson",
			Edad: 25,
		}, {
			ID: 2,
			Nombre: "Jair",
			Apellido: "Bolsonaro",
			Edad: 32,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock: &dbMock,
	} */

	dataByte := []byte(perso)
	dbMock := store.Mock{Data: dataByte}

	var personasEsperadas []Persona
	json.Unmarshal(dataByte, &personasEsperadas)

	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)

	//Act
	misPersonas, _ := repo.GetAll()

	//Assert
	assert.Equal(t, personasEsperadas, misPersonas)
}