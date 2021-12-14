package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

var productos string = "[{\"id\": 1,\"nombre\": \"Otro producto\",\" color\": \" Otro color\",\"precio\": \"$800\",\"stock\": 10,\"codigo\": \"AAAAAAAAAA\",\"publicado\": false,\"creado\": \"10/10/2020\"},{\"id\": 3,\"nombre\": \"Producto\",\"color\": \"\",\"precio\": \"$700\",\"stock\": 21,\"codigo\": \"32CRHI85275114\",\"publicado\": true,\"creado\": \"23/5/2020\"}]"

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(productos), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	store := StubStore{}
	repo := NewRepository(&store)

	ps, _ := repo.GetAll()
	var expected []Producto
	_ = json.Unmarshal([]byte(productos), &expected)

	assert.Equal(t, ps, expected)
}

func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Pedro"

	//Act
	personaActualizada, err := repo.Edit(1, nombreExpected, "Rojo", "$45.00", 24, "abcd", true, "25/12/2003")

	//Assert
	assert.Equal(t, nombreExpected, personaActualizada.Nombre)
	assert.Nil(t, err)
}
