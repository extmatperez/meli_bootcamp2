package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var prodData string = `[
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

type storeMock struct {
	SpyRead bool
}

func (s *storeMock) Read(data interface{}) error {
	s.SpyRead = true
	json.Unmarshal([]byte(prodData), &data)

	return nil
}
func (s *storeMock) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arr
	store := storeMock{}
	repo := NewRepository(&store)

	misProd, err := repo.GetAll()

	var expected []Producto
	json.Unmarshal([]byte(prodData), &expected)

	assert.Equal(t, expected, misProd)
	assert.Nil(t, err)
}

func TestUpdateName(t *testing.T) {
	//arr
	store := storeMock{}
	repo := NewRepository(&store)

	miProd, err := repo.Update(1, "EspejoUpdated", "Azul", 40, 10, "156ujma8ssssDA", true, "25/10/2020")

	var expected []Producto
	json.Unmarshal([]byte(prodData), &expected)

	assert.NotEqual(t, expected[0].Nombre, miProd.Nombre)
	assert.Nil(t, err)
	assert.True(t, store.SpyRead)
}
