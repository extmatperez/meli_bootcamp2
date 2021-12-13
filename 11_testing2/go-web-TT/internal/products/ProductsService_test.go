package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/11_testing2/go-web-TT/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubResposirity struct {
	useGetAll bool
}

var productss string = `[
		{"id": 1, "nombre": "producto1modificadoconpatch","color": "rojo", "precio": 20, "stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"},
   		{"id": 2,"nombre": "producto1","color": "rojo","precio": 20,"stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"
	}]`

func (s *StubResposirity) GetAll() ([]Product, error) {
	var salida []Product
	err := json.Unmarshal([]byte(productss), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubResposirity) Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) UpdateNombre(id int, nombre string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) UpdatePrecio(id int, precio int) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) Delete(id int) error {
	return nil
}
func (s *StubResposirity) LastID() (int, error) {
	return 0, nil
}
func TestGetAllService(t *testing.T) {
	stubRepo := StubResposirity{}
	service := NewService(&stubRepo)
	misProductos, _ := service.GetAll()

	assert.Equal(t, 2, len(misProductos))
	assert.True(t, true, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) { // implementa el delete

	stubRepo := StubResposirity{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}
func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(productss)
	var productosEsperados []Product
	json.Unmarshal(dataByte, &productosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misPersonas, _ := service.GetAll()

	assert.Equal(t, productosEsperados, misPersonas)
}

func TestGetAllServiceMockError(t *testing.T) {
	//Arrange
	errorEsperado := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misProductos, errorRecibido := service.GetAll()

	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, misProductos)
}
func TestStoreServiceMock(t *testing.T) {
	//Arrange
	productoNuevo := Product{
		Nombre:        "producto test",
		Color:         "rojo",
		Precio:        20,
		Stock:         "alguno",
		Codigo:        "SADFHJK9",
		Publicado:     true,
		FechaCreacion: "01/12/2021",
	}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoCreado, _ := service.Store(productoNuevo.Nombre, productoNuevo.Color, productoNuevo.Precio, productoNuevo.Stock, productoNuevo.Codigo, productoNuevo.Publicado, productoNuevo.FechaCreacion)

	assert.Equal(t, productoNuevo.Nombre, productoCreado.Nombre)
	assert.Equal(t, productoNuevo.Color, productoCreado.Color)
	// assert.Nil(t, misPersonas)
}
func TestStoreServiceMockError(t *testing.T) {
	//Arrange
	productoNuevo := Product{
		Nombre:        "producto test",
		Color:         "rojo",
		Precio:        20,
		Stock:         "alguno",
		Codigo:        "SADFHJK9",
		Publicado:     true,
		FechaCreacion: "01/12/2021",
	}
	errorEsperado := errors.New("No hay datos en el Mock")
	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productoCreado, err := service.Store(productoNuevo.Nombre, productoNuevo.Color, productoNuevo.Precio, productoNuevo.Stock, productoNuevo.Codigo, productoNuevo.Publicado, productoNuevo.FechaCreacion)

	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, Product{}, productoCreado)
	// assert.Nil(t, misPersonas)
}
