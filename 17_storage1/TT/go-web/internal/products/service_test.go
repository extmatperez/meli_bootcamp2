package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 1, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	prod, _ := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)

	assert.Equal(t, prodEsperado, prod)

}

func TestUpdateNotFoundMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 4, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	_, err := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.Error(t, err)
}

func TestUpdateErrorMock(t *testing.T) {
	dbMock := store.Mock{Err: errors.New("Error en la conexion")}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := Product{ID: 4, Nombre: "Change Name", Color: "Change Color", Precio: 2, Stock: 2, Codigo: "My code", Publicado: true, FechaCreacion: "01/12/2021"}
	_, err := service.Update(prodEsperado.ID, prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.NotNil(t, err)
}

func TestDeleteMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	err := service.Delete(1)
	assert.Nil(t, err)
	_, err = service.FindById(1)
	assert.NotNil(t, err)

}

func TestDeleteNotFoundMock(t *testing.T) {

	productsJsonTestUpdate := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestUpdate)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	err := service.Delete(4)
	assert.Error(t, err)
}

func TestDeleteErrorMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en la conexion")}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	err := service.Delete(4)
	assert.Error(t, err)
}

func TestGetAllMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := []Product{}
	json.Unmarshal([]byte(productsJsonTestGet), &prodEsperado)

	prod, _ := service.GetAll()
	assert.Equal(t, prodEsperado, prod)

}

func TestGetAllErrorMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en base de datos")}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.GetAll()
	assert.Error(t, err)

}

func TestGetAllEmptyMock(t *testing.T) {

	productsJsonTestGet := `[]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	prodEsperado := []Product{}
	json.Unmarshal([]byte(productsJsonTestGet), &prodEsperado)

	prod, _ := service.GetAll()
	assert.Equal(t, prodEsperado, prod)

}

func TestStoreMock(t *testing.T) {

	productsJsonTestGet := `[]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	prodEsperado := Product{ID: 1, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod, _ := service.Store(prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.Equal(t, prodEsperado, prod)

}
func TestStoreTwoElementsMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	prodEsperado := Product{ID: 3, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod, _ := service.Store(prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.Equal(t, prodEsperado, prod)

}

func TestStoreErrorMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en base de datos")}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	prodEsperado := Product{ID: 1, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	_, err := service.Store(prodEsperado.Nombre, prodEsperado.Color, prodEsperado.Precio, prodEsperado.Stock, prodEsperado.Codigo, prodEsperado.Publicado, prodEsperado.FechaCreacion)
	assert.Error(t, err)

}

func TestFindByIdMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	prodEsperado := Product{ID: 1, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod, _ := service.FindById(1)
	assert.Equal(t, prodEsperado, prod)
}

func TestFindByIdNotFoundMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.FindById(3)
	assert.Error(t, err)
}

func TestFindByIdErrMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en base de datos")}
	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.FindById(1)
	assert.Error(t, err)
}

func TestUpdateNameAndPriceMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	prodEsperado := Product{ID: 1, Nombre: "Update Name", Color: "string", Precio: 100, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod, _ := service.UpdateNameAndPrice(1, "Update Name", 100)
	assert.Equal(t, prodEsperado, prod)
}

func TestUpdateNameAndPriceNotFoundMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.UpdateNameAndPrice(3, "Update Name", 100)
	assert.Error(t, err)
}

func TestUpdateNameAndPriceErrordMock(t *testing.T) {

	dbMock := store.Mock{Err: errors.New("Error en base de datos")}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)

	_, err := service.UpdateNameAndPrice(3, "Update Name", 100)
	assert.Error(t, err)
}

func TestFilterMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	params := map[string][]string{"id": {"1"}, "nombre": {"Producto"}, "color": {"string"}, "precio": {"1"}, "stock": {"1"}, "codigo": {"string"}, "publicado": {"true"}, "fechaCreacion": {"string"}}
	prodEsperado := Product{ID: 1, Nombre: "Producto", Color: "string", Precio: 1, Stock: 1, Codigo: "string", Publicado: true, FechaCreacion: "string"}
	prod, _ := service.Filter(params)
	assert.Equal(t, prodEsperado.ID, prod[0].ID)
}

func TestFilterWithOutFiltersMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	params := map[string][]string{}
	prod, _ := service.Filter(params)
	assert.Equal(t, 2, len(prod))
}

func TestFilterPriceAndStockMock(t *testing.T) {

	productsJsonTestGet := `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":false,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":false,"fechaCreacion":"09/06/2021"}]`
	dbMock := store.Mock{Data: []byte(productsJsonTestGet)}

	db := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&db)
	service := NewService(repo)
	params := map[string][]string{"precio": {"0", "451"}, "stock": {"0", "100"}, "publicado": {"false"}}
	prod, _ := service.Filter(params)
	assert.Equal(t, 2, len(prod))
}
