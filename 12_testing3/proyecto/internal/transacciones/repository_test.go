package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

type MockStore struct {
	used bool
}

var trans string = `[
{"id":1,"codigo_transaccion":"cybqf7i0bo","moneda":"Baht","monto":3011534.4,"emisor":"aburdis0","receptor":"fodoireidh0","fecha_creacion":"9/25/2021"},
{"id":2,"codigo_transaccion":"w9fm5sk8gt","moneda":"Yuan Renminbi","monto":5069586.0,"emisor":"lfilyushkin1","receptor":"nwestmancoat1","fecha_creacion":"4/24/2021"}
]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(trans), &data)
}
func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (m *MockStore) Read(data interface{}) error {
	m.used = true
	return json.Unmarshal([]byte(trans), &data)
}

func (m *MockStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//act
	misTransacciones, err := repo.GetAll()
	var expected []Transaccion
	json.Unmarshal([]byte(trans), &expected)

	//assert
	assert.Equal(t, expected, misTransacciones)
	assert.Nil(t, err)
}

func TestUpdateEmisor(t *testing.T) {
	//arrange
	store := MockStore{false}
	repo := NewRepository(&store)

	nameExpected := "after update"

	//act
	result, err := repo.Update(1, "cybqf7i0bo", "Baht", 3011534.4, nameExpected, "fodoireidh0", "9/25/2021")

	//assert
	assert.Equal(t, nameExpected, result.Emisor)
	assert.Nil(t, err)
	assert.True(t, store.used)

}

func TestUpdateError(t *testing.T) {
	//arrange
	store := MockStore{false}
	repo := NewRepository(&store)

	nameExpected := "after update"

	//act
	_, err := repo.Update(90, "cybqf7i0bo", "Baht", 3011534.4, nameExpected, "fodoireidh0", "9/25/2021")

	//assert
	assert.Error(t, err)

}

func TestUpdateRepositoryMock(t *testing.T) {
	//arrange
	dataByte := []byte(trans)
	dbMock := store.Mock{Data: dataByte}

	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)

	nameExpected := "after update"

	//act
	result, _ := repo.Update(1, "cybqf7i0bo", "Baht", 3011534.4, nameExpected, "fodoireidh0", "9/25/2021")

	//assert
	assert.Equal(t, nameExpected, result.Emisor)

}
