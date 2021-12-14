package internal

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TTarde/proyecto/pkg/store"
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
  

  var per string = `[{
    "id": 1,
    "codigo_transaccion": 486499,
    "moneda": "pesos",
    "monto": 84.16,
    "emisor": "Emisor Anterior",
    "receptor": "Hauck-Carter",
    "fecha_transaccion": "13/08/2021"
  }]`

type StubStore struct{}
type mockStorage struct{
  readed bool
}

func (m *mockStorage) Read(data interface{}) error{
  m.readed = true
  return json.Unmarshal([]byte(per), &data)
}

func (m *mockStorage) Write(data interface{}) error{
  return nil
}

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
	err := json.Unmarshal([]byte(perso), &esperado)

  if err != nil {
    log.Fatal(err)
  }

	assert.Equal(t, esperado, personas)
}

func TestUpdateEmisor(t *testing.T) {
  store := mockStorage{}
	repo := NewRepository(&store)

  store.readed = false

  actualizado, _ := repo.UpdateEmisor(1, "Emisor actualizado")

  var esperado Transaccion
  json.Unmarshal([]byte(per), &esperado)

  assert.Equal(t, store.readed, true)
  assert.NotEqual(t, esperado.Emisor, actualizado.Emisor)

}

func TestUpdateMock(t *testing.T){
  dataByte := []byte(per)

  trNuevo := Transaccion{
    ID: 1,
    CodigoTransaccion: 556111,
    Moneda: "Pesos",
    Monto: 80.00,
    Emisor: "Locomotion",
    Receptor: "Disney",
    FechaTransaccion: "13/08/2021",
  }

  dbMock := store.Mock{Data: dataByte}
  storeStub := store.FileStore{Mock: &dbMock}
  repo := NewRepository(&storeStub)


  per_actualizada, _ := repo.Update(trNuevo.ID, trNuevo.CodigoTransaccion, trNuevo.Moneda, trNuevo.Monto, trNuevo.Emisor, trNuevo.Receptor, trNuevo.FechaTransaccion)


  assert.Equal(t, trNuevo, per_actualizada)
}

func TestDeleteMock(t *testing.T) {
  dataByte := []byte(perso)

  dbMock := store.Mock{Data: dataByte}
  storeStub := store.FileStore{Mock: &dbMock}
  repo := NewRepository(&storeStub)

  err := repo.Delete(1)
  todos, _ := repo.GetAll()

  assert.Equal(t, err, nil)
  assert.Equal(t, len(todos), 1)
}


func TestDeleteErrorMock(t *testing.T) {
  dataByte := []byte(perso)

  dbMock := store.Mock{Data: dataByte}
  storeStub := store.FileStore{Mock: &dbMock}
  repo := NewRepository(&storeStub)

  err := repo.Delete(5)

  assert.NotEqual(t, err, nil)
}

func TestStoreMock(t *testing.T) {
  dataByte := []byte(per)
  trNuevo := Transaccion{
    ID: 2,
    CodigoTransaccion: 556111,
    Moneda: "Pesos",
    Monto: 80.00,
    Emisor: "Locomotion",
    Receptor: "Disney",
    FechaTransaccion: "13/08/2021",
  }

  dbMock := store.Mock{Data: dataByte}
  storeStub := store.FileStore{Mock: &dbMock}
  repo := NewRepository(&storeStub)

  tr, _ := repo.Store(trNuevo.ID, trNuevo.CodigoTransaccion, trNuevo.Moneda, trNuevo.Monto, trNuevo.Emisor, trNuevo.Receptor, trNuevo.FechaTransaccion)

  assert.Equal(t, trNuevo, tr)

}



