package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

var tr string = `[
	{"id": 5,"codigo_de_transaccion": "asd12321","moneda": "USD","monto": 23400.85,"emisor": "sebac2","receptor": "rec","fecha_de_transaccion": "10/18/2021"},
	{"id": 6,"codigo_de_transaccion": "asd12321","moneda": "UYU","monto": 2344.8,"emisor": "sebac","receptor": "rec","fecha_de_transaccion": "10/18/2021"}
	]`

//Clase 1
type stubStore struct{}

func (st *stubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(tr), &data)
}

func (st *stubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := &stubStore{}
	repo := NewRepository(store)

	//Act
	misTransacciones, _ := repo.GetAll()
	var expected []Transaction
	json.Unmarshal([]byte(tr), &expected)
	//Assert
	assert.Equal(t, expected, misTransacciones)
}

type mockStore struct {
	called bool
}

func (ms *mockStore) Read(data interface{}) error {
	ms.called = true
	return json.Unmarshal([]byte(`[{"id": 5,"codigo_de_transaccion": "beforeupdate","moneda": "USD","monto": 100.0,"emisor": "sebac2","receptor": "rec","fecha_de_transaccion": "10/18/2021"}]`), &data)
}

func (ms *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateCodigoYMontoSuccess(t *testing.T) {
	//Arrange
	store := &mockStore{}
	repo := NewRepository(store)

	//Act
	updatedTransaction, err := repo.UpdateCodigoYMonto(5, "afterupdate", 200.0)
	var expected Transaction
	json.Unmarshal([]byte(`{"id": 5,"codigo_de_transaccion": "afterupdate","moneda": "USD","monto": 200.0,"emisor": "sebac2","receptor": "rec","fecha_de_transaccion": "10/18/2021"}`), &expected)
	//Assert
	assert.Equal(t, expected, updatedTransaction)
	assert.True(t, store.called)
	assert.Nil(t, err)
}

// Clase 2

func TestRepositoryGetAllMock(t *testing.T) {
	//Arrange
	tranByte := []byte(tr)
	var transactionsExpected []Transaction
	json.Unmarshal(tranByte, &transactionsExpected)
	dbMock := store.Mock{Data: tranByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	//Act
	misTransacciones, _ := repo.GetAll()

	//Assert
	assert.Equal(t, transactionsExpected, misTransacciones)
	assert.True(t, dbMock.Called)
}

func TestRepositoryUpdateMock(t *testing.T) {
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

	//Act
	miTransaccion, _ := repo.Update(5, "prueba1", "MXN", 123.45, "em1", "rec1", "12-12-2020")

	//Assert
	assert.Equal(t, transactionExpected, miTransaccion)
	assert.True(t, dbMock.Called)
}
