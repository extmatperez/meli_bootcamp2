package internal

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/12_testing3/PracticaTM/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/internal/models"
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

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	transaccionNueva := models.Transaccion{
		Moneda:   "Pesos",
		Monto:    65002.45,
		Emisor:   "Facundo",
		Receptor: "Matias",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionCreada, _ := service.Store(transaccionNueva.Moneda, transaccionNueva.Monto, transaccionNueva.Emisor, transaccionNueva.Receptor)

	assert.Equal(t, transaccionNueva.Moneda, transaccionCreada.Moneda)
	assert.Equal(t, transaccionNueva.Monto, transaccionCreada.Monto)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	transaccionNueva := models.Transaccion{
		Moneda:   "Pesos",
		Monto:    6500.45,
		Emisor:   "Facundo",
		Receptor: "Matias",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionCargada := service.GetOne(1)

	assert.Equal(t, transaccionNueva.Moneda, transaccionCargada.Moneda)
	assert.Equal(t, transaccionNueva.Monto, transaccionCargada.Monto)
	// assert.Nil(t, misPersonas)
}

func TestGetByNameSQL(t *testing.T) {
	//Arrange
	name1 := "Facundo"
	name2 := "Rebeca"

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionesFacundo := service.GetByName(name1)
	transaccionesRebeca := service.GetByName(name2)

	assert.Equal(t, len(transaccionesFacundo), 3)
	assert.Equal(t, len(transaccionesRebeca), 2)
}

func TestGetAllServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misPersonasDB, err := service.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(misPersonasDB) > 0)
}

func TestDeleteServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionNueva := models.Transaccion{
		Moneda:   "Pesos",
		Monto:    6500.45,
		Emisor:   "Facundo",
		Receptor: "Matias",
	}

	transaccionCreada, _ := service.Store(transaccionNueva.Moneda, transaccionNueva.Monto, transaccionNueva.Emisor, transaccionNueva.Receptor)

	idTransaccion, _ := strconv.Atoi(transaccionCreada.CodTransaccion)
	err := service.Delete(idTransaccion)

	assert.Nil(t, err)
}

func TestDeleteNotFoundServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	err := service.Delete(0)

	assert.Error(t, err)
	assert.Equal(t, "no se encontro la transaccion", err.Error())
}

func TestGetFullDataServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misPersonasDB, err := service.GetFullData()

	assert.Nil(t, err)
	assert.Equal(t, "Cordoba", misPersonasDB[1].Ciudad.Nombre)
}

func TestGetOneContextServiceSQL(t *testing.T) {
	//Arrange
	transaccionNueva := models.Transaccion{
		Moneda:   "Pesos",
		Monto:    6500.45,
		Emisor:   "Facundo",
		Receptor: "Matias",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	transaccionCargada, err := service.GetOneWithContext(context.Background(), 1)

	assert.Equal(t, transaccionNueva.Moneda, transaccionCargada.Moneda)
	assert.Equal(t, transaccionNueva.Monto, transaccionCargada.Monto)
	assert.Nil(t, err)
}
