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

func TestGetByNameServiceSQL(t *testing.T) {
	//Arrange
	nombreProd := "Yogurt"

	repo := newRepositorySQL()

	service := NewServiceSQL(repo)

	prodObtenidos, err := service.GetByName(nombreProd)

	assert.Error(t, err)
	assert.True(t, len(prodObtenidos) >= 0)
	// assert.Equal(t, .Nombre, .Nombre)
}
