package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TT/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceMock(t *testing.T) {
	//Arrenge
	newProduct := Producto{
		ID:            1,
		Nombre:        "Asdf",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "01/06/1996",
	}

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	createdProduct, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Precio, newProduct.Stock, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion)

	//Assert
	assert.Equal(t, newProduct.Nombre, createdProduct.Nombre)
	assert.Equal(t, newProduct.Codigo, createdProduct.Codigo)
	assert.Nil(t, err)
	assert.True(t, dbMock.Used)
}

func TestCreateServiceMockError(t *testing.T) {
	//Arrenge
	newProduct := Producto{
		ID:            1,
		Nombre:        "Asdf",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "01/06/1996",
	}

	expectedError := errors.New("error en los datos del Mock")

	dbMock := store.Mock{Data: []byte(productosTest), Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	_, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Precio, newProduct.Stock, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion)

	//Assert
	assert.Equal(t, expectedError, err)
	assert.True(t, dbMock.Used)
}

func TestGetAllMock(t *testing.T) {
	//Arrenge
	dbMock := store.Mock{Data: []byte(productosTest)}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	var expectedList []Producto
	errUnmarshal := json.Unmarshal([]byte(productosTest), &expectedList)

	//Act
	productsList, err := service.GetAll()

	//Assert
	assert.Equal(t, expectedList, productsList)
	assert.Nil(t, err)
	assert.Nil(t, errUnmarshal)
	assert.True(t, dbMock.Used)
}

func TestGetAllMockError(t *testing.T) {
	//Arrenge
	expectedError := errors.New("The Mock is empty")

	dbMock := store.Mock{Data: []byte(``), Err: expectedError}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	_, err := service.GetAll()

	//Assert
	assert.Equal(t, expectedError, err)
	assert.True(t, dbMock.Used)
}

func TestUpdateServiceMock(t *testing.T) {
	//Arrenge
	productToUpdate := Producto{
		ID:            1,
		Nombre:        "Asdf",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "01/06/1996",
	}

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	updatedProduct, err := service.Update(productToUpdate.ID, productToUpdate.Nombre, productToUpdate.Color, productToUpdate.Precio, productToUpdate.Stock, productToUpdate.Codigo, productToUpdate.Publicado, productToUpdate.FechaCreacion)

	//Assert
	assert.Equal(t, productToUpdate, updatedProduct)
	assert.Nil(t, err)
	assert.True(t, dbMock.Used)
}

func TestUpdateServiceMockError(t *testing.T) {
	//Arrenge
	productToUpdate := Producto{
		ID:            0,
		Nombre:        "Asdf",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "01/06/1996",
	}

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	_, err := service.Update(productToUpdate.ID, productToUpdate.Nombre, productToUpdate.Color, productToUpdate.Precio, productToUpdate.Stock, productToUpdate.Codigo, productToUpdate.Publicado, productToUpdate.FechaCreacion)

	//Assert
	assert.Error(t, err)
	assert.True(t, dbMock.Used)
}

func TestDeleteServiceMock(t *testing.T) {
	//Arrenge
	productToUpdateID := 1

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	err := service.Delete(productToUpdateID)

	//Assert
	assert.Nil(t, err)
	assert.True(t, dbMock.Used)
}

func TestDeleteServiceMockError(t *testing.T) {
	//Arrenge
	productToUpdateID := 0

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	err := service.Delete(productToUpdateID)

	//Assert
	assert.Error(t, err)
	assert.True(t, dbMock.Used)
}

func TestUpdateNombrePrecioServiceMock(t *testing.T) {
	//Arrenge
	productToUpdateID := 1
	productToUpdateNombre := "Coca"
	productToUpdatePrecio := "$150"

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	updatedProduct, err := service.UpdateNombrePrecio(productToUpdateID, productToUpdateNombre, productToUpdatePrecio)

	//Assert
	assert.Equal(t, productToUpdateID, updatedProduct.ID)
	assert.Equal(t, productToUpdateNombre, updatedProduct.Nombre)
	assert.Equal(t, productToUpdatePrecio, updatedProduct.Precio)
	assert.Nil(t, err)
	assert.True(t, dbMock.Used)
}

func TestUpdateNombrePrecioServiceMockError(t *testing.T) {
	//Arrenge
	productToUpdateID := 0
	productToUpdateNombre := "Coca"
	productToUpdatePrecio := "$150"

	dbMock := store.Mock{Data: []byte(productosTest)}

	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	//Act
	_, err := service.UpdateNombrePrecio(productToUpdateID, productToUpdateNombre, productToUpdatePrecio)

	//Assert
	assert.Error(t, err)
	assert.True(t, dbMock.Used)
}