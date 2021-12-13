package internal

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	useRead bool
}

var product string = `[
		{"id": 1, "nombre": "prodcuto before update","color": "rojo", "precio": 20, "stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"},
   		{"id": 2,"nombre": "producto1","color": "rojo","precio": 20,"stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"
	}]`

func (s *StubStore) Read(data interface{}) error {
	s.useRead = true
	return json.Unmarshal([]byte(product), data)

}

func (s *StubStore) Write(data interface{}) error {
	return nil
}
func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//Act
	misProductos, _ := repo.GetAll()
	var expected []Product
	json.Unmarshal([]byte(product), &expected)
	fmt.Println(expected)
	//assert
	assert.Equal(t, expected, misProductos)
}
func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "producto After"
	//productoToUpdate := repo.
	//Act
	productActualizado, err := repo.UpdateNombre(1, nombreExpected)
	//assert
	assert.Equal(t, nombreExpected, productActualizado.Nombre)
	assert.Nil(t, err)
	assert.True(t, true, store.useRead)
	fmt.Println(productActualizado)
}

// falta prbar el error
