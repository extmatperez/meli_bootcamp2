package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

var produ string = `[ 
	{	"id": 1,	"nombre": "Banana",	"color": "Amarillo",	"precio": 27.99 , "stock" : 100, "codigo":"300","publicado": true,"fecha_de_creacion":"22/04/1991"  },
	{	"id": 2,	"nombre": "Manzana",	"color": "Rojo",	"precio": 17.99 , "stock" : 50, "codigo":"210","publicado": true,"fecha_de_creacion":"22/05/1995"  }]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(produ), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//Act
	misProducts, _ := repo.GetAll()
	var expected []Product
	err := json.Unmarshal([]byte(produ), &expected)
	if err != nil {
		panic(err)
	}
	//Assert
	assert.Equal(t, expected, misProducts)
}
func TestLastID(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	lastID := 2

	//Act
	ultimoID, _ := repo.LastId()

	//Assert
	assert.Equal(t, lastID, ultimoID)
}

func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Banana"

	//Act
	produActualizado, err := repo.Update(1, nombreExpected, "Naranja", 45.33, 100, "222", true, "22/10/20")

	//Assert
	assert.Equal(t, nombreExpected, produActualizado.Nombre)
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Pedro"

	//Act
	_, err := repo.Update(4, nombreExpected, "Naranja", 45.33, 100, "222", true, "22/10/20")

	//Assert
	assert.Error(t, err)
}
