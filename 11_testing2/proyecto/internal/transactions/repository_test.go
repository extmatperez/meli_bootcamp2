package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

var tr string = `[
	{"id": 5,"codigo_de_transaccion": "asd12321","moneda": "USD","monto": 23400.85,"emisor": "sebac2","receptor": "rec","fecha_de_transaccion": "10/18/2021"},
	{"id": 6,"codigo_de_transaccion": "asd12321","moneda": "UYU","monto": 2344.8,"emisor": "sebac","receptor": "rec","fecha_de_transaccion": "10/18/2021"}
	]`

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
	updatedTransaction, _ := repo.UpdateCodigoYMonto(5, "afterupdate", 200.0)
	var expected Transaction
	json.Unmarshal([]byte(`{"id": 5,"codigo_de_transaccion": "afterupdate","moneda": "USD","monto": 200.0,"emisor": "sebac2","receptor": "rec","fecha_de_transaccion": "10/18/2021"}`), &expected)
	//Assert
	assert.Equal(t, expected, updatedTransaction)
	assert.True(t, store.called)
}
