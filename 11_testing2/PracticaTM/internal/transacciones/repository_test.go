package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var transac string = `[
	{
	"id": 1,
	"cod_transaccion": "123asd456",
	"moneda": "dolar",
	"monto": "20.55",
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   },
   {
	"id": 2,
	"cod_transaccion": "123asd456",
	"moneda": "dolar",
	"monto": "20.55",
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   },
   {
	"id": 3,
	"cod_transaccion": "123asd456",
	"moneda": "dolar",
	"monto": "20.55",
	"emisor": "Facundo",
	"receptor": "Ezequiel",
	"fecha_trans": "21/01/2021"
   }
   ]
`

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(transac), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	store := StubStore{}
	repo := NewRepository(&store)
	var transacExpected []Transaccion

	json.Unmarshal([]byte(transac), &transacExpected)

	myTransac, _ := repo.getAll()

	assert.Equal(t, transacExpected, myTransac)

}
