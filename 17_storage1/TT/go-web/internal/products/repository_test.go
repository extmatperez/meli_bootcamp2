package internal

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/pkg/util"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	withError bool
}

var productsJsonTestGetAll string = `[{"id":1,"nombre":"Producto","color":"string","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`

func (s *StubStore) Read(data interface{}) error {
	if s.withError {
		return errors.New("Error de conexion a base de datos")
	}
	return json.Unmarshal([]byte(productsJsonTestGetAll), data)
}
func (s *StubStore) Write(data interface{}) error {
	return nil
}

type SpyStore struct {
	passRead bool
}

var productsJsonTestUpdateName string = `[{"id":1,"nombre":"Before Update","color":"Blue","precio":1,"stock":1,"codigo":"string","publicado":true,"fechaCreacion":"string"},{"id":2,"nombre":"Agenda - gris","color":"Indigo","precio":450,"stock":59,"codigo":"WBADX1C50CE880177","publicado":true,"fechaCreacion":"09/06/2021"}]`

func (s *SpyStore) Read(data interface{}) error {
	s.passRead = true
	return json.Unmarshal([]byte(productsJsonTestUpdateName), data)
}
func (s *SpyStore) Write(data interface{}) error {
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	productsJsonTestUpdateName = string(content)
	return nil
}

/*
func TestStoreInDataBase(t *testing.T) {
	repository := NewRepositorySQL(db.StorageDB)

	prod, err := repository.Store("Producto Nuevo", "Rojo", 1000, 10, "UK-0000", false, "2021-12-12")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, prod.ID)
}

func TestGetByNameInDataBase(t *testing.T) {
	repository := NewRepositorySQL(db.StorageDB)

	prod, err := repository.GetByName("Juan")
	assert.Nil(t, err)
	assert.Equal(t, 1, prod.ID)
}

func TestGetAllInDataBase(t *testing.T) {
	repository := NewRepositorySQL(db.StorageDB)

	prods, err := repository.GetAll()
	assert.Nil(t, err)
	assert.True(t, len(prods) > 0)
}

func TestUpdateInDataBase(t *testing.T) {
	repository := NewRepositorySQL(db.StorageDB)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prods, err := repository.Update(ctx, 2, "Change Name", "Change Color", 100, 1, "Change CODE", true, "2020/10/01")
	assert.Nil(t, err)
	assert.Equal(t, prods.Nombre, "Change Name")
	assert.Equal(t, prods.Color, "Change Color")
	assert.Equal(t, prods.Precio, 100)
	assert.Equal(t, prods.Stock, 1)
	assert.Equal(t, prods.Codigo, "Change CODE")
	assert.Equal(t, prods.Publicado, true)
	assert.Equal(t, prods.FechaCreacion, "2020/10/01")

	_, err = repository.Update(ctx, 2, "Update Product", "RED AND WHITE", 1001, 100, "UK-9087600", true, "2020/10/01")
	assert.Nil(t, err)

}*/

// -------------Test con go-txdb

func TestStoreInDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prod, err := repository.Store("Producto Nuevo", "Rojo", 1000, 10, "UK-0000", false, "2021-12-12 00:00:00")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, prod.ID)

	prodGetOne, err := repository.GetOne(prod.ID)
	assert.Nil(t, err)
	assert.Equal(t, prod, prodGetOne)
}

func TestGetOneDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prod, err := repository.GetOne(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, prod.ID)
}

func TestGetOneDataBaseMockNotFound(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prod, err := repository.GetOne(-1)
	assert.Error(t, err)
	assert.Equal(t, models.Product{}, prod)
}

func TestGetByNameInDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prod, err := repository.GetByName("Juan")
	assert.Nil(t, err)
	assert.Equal(t, 1, prod.ID)
}

func TestGetAllInDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prods, err := repository.GetAll()
	assert.Nil(t, err)
	assert.True(t, len(prods) > 0)
}

func TestUpdateInDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prods, err := repository.Update(ctx, 2, "Change Name", "Change Color", 100, 1, "Change CODE", true, "2020-10-01 00:00:00")
	assert.Nil(t, err)
	assert.Equal(t, "Change Name", prods.Nombre)

	prodGetOne, err := repository.GetOne(2)
	assert.Nil(t, err)
	assert.Equal(t, prodGetOne, prods)

}

func TestDeleteInDataBaseMock(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	prod, err := repository.Store("Producto Nuevo", "Rojo", 1000, 10, "UK-0000", false, "2021-12-12")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, prod.ID)

	err = repository.Delete(prod.ID)
	//verificando que elimine correctamente
	assert.Nil(t, err)

	//Verificando que fue borrado
	_, err = repository.GetOne(prod.ID)
	assert.Error(t, err)

}

func TestDeleteInDataBaseMockNotFound(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)

	err := repository.Delete(-1)
	//verificando que tenga error al momento de eliminar uno que no existe
	assert.Error(t, err)

}

func TestUpdateInDataBaseMockNotFound(t *testing.T) {
	db, _ := util.InitDb()
	repository := NewRepositorySQL(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repository.Update(ctx, -1, "Change Name", "Change Color", 100, 1, "Change CODE", true, "2020/10/01")
	assert.Error(t, err)

}

// -------------Test con go-sqlmock

func TestStoreWithSQLMock(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	colmuns := []string{
		"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fechaCreacion",
	}
	rows := sqlmock.NewRows(colmuns)
	rows.AddRow(newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products").WithArgs(newProd.ID).WillReturnRows(rows)
	repository := NewRepositorySQL(db)
	prodCreated, err := repository.Store(newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	assert.Nil(t, err)
	assert.Equal(t, newProd, prodCreated)

	prodGetOne, err := repository.GetOne(newProd.ID)
	assert.Nil(t, err)
	assert.Equal(t, prodCreated, prodGetOne)

}

func TestStoreErrorWithSQLMock(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnError(errors.New("Error Insertando"))

	repository := NewRepositorySQL(db)
	_, err = repository.Store(newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	assert.Error(t, err)

}

func TestGetOneWithSQLMock(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	colmuns := []string{
		"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fechaCreacion",
	}
	rows := sqlmock.NewRows(colmuns)
	rows.AddRow(newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products").WithArgs(newProd.ID).WillReturnRows(rows)
	repository := NewRepositorySQL(db)

	prodGetOne, err := repository.GetOne(newProd.ID)
	assert.Nil(t, err)
	assert.Equal(t, newProd, prodGetOne)

}

func TestGetOneWithSQLMockNotFound(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	colmuns := []string{
		"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fechaCreacion",
	}
	rows := sqlmock.NewRows(colmuns)
	rows.AddRow(newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products").WithArgs(2).WillReturnRows(rows)
	repository := NewRepositorySQL(db)

	prodGetOne, err := repository.GetOne(newProd.ID)
	assert.Error(t, err)
	assert.Equal(t, models.Product{}, prodGetOne)

}

func TestUpdateWithSQLMock(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("UPDATE products")
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))

	colmuns := []string{
		"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fechaCreacion",
	}
	rows := sqlmock.NewRows(colmuns)
	rows.AddRow(newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fechaCreacion from products").WithArgs(newProd.ID).WillReturnRows(rows)
	repository := NewRepositorySQL(db)

	prodUpdated, err := repository.Update(context.Background(), newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	assert.Nil(t, err)
	assert.Equal(t, newProd, prodUpdated)

	prodGetOne, err := repository.GetOne(newProd.ID)
	assert.Nil(t, err)
	assert.Equal(t, prodUpdated, prodGetOne)

}

func TestUpdateWithSQLMockNotFound(t *testing.T) {
	newProd := models.Product{ID: 1, Nombre: "Producto Nuevo", Color: "Rojo", Precio: 1000, Stock: 10, Codigo: "UK-0000", Publicado: false, FechaCreacion: "2021-12-12 00:00:00"}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("UPDATE products")
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 0))

	repository := NewRepositorySQL(db)

	prodUpdated, err := repository.Update(context.Background(), newProd.ID, newProd.Nombre, newProd.Color, newProd.Precio, newProd.Stock, newProd.Codigo, newProd.Publicado, newProd.FechaCreacion)
	assert.Error(t, err)
	assert.Equal(t, models.Product{}, prodUpdated)

}

func TestDeleteWithSQLMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("DELETE FROM products")
	mock.ExpectExec("DELETE FROM products").WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewRepositorySQL(db)

	err = repository.Delete(1)
	assert.Nil(t, err)

	_, err = repository.GetOne(1)
	assert.Error(t, err)

}

func TestDeleteWithSQLMockNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("DELETE FROM products")
	mock.ExpectExec("DELETE FROM products").WillReturnResult(sqlmock.NewResult(1, 0))

	repository := NewRepositorySQL(db)

	err = repository.Delete(1)
	assert.Error(t, err)
}
