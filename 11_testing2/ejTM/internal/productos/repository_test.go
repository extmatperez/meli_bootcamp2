package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/11_testing2/ejTM/pkg/store"
	"github.com/stretchr/testify/assert"
)

var prods string = `[{"id": 1,"nombre": "BeforeUpdate", "color": "azul", "precio": 1.8},
	{
	 "id": 2,
	 "nombre": "poyitrdcvh",
	 "color": "dorado",
	 "precio": 1.8
	}]`
var productosInstancias []Producto = []Producto{
	{Id: 1, Nombre: "BeforeUpdate", Color: "azul", Precio: 1.8},
	{Id: 2, Nombre: "poyitrdcvh", Color: "dorado", Precio: 1.8},
}

type MockStore struct {
	MetodoLlamado bool
}
type StubStore struct{}

// type request struct {
// 	Nombre string  `json:"nombre"`
// 	Color  string  `json:"color" `
// 	Precio float64 `json:"precio" `
// }
func (s *MockStore) Read(data interface{}) error {
	s.MetodoLlamado = true
	return json.Unmarshal([]byte(prods), &data)
}
func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(prods), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}
func (s *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	prods = string(byteData)
	return nil
}

func NewStubStore() store.Store {
	return &StubStore{}
}
func NewMockStore() store.Store {
	return &MockStore{false}
}

func TestGetAll(t *testing.T) {
	productosEsperados := productosInstancias
	stubStore := NewStubStore()
	repo := NewRepository(stubStore)
	resultadoGetAll, err := repo.GetAll()

	assert.Equal(t, productosEsperados, resultadoGetAll, "en el get all los productos tendiran que ser iguales ")
	assert.Nil(t, err, "que no devuelva un error ")
}

func TestUpdate(t *testing.T) {
	mockStore := MockStore{}
	repo := NewRepository(&mockStore)
	prod, _ := repo.UpdateNombrePrecio(1, "After Update", 20.1)

	assert.Equal(t, prod.Nombre, "After Update", "se tiene que actualizar")
	assert.True(t, mockStore.MetodoLlamado)

}
