package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/11_testing2/TM/Ejercicios/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

type MockStore struct {
	CalledMethod bool
}

var pays string = `[
	{
		"id": 1,
		"codigo": "0000001",
		"moneda": "ARS",
		"monto": 10535.26,
		"emisor": "Rodrigo Vega Gimenez",
		"receptor": "Juan Pablo Nieto",
		"fecha": "2021-12-01"
	   },
	   {
		"id": 2,
		"codigo": "0000002",
		"moneda": "ARS",
		"monto": 9563.45,
		"emisor": "Rodrigo Vega Gimenez",
		"receptor": "Maximiliano Caceres",
		"fecha": "2021-12-01"
	   }]`

var paymentTestCheckings []Payment = []Payment{
	{Id: 1, Codigo: "0000001", Moneda: "ARS", Monto: 10535.26, Emisor: "Rodrigo Vega Gimenez", Receptor: "Juan Pablo Nieto", Fecha: "2021-12-01"},
	{Id: 2, Codigo: "0000002", Moneda: "ARS", Monto: 9563.45, Emisor: "Rodrigo Vega Gimenez", Receptor: "Maximiliano Caceres", Fecha: "2021-12-01"},
}

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(pays), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (m *MockStore) Read(data interface{}) error {
	m.CalledMethod = true
	return json.Unmarshal([]byte(pays), &data)
}

func (m *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	pays = string(byteData)
	return nil
}

// Creamos los stores Stub y Mock.
func NewStubStore() store.Store {
	return &StubStore{}
}

// Aqui por defecto la llamada al metodo va a ser false, porque inicialmente no se corrio ese metodo y el mock debe devolver false.
func NewMockStore() store.Store {
	return &MockStore{false}
}

func TestGetAll(t *testing.T) {
	// Arrange.
	store := StubStore{}
	repo := NewRepository(&store)

	// Act.
	myPayments, _ := repo.GetAll()

	// Esto queda para otro test, tipo TestGetAllError.
	// assert.Nil(t, err)

	var expected []Payment
	json.Unmarshal([]byte(pays), &expected)

	// Assert.
	assert.Equal(t, expected, myPayments)
}

func TestGetAllError(t *testing.T) {
	// Arrange.
	store := StubStore{}
	repo := NewRepository(&store)

	// Act.
	_, err := repo.GetAll()

	assert.Nil(t, err)
}

func TestLastId(t *testing.T) {
	// Arrange.
	store := StubStore{}
	repo := NewRepository(&store)
	lastId := 2

	// Act.
	ultimoId, _ := repo.LastId()

	// Assert.
	assert.Equal(t, lastId, ultimoId)
}

func TestLastIdError(t *testing.T) {
	// Arrange.
	store := StubStore{}
	repo := NewRepository(&store)

	// Act.
	_, err := repo.LastId()

	// Assert.
	assert.Nil(t, err)
}

func TestUpdateCode(t *testing.T) {
	mockStore := MockStore{}
	repo := NewRepository(&mockStore)
	pay, _ := repo.UpdateCodigo(1, "000001")

	assert.Equal(t, pay.Codigo, "000001", "Se tiene que actualizar.")
	assert.True(t, mockStore.CalledMethod)
}

func TestUpdateCodeError(t *testing.T) {
	mockStore := MockStore{}
	repo := NewRepository(&mockStore)
	_, err := repo.UpdateCodigo(1, "000001")

	assert.Nil(t, err)
}

func TestGetAllRepositoryMock(t *testing.T) {
	dataBytes := []byte(pays)
	var expectedPayments []Payment
	json.Unmarshal(dataBytes, &expectedPayments)

	dbMock := store.Mock{Data: dataBytes}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	myPayments, _ := repo.GetAll()

	assert.Equal(t, expectedPayments, myPayments)
}

func TestGetLastIdRepositoryMock(t *testing.T) {
	dataBytes := []byte(pays)
	var expectedPayments []Payment
	json.Unmarshal(dataBytes, &expectedPayments)

	dbMock := store.Mock{Data: dataBytes}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	ultimoId, _ := repo.LastId()

	assert.Equal(t, expectedPayments[len(expectedPayments)-1].Id, ultimoId)
}
