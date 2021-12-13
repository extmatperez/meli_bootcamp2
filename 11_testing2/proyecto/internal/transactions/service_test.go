package internal

import (
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceUpdateSuccess(t *testing.T) {
	//Arrange
	tranByte := []byte(tr)
	transactionExpected := Transaction{
		ID:                  5,
		CodigoDeTransaccion: "prueba1",
		Moneda:              "MXN",
		Monto:               123.45,
		Emisor:              "em1",
		Receptor:            "rec1",
		FechaDeTransaccion:  "12-12-2020",
	}
	dbMock := store.Mock{Data: tranByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	miTransaccion, err := service.Update(5, "prueba1", "MXN", 123.45, "em1", "rec1", "12-12-2020")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, miTransaccion)
	assert.True(t, dbMock.Called)
}

func TestServiceUpdateError(t *testing.T) {
	//Arrange
	ExpectedError := fmt.Errorf("db error")
	dbMock := store.Mock{Err: ExpectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	miTransaccion, err := service.Update(5, "prueba1", "MXN", 123.45, "em1", "rec1", "12-12-2020")

	//Assert
	assert.Error(t, err)
	assert.Equal(t, Transaction{}, miTransaccion)
	assert.True(t, dbMock.Called)
}

func TestServiceDeleteSuccess(t *testing.T) {
	//Arrange
	tranByte := []byte(tr)
	// var transactionsExpected []Transaction
	// json.Unmarshal(tranByte, &transactionsExpected)
	dbMock := store.Mock{Data: tranByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	err := service.Delete(5)
	transacciones, _ := service.GetAll(nil)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(transacciones))
	assert.True(t, dbMock.Called)
}

func TestServiceDeleteError(t *testing.T) {
	//Arrange
	ExpectedError := fmt.Errorf("db error")
	dbMock := store.Mock{Err: ExpectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	err := service.Delete(5)
	//Assert
	assert.Error(t, err)
	assert.True(t, dbMock.Called)
}
