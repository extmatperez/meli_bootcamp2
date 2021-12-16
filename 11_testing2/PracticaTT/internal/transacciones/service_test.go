package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTT/pkg/store"
	"github.com/stretchr/testify/assert"
)

// func (serv *service) Update(id int, codTransaccion, moneda string,
// 	monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)

var transactions []Transaccion = []Transaccion{
	{1, "a1b2b3", "pesos", 4444.33, "Matias", "Esteban", "21/10/2021"},
	{2, "c1c2c3", "pesos", 5555.33, "Jorge", "Esteban", "21/09/2021"},
	{3, "asdfv32", "pesos", 1231.33, "Sebastian", "Esteban", "21/09/2021"},
}

func TestUpdateTransacMock(t *testing.T) {
	//Arrange
	var newTransac Transaccion = Transaccion{1, "abcd1234", "dolar", 550.33, "Facundo", "Esteban", "21/11/2021"}

	dataByte, _ := json.Marshal(transactions)
	// fmt.Println(string(dataByte))
	dbMock := store.MockStore{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	resTransac, err := service.Update(newTransac.Id, newTransac.CodTransaccion, newTransac.Moneda, newTransac.Monto, newTransac.Emisor, newTransac.Receptor, newTransac.FechaTrans)

	//Assert
	assert.Equal(t, resTransac, newTransac)
	assert.Nil(t, err)
}

func TestDeleteTransacMock(t *testing.T) {
	//Arrange
	transSpected := Transaccion{2, "c1c2c3", "pesos", 5555.33, "Jorge", "Esteban", "21/09/2021"}

	dataByte, _ := json.Marshal(transactions)
	dbMock := store.MockStore{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	transDeleted1, err1 := service.Delete(2)
	transDeleted2, err2 := service.Delete(8)

	transDeleted3, err3 := service.Search("2")

	//Assert
	assert.Equal(t, transSpected, transDeleted1)
	assert.Nil(t, err1)

	assert.Empty(t, transDeleted2)
	assert.NotNil(t, err2)
	assert.Error(t, err2)

	assert.Empty(t, transDeleted3)
	assert.NotNil(t, err3)
	assert.Error(t, err3)
}
