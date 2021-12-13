package internal

import (
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 1, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	prod, _ := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)

	assert.Equal(t, prodEsperado, prod)

}

func TestUpdateNotFoundMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 4, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	_, err := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.Error(t, err)
}

func TestUpdateErrorMock(t *testing.T) {
	dbMock := store.Mock{Err: errors.New("Error en la conexion")}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 4, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	_, err := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.NotNil(t, err)
}

func TestDeleteMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	err := service.Delete(1)
	assert.Nil(t, err)
	_, err = service.FindById(1)
	assert.NotNil(t, err)

}

func TestDeleteNotFoundMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	err := service.Delete(4)
	assert.Error(t, err)
}

func TestDeleteErrorMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en la conexion")}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	err := service.Delete(4)
	assert.Error(t, err)
}
