package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	withError bool
}

var productsJsonTestGetAll string = `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`

func (s *StubStore) Read(data interface{}) error {
	if s.withError {
		return errors.New("Error de conexion a base de datos")
	}
	return json.Unmarshal([]byte(productsJsonTestGetAll), data)
}
func (s *StubStore) Write(data interface{}) error {
	return nil
}

type SpyStore struct {
	passRead bool
}

var productsJsonTestUpdateName string = `[{"id":1,"nombre":"Before Update","color":"Blue","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`

func (s *SpyStore) Read(data interface{}) error {
	s.passRead = true
	return json.Unmarshal([]byte(productsJsonTestUpdateName), data)
}
func (s *SpyStore) Write(data interface{}) error {
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	productsJsonTestUpdateName = string(content)
	return nil
}

func TestGetAll(t *testing.T) {
	db := StubStore{}
	prod1 := Product{ID: 1, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod2 := Product{ID: 2, Nombre: "Agenda - gris", Color: "Indigo", Precio: 450, Stock: 59, Codigo: "WBADX1C50CE880177", Publicado: true, FechaCreacion: "09/06/2021"}
	repository := NewRepository(&db)
	resultProds, _ := repository.GetAll()
	assert.Equal(t, len(resultProds), 2)
	assert.Equal(t, resultProds[0], prod1)
	assert.Equal(t, resultProds[1], prod2)
}

func TestGetAllWithError(t *testing.T) {
	db := StubStore{true}
	repository := NewRepository(&db)
	_, err := repository.GetAll()
	assert.Error(t, err)
}

func TestUpdateNameAndPrice(t *testing.T) {
	db := SpyStore{}
	repository := NewRepository(&db)
	result, _ := repository.UpdateNameAndPrice(1, "After Update", 1)
	assert.True(t, db.passRead)
	assert.Equal(t, result.Nombre, "After Update")
}

func TestUpdateNameAndPriceNotFound(t *testing.T) {
	db := SpyStore{}
	repository := NewRepository(&db)
	_, err := repository.UpdateNameAndPrice(4, "After Update", 1)
	assert.True(t, db.passRead)
	assert.Error(t, err)
}

func TestUpdateNameAndPriceError(t *testing.T) {
	db := StubStore{true}
	repository := NewRepository(&db)
	_, err := repository.UpdateNameAndPrice(1, "After Update", 1)
	assert.Error(t, err)
}
