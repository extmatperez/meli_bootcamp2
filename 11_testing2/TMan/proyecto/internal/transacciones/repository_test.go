package internal

import (
	"encoding/json"
	"testing"
	"github.com/go-playground/assert/v2"
)


var perso string = `[{
    "id": 1,
    "codigo_transaccion": 486499,
    "moneda": "pesos",
    "monto": 84.16,
    "emisor": "Lakin LLC",
    "receptor": "Hauck-Carter",
    "fecha_transaccion": "13/08/2021"
  }, {
    "id": 2,
    "codigo_transaccion": 232323,
    "moneda": "yenes",
    "monto": 99.13,
    "emisor": "Breitenberg and Sons",
    "receptor": "Hickle-Barrows",
    "fecha_transaccion": "25/06/2021"
  }]`

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error{
	return json.Unmarshal([]byte(perso), &data)
}

func (s *StubStore) Write(data interface{}) error{
	return nil
}



func TestGetAll(t *testing.T){
	store := StubStore{}
	repo := NewRepository(&store)

	personas, _ := repo.GetAll()
	var esperado []Transaccion
	json.Unmarshal([]byte(perso), &esperado)

	assert.Equal(t, esperado, personas)
}
