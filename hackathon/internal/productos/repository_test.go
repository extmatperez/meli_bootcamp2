package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/internal/models"
	"github.com/stretchr/testify/assert"
)

// var prods string = `[{"id": 1,"nombre": "BeforeUpdate", "color": "azul", "precio": 1.8},
// 	{
// 	 "id": 2,
// 	 "nombre": "poyitrdcvh",
// 	 "color": "dorado",
// 	 "precio": 1.8
// 	}]`
// var productosInstancias []Producto = []Producto{
// 	{Id: 1, Nombre: "BeforeUpdate", Color: "azul", Precio: 1.8},
// 	{Id: 2, Nombre: "poyitrdcvh", Color: "dorado", Precio: 1.8},
// }

// type MockStore struct {
// 	MetodoLlamado bool
// }
// type StubStore struct{}

// func (s *MockStore) Read(data interface{}) error {
// 	s.MetodoLlamado = true
// 	return json.Unmarshal([]byte(prods), &data)
// }
// func (s *StubStore) Read(data interface{}) error {
// 	return json.Unmarshal([]byte(prods), &data)
// }

// func (s *StubStore) Write(data interface{}) error {
// 	return nil
// }
// func (s *MockStore) Write(data interface{}) error {
// 	byteData, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	prods = string(byteData)
// 	return nil
// }

// func NewStubStore() store.Store {
// 	return &StubStore{}
// }
// func NewMockStore() store.Store {
// 	return &MockStore{false}
// }

// func TestGetAll(t *testing.T) {
// 	productosEsperados := productosInstancias
// 	stubStore := NewStubStore()
// 	repo := NewRepository(stubStore)
// 	resultadoGetAll, err := repo.GetAll()

// 	assert.Equal(t, productosEsperados, resultadoGetAll, "en el get all los productos tendiran que ser iguales ")
// 	assert.Nil(t, err, "que no devuelva un error ")
// }

// func TestUpdate(t *testing.T) {
// 	mockStore := MockStore{}
// 	repo := NewRepository(&mockStore)
// 	prod, _ := repo.UpdateNombrePrecio(1, "After Update", 20.1)

// 	assert.Equal(t, prod.Nombre, "After Update", "se tiene que actualizar")
// 	assert.True(t, mockStore.MetodoLlamado)

// }

// var productosParaMock []Producto = []Producto{
// 	{Id: 1, Nombre: "BeforeUpdate", Color: "azul", Precio: 1.8},
// 	{Id: 2, Nombre: "poyitrdcvh", Color: "dorado", Precio: 1.8},
// }

// func TestDeleteService(t *testing.T) {
// 	prodByte, _ := json.Marshal(productosParaMock)
// 	mock := store.Mock{Data: prodByte}
// 	fileStore := store.FileStore{Mock: &mock}
// 	repo := NewRepository(&fileStore)
// 	service := NewService(repo)

// 	err := service.Delete(1)

// 	assert.Nil(t, err, "se tendria que borrar exitosamente")

// 	productosPostBorrado, _ := service.GetAll()
// 	assert.Equal(t, len(productosParaMock)-1, len(productosPostBorrado), "la longitud tendria que ser menor despues del borrado")

// }

// func TestUpdateNombrePrecioService(t *testing.T) {
// 	prodByte, _ := json.Marshal(productosParaMock)
// 	fmt.Println("PRODUCTOS PARA MOCK : ", productosParaMock)
// 	mock := store.Mock{Data: prodByte}
// 	fileStore := store.FileStore{Mock: &mock}
// 	repo := NewRepository(&fileStore)
// 	service := NewService(repo)

// 	productoActualizado, err := service.UpdateNombrePrecio(1, "Nuevo", 22.2)
// 	assert.Equal(t, productoActualizado.Nombre, "Nuevo", "el nombre del rpoducto actualizado debe ser igual al parametro con el que se lo llama")
// 	assert.Nil(t, err, "no debe arrojar ningun error el actualizado")
// }

var prod models.Producto = models.Producto{Nombre: "tenedor", Color: "azul", Precio: 22.4}

// func TestStoreSQL(t *testing.T) {
// 	repoSQL := NewRepositorySQL()

// 	productoResultado, err := repoSQL.Store(prod)

// 	assert.Equal(t, prod.Nombre, productoResultado.Nombre)
// 	assert.Equal(t, prod.Color, productoResultado.Color)
// 	assert.Nil(t, err)

// }

func TestGetAllSQL(t *testing.T) {
	repoSQL := NewRepositorySQL()

	sliceProductos, err := repoSQL.GetAll()

	assert.Nil(t, err, "no debe arrojar error")
	assert.True(t, len(sliceProductos) > 0)

}

// func TestUpdateContextSQL(t *testing.T) {
// 	repoSQL := NewRepositorySQL()
// 	productoParaActualizar := models.Producto{Id: 7, Nombre: "axtualizado", Color: "nuevo", Precio: 57.6}

// 	productoAntesDeActualizar, _ := repoSQL.GetById(productoParaActualizar.Id)

// 	contextNuevo := context.Background()
// 	productoActualizado, _ := repoSQL.UpdateContext(contextNuevo, productoParaActualizar)

// 	assert.Equal(t, productoActualizado.Nombre, productoParaActualizar.Nombre)
// 	assert.Equal(t, productoActualizado.Color, productoParaActualizar.Color)

// 	_, err := repoSQL.UpdateContext(contextNuevo, productoAntesDeActualizar)
// 	assert.Nil(t, err, "no debe arrojar error2 ")
// }

func TestGetFullData(t *testing.T) {
	repo := NewRepositorySQL()
	productos, err := repo.GetFullData()
	assert.NotNil(t, productos)
	assert.Nil(t, err)
}
