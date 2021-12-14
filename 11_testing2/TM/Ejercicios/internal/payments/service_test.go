package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/11_testing2/TM/Ejercicios/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

var pays_bis string = `[
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

func (s *StubRepository) GetAll() ([]Payment, error) {
	var out []Payment
	err := json.Unmarshal([]byte(pays), &out)
	s.useGetAll = true
	return out, err
}

func (s *StubRepository) Filtrar(values ...string) ([]Payment, error) {
	return []Payment{}, nil
}

func (s *StubRepository) Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) Update(id int, codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) UpdateCodigo(id int, codigo string) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) UpdateMonto(id int, monto float64) (Payment, error) {
	return Payment{}, nil
}

func (s *StubRepository) Delete(id int) (string, error) {
	return "", nil
}

func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

/*func TestGetAllService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	myPayments, _ := service.GetAll()

	assert.Equal(t, 2, len(myPayments))
	assert.True(t, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	// Para que en caso de borrar el primer Id entonces tire un error.
	_, err := service.Delete(1)

	assert.Nil(t, err)
}*/

func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(pays_bis)
	var expectedPayments []Payment
	json.Unmarshal(dataByte, &expectedPayments)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myPayments, _ := service.GetAll()

	assert.Equal(t, expectedPayments, myPayments)
}

func TestGetAllServiceMockError(t *testing.T) {
	//Arrange
	expectedError := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myPayments, receivedError := service.GetAll()

	assert.Equal(t, expectedError, receivedError)
	assert.Nil(t, myPayments)
}

func TestStoreServiceMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	createdPayment, _ := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, newPayment.Codigo, createdPayment.Codigo)
	assert.Equal(t, newPayment.Emisor, createdPayment.Emisor)
	assert.Equal(t, newPayment.Receptor, createdPayment.Receptor)
}

func TestStoreServiceMockError(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	expectedError := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	createdPayment, err := service.Store(newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, Payment{}, createdPayment)
}

func TestUpdateServiceMock(t *testing.T) {
	//Arrange
	newPayment := Payment{Codigo: "0000003", Moneda: "ARS", Monto: 674.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Carlos Miño", Fecha: "2021-12-02"}

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	updatedPayment, _ := service.Update(1, newPayment.Codigo, newPayment.Moneda, newPayment.Emisor, newPayment.Receptor, newPayment.Fecha, newPayment.Monto)

	assert.Equal(t, newPayment.Codigo, updatedPayment.Codigo)
	assert.Equal(t, newPayment.Emisor, updatedPayment.Emisor)
	assert.Equal(t, newPayment.Receptor, updatedPayment.Receptor)
	assert.Equal(t, 1, updatedPayment.Id)
}

func TestUpdateNombreServiceMock(t *testing.T) {
	//Arrange
	newCode := "0000004"

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	updatedPayment, _ := service.UpdateCodigo(2, newCode)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, newCode, updatedPayment.Codigo)
	assert.Equal(t, 2, updatedPayment.Id)
	// assert.Nil(t, misPersonas )
}

func TestDeleteNombreServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(pays_bis)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Delete(2)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Nil(t, err)

	allPayments, _ := service.GetAll()

	assert.Equal(t, 1, len(allPayments))
	// assert.Nil(t, misPersonas)
}
