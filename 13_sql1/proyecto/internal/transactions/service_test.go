package internal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp2/13_sql1/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/13_sql1/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAllSuccess(t *testing.T) {
	//Arrange
	dbMock := store.Mock{Data: []byte(tr)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	filtro := map[string]string{"Moneda": "UYU"}
	transacciones, err := service.GetAll(filtro)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(transacciones))
}

func TestServiceGetAllErrors(t *testing.T) {
	//Arrange
	ExpectedError := fmt.Errorf("db error")
	dbMock := store.Mock{Err: ExpectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	filtro := map[string]string{"Moneda": "UYU"}
	transacciones, err := service.GetAll(filtro)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, 0, len(transacciones))
}

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

func TestServiceStoreSuccess(t *testing.T) {
	//Arrange
	tranByte := []byte(tr)
	transactionExpected := Transaction{
		ID:                  7,
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
	miTransaccion, err := service.Store("prueba1", "MXN", 123.45, "em1", "rec1", "12-12-2020")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, miTransaccion)
	assert.True(t, dbMock.Called)
}

func TestServiceStoreError(t *testing.T) {
	//Arrange
	ExpectedError := fmt.Errorf("Error del servidor")
	dbMock := store.Mock{Err: ExpectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	//Act
	miTransaccion, err := service.Store("prueba1", "MXN", 123.45, "em1", "rec1", "12-12-2020")

	//Assert
	assert.Error(t, err)
	assert.Equal(t, Transaction{}, miTransaccion)
	assert.True(t, dbMock.Called)
}

//SQL

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	transactionExpected := models.Transaction{
		ID:                  7,
		CodigoDeTransaccion: "pruebaStore",
		Moneda:              "MXN",
		Monto:               123.45,
		Emisor:              "em1",
		Receptor:            "rec1",
		FechaDeTransaccion:  "2020-12-12 00:00:00",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	//Act
	miTransaccion, err := service.Store(transactionExpected.CodigoDeTransaccion, transactionExpected.Moneda, transactionExpected.Monto, transactionExpected.Emisor, transactionExpected.Receptor, transactionExpected.FechaDeTransaccion)
	defer service.Delete(miTransaccion.ID)
	transactionExpected.ID = miTransaccion.ID
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, miTransaccion)
}

func TestUpdateWithContextServiceSQL(t *testing.T) {
	//Arrange
	transactionExpected := models.Transaction{
		CodigoDeTransaccion: "testUpdatewc",
		Moneda:              "MXN",
		Monto:               123.45,
		Emisor:              "em1",
		Receptor:            "rec1",
		FechaDeTransaccion:  "2020-12-12 00:00:00",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)
	recibida, _ := service.Store(transactionExpected.CodigoDeTransaccion, "UYU", 254.33, "emisor3", "rec0", transactionExpected.FechaDeTransaccion)
	defer service.Delete(recibida.ID)
	transactionExpected.ID = recibida.ID
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//Act
	miTransaccion, err := service.UpdateWithContext(ctx, transactionExpected.ID, transactionExpected.CodigoDeTransaccion, transactionExpected.Moneda, transactionExpected.Monto, transactionExpected.Emisor, transactionExpected.Receptor, transactionExpected.FechaDeTransaccion)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, miTransaccion)

}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	transactionExpected := models.Transaction{
		CodigoDeTransaccion: "testUpdate",
		Moneda:              "MXN",
		Monto:               123.45,
		Emisor:              "em1",
		Receptor:            "rec1",
		FechaDeTransaccion:  "2020-12-12 00:00:00",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)
	recibida, _ := service.Store(transactionExpected.CodigoDeTransaccion, "UYU", 254.33, "emisor3", "rec0", transactionExpected.FechaDeTransaccion)
	defer service.Delete(recibida.ID)
	transactionExpected.ID = recibida.ID
	//Act
	miTransaccion, err := service.Update(transactionExpected.ID, transactionExpected.CodigoDeTransaccion, transactionExpected.Moneda, transactionExpected.Monto, transactionExpected.Emisor, transactionExpected.Receptor, transactionExpected.FechaDeTransaccion)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, miTransaccion)

}

func TestDeleteServiceSQL(t *testing.T) {

	//Arrange
	transactionToDelete := models.Transaction{
		CodigoDeTransaccion: "testDelete",
		Moneda:              "MXN",
		Monto:               123.45,
		Emisor:              "em1",
		Receptor:            "rec1",
		FechaDeTransaccion:  "2020-12-12 00:00:00",
	}
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)
	recibida, _ := service.Store(transactionToDelete.CodigoDeTransaccion, transactionToDelete.Moneda, transactionToDelete.Monto, transactionToDelete.Emisor, transactionToDelete.Receptor, transactionToDelete.FechaDeTransaccion)
	//Act
	err := service.Delete(recibida.ID)
	//Assert
	assert.Nil(t, err)

}
