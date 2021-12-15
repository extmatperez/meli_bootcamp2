package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

var produs string = `[ 
	{	"id": 1,	"nombre": "Banana",		"color": "Amarillo",	"precio": 27.99 , "stock" : 100, "codigo":"300","publicado": true,"fecha_de_creacion":"22/04/1991"  },
	{	"id": 2,	"nombre": "Manzana",	"color": "Rojo",	"precio": 17.99 , "stock" : 50, "codigo":"210","publicado": true,"fecha_de_creacion":"22/05/1995"  }]`

func (s *StubRepository) GetAll() ([]Product, error) {
	var salida []Product
	err := json.Unmarshal([]byte(produs), &salida)
	s.useGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubRepository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubRepository) UpdateProd(id int, name string, price float64) (Product, error) {
	return Product{}, nil
}
func (s *StubRepository) Delete(id int) error {
	return nil
}
func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misProducts, _ := service.GetAll()

	assert.Equal(t, 2, len(misProducts))
	assert.True(t, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}
func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(produs)
	var productosEsperados []Product
	json.Unmarshal(dataByte, &productosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misProducts, _ := service.GetAll()

	assert.Equal(t, productosEsperados, misProducts)
}

func TestGetAllServiceMockError(t *testing.T) {
	//Arrange
	errorEsperado := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misProducts, errorRecibido := service.GetAll()

	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, misProducts)
}

func TestStoreServiceMock(t *testing.T) {
	//Arrange
	produNuevo := Product{
		Nombre:          "Banana",
		Color:           "Naranja",
		Precio:          29.99,
		Stock:           35,
		Codigo:          "banaro",
		Publicado:       true,
		FechaDeCreacion: "22/11/21",
	}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoCreado, _ := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.Equal(t, produNuevo.Nombre, productoCreado.Nombre)
	assert.Equal(t, produNuevo.Color, productoCreado.Color)
	// assert.Nil(t, misProducts)
}

func TestStoreServiceMockError(t *testing.T) {
	//Arrange
	produNuevo := Product{
		Nombre:          "Manzana",
		Color:           "verde",
		Precio:          29.99,
		Stock:           35,
		Codigo:          "mave",
		Publicado:       true,
		FechaDeCreacion: "22/11/21",
	}
	errorEsperado := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoCreado, err := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, Product{}, productoCreado)
}

func TestUpdateServiceMock(t *testing.T) {
	//Arrange
	produNuevo := Product{
		Nombre:          "Manzana",
		Color:           "verde",
		Precio:          29.99,
		Stock:           35,
		Codigo:          "mave",
		Publicado:       true,
		FechaDeCreacion: "22/11/21",
	}

	dataByte := []byte(produs)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoActualizado, _ := service.Update(1, produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)
	//productoCreado, err := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.Equal(t, produNuevo.Nombre, produNuevo.Nombre)
	assert.Equal(t, produNuevo.Color, produNuevo.Color)
	assert.Equal(t, 1, productoActualizado.ID)
	// assert.Nil(t, misProducts)
}

func TestUpdateProdServiceMock(t *testing.T) {
	//Arrange
	nuevoNombre := "Melon"
	nuevoPrecio := 39.99
	dataByte := []byte(produs)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoActualizado, _ := service.UpdateProd(2, nuevoNombre, nuevoPrecio)
	//productoCreado, err := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.Equal(t, nuevoNombre, productoActualizado.Nombre)
	assert.Equal(t, 2, productoActualizado.ID)
	// assert.Nil(t, misProducts )
}

func TestUpdateProdServiceMockError(t *testing.T) {
	//Arrange
	nuevoNombre := "Sandia"
	nuevoPrecio := 39.99

	dataByte := []byte(produs)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateProd(15, nuevoNombre, nuevoPrecio)
	//productoCreado, err := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.NotNil(t, err)
	// assert.Nil(t, misProducts )
}

func TestDeleteNombreServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(produs)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(2)
	//productoCreado, err := service.Store(produNuevo.Nombre, produNuevo.Color, produNuevo.Precio, produNuevo.Stock, produNuevo.Codigo, produNuevo.Publicado, produNuevo.FechaDeCreacion)

	assert.Nil(t, err)

	todos, _ := service.GetAll()

	assert.Equal(t, 1, len(todos))
	// assert.Nil(t, misProducts)
}
