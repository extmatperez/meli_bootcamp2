package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllServiceMock(t *testing.T) {

	dataByte := []byte(produc)                           //* obtengo los bytes de product
	var productosEsperados []Productos                   //* declaro una variable
	err := json.Unmarshal(dataByte, &productosEsperados) //* guardo los datos en la variable
	if err != nil {
		assert.Error(t, err)
	}

	dbMock := store.Mock{Data: dataByte}    //* paso los datos en bytes
	store := store.FileStore{Mock: &dbMock} //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)           //* Creo un repo y le paso mi store que tiene todos los metodos

	service := NewService(repo) //? creo un service con el repo

	productos, _ := service.GetAll() //* devuelve los productos y un error

	assert.Equal(t, productosEsperados, productos)

}

func TestGetAllServiceErrMock(t *testing.T) {

	errorEsperado := errors.New("error: No hay datos en el Mock")

	dbMock := store.Mock{Err: errorEsperado} //* paso los datos en bytes
	store := store.FileStore{Mock: &dbMock}  //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)            //* Creo un repo y le paso mi store que tiene todos los metodos

	service := NewService(repo) //? creo un service con el repo

	productos, errorRecibido := service.GetAll() //* devuelve los productos y un error

	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, productos)

}

func TestStoreServiceMock(t *testing.T) {

	prodNuevo := Productos{
		Nombre:            "prod nuevo",
		Color:             "verde",
		Precio:            150.2,
		Stock:             5,
		Codigo:            "ddasd",
		Publicado:         true,
		Fecha_de_creacion: "10/10/10",
	}

	dbMock := store.Mock{Data: []byte(`[]`)} //* creo el db mock y le paso un json vacio
	store := store.FileStore{Mock: &dbMock}  //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)            //* Creo un repo y le paso mi store que tiene todos los metodos

	service := NewService(repo) //? creo un service con el repo

	producto, _ := service.Store(prodNuevo.Stock, prodNuevo.Nombre, prodNuevo.Color, prodNuevo.Codigo,
		prodNuevo.Fecha_de_creacion, prodNuevo.Precio, prodNuevo.Publicado) //* devuelve el producto y un error

	assert.Equal(t, prodNuevo.Nombre, producto.Nombre)
	assert.Equal(t, prodNuevo.Codigo, producto.Codigo)
}

func TestStoreServiceMockErr(t *testing.T) {

	prodNuevo := Productos{
		Nombre:            "prod nuevo",
		Color:             "verde",
		Precio:            150.2,
		Stock:             5,
		Codigo:            "ddasd",
		Publicado:         true,
		Fecha_de_creacion: "10/10/10",
	}

	errorEsperado := errors.New("error: No hay datos en el Mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado} //* creo el db mock y le paso un json vacio y un err
	store := store.FileStore{Mock: &dbMock}                      //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)                                //* Creo un repo y le paso mi store que tiene todos los metodos

	service := NewService(repo) //? creo un service con el repo

	producto, err := service.Store(prodNuevo.Stock, prodNuevo.Nombre, prodNuevo.Color, prodNuevo.Codigo,
		prodNuevo.Fecha_de_creacion, prodNuevo.Precio, prodNuevo.Publicado) //* devuelve el producto y un error

	assert.Equal(t, errorEsperado, err)

	assert.Equal(t, Productos{}, producto)
}
