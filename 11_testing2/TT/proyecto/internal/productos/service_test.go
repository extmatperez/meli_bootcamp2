/*
Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
	Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
	Para dar el test como OK debe validarse que al invocar el método del Service Update(), retorne el producto
	con mismo Id y los datos actualizados. Validar también que  Read() del Repository haya sido ejecutado durante el test.
*/

/*
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la correcta eliminación de un
producto, y el error cuando el producto no existe. Para lograrlo puede:
	Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
	Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete().
	También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.
*/

package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TT/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSeriveMock(t *testing.T) {

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

func TestUpdateSeriveMockError(t *testing.T) {

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

func TestDeleteSeriveMock(t *testing.T) {

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

func TestDeleteSeriveMockError(t *testing.T) {

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