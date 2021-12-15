package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/12_testing3/ProyectoEstructura/pkg/store"
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
	err := json.Unmarshal([]byte(prodData), &data)
	if err != nil {
		return err
	}
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

	erro := json.Unmarshal([]byte(prodData), &expected)

	if erro != nil {
		assert.Error(t, erro)
	}
	assert.Equal(t, expected, misProd)
	assert.Nil(t, err)
}

func TestUpdateName(t *testing.T) {
	//arr
	store := storeMock{}
	repo := NewRepository(&store)

	miProd, err := repo.Update(1, "EspejoUpdated", "Azul", 40, 10, "156ujma8ssssDA", true, "25/10/2020")

	var expected []Producto
	errUnm := json.Unmarshal([]byte(prodData), &expected)
	if errUnm != nil {
		assert.Error(t, errUnm)
	}

	assert.NotEqual(t, expected[0].Nombre, miProd.Nombre)
	assert.Nil(t, err)
	assert.True(t, store.SpyRead)
}
func TestStore(t *testing.T) {
	//arr
	dataByte := []byte(prodData)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	// store := storeServiceMock{}
	repo := NewRepository(&storeStub)
	// store := storeMock{}
	// repo := NewRepository(&store)

	miProd, err := repo.Store(1, "EspejoCreated", "Azul", 40, 10, "156ujma8ssssDA", true, "25/10/2020")

	var expected []Producto
	errUnm := json.Unmarshal([]byte(prodData), &expected)
	if errUnm != nil {
		assert.Error(t, errUnm)
	}

	assert.NotEqual(t, expected[0].Nombre, miProd.Nombre)
	assert.Nil(t, err)
}
