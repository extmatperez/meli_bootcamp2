package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TTarde/proyecto/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
	"github.com/go-playground/assert/v2"
)

func TestServiceUpdateMock(t *testing.T) {
	dataByte := []byte(per)

	trNuevo := Transaccion{
		ID:                1,
		CodigoTransaccion: 556111,
		Moneda:            "Pesos",
		Monto:             80.00,
		Emisor:            "Locomotion",
		Receptor:          "Disney",
		FechaTransaccion:  "13/08/2021",
	}

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	per_actualizada, _ := service.Update(trNuevo.ID, trNuevo.CodigoTransaccion, trNuevo.Moneda, trNuevo.Monto, trNuevo.Emisor, trNuevo.Receptor, trNuevo.FechaTransaccion)

	assert.Equal(t, trNuevo, per_actualizada)



}


func TestServiceDeleteMock(t *testing.T) {
	dataByte := []byte(perso)
  
	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	err := service.Delete(1)
	todos, _ := service.GetAll()
  
	assert.Equal(t, err, nil)
	assert.Equal(t, len(todos), 1)
  }

  func TestStoreSQL(t *testing.T) {
	  transaccionNueva := models.Transaccion{
		CodigoTransaccion: 556111,
		Moneda:            "Pesos",
		Monto:             80.00,
		Emisor:            "Locomotion",
		Receptor:          "Disney",
		FechaTransaccion:  "13/08/2021",
	  }

	  repo := NewRepositorySQL()

	  service := NewServiceSQL(repo)

	  transaccionCreada, _ := service.Store(transaccionNueva.CodigoTransaccion, transaccionNueva.Moneda, transaccionNueva.Monto, transaccionNueva.Emisor, transaccionNueva.Receptor, transaccionNueva.FechaTransaccion)

	  assert.Equal(t, transaccionNueva.Emisor, transaccionCreada.Emisor)
  }


  func TestGetByNameSQL(t *testing.T) {
	transaccionNueva := models.Transaccion{
	  CodigoTransaccion: 556111,
	  Moneda:            "Pesos",
	  Monto:             80.00,
	  Emisor:            "Locomotion",
	  Receptor:          "Disney",
	  FechaTransaccion:  "13/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionCargada := service.GetByName("Locomotion")

	assert.Equal(t, transaccionNueva.Emisor, transaccionCargada.Emisor)
}