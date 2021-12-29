package internal

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/proyecto/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/proyecto/pkg/store"
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

func TestStoreServiceSQL(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Chair",
		Color:         "Black",
		Precio:        "$560",
		Stock:         44,
		Codigo:        "kbsfdd9b-jbaf8-jkvaf898-jkgs8",
		Publicado:     true,
		FechaCreacion: "2010-01-01",
		Tipo:          models.Tipo{ID: 3},
	}

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	//Act
	productoCreado, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Stock, newProduct.Precio, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion, newProduct.Tipo.ID)

	//Assert
	assert.Equal(t, newProduct.Nombre, productoCreado.Nombre)
	assert.Nil(t, err)

	//Delete inserted
	service.Delete(productoCreado.ID)
}

func TestGetAllServiceSQL(t *testing.T) {
	//Arrenge
	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	//Act
	productosLeidos, err := service.GetAll()

	//Assert
	assert.Len(t, productosLeidos, 3)
	assert.Nil(t, err)
}

func TestGetServiceSQL(t *testing.T) {
	//Arrenge
	id := 21

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	//Act
	productoLeido, err := service.Get(id)

	//Assert
	assert.Equal(t, "Grove", productoLeido.Nombre)
	assert.Equal(t, "Crimson", productoLeido.Color)
	assert.Nil(t, err)
}

func TestGetByNameServiceSQL(t *testing.T) {
	//Arrenge
	nombre := "Grove"

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	//Act
	productosLeidos, err := service.GetByName(nombre)

	//Assert
	assert.Equal(t, "Crimson", productosLeidos[0].Color)
	assert.Equal(t, "Yellow", productosLeidos[1].Color)
	assert.Len(t, productosLeidos, 2)
	assert.Nil(t, err)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrenge
	productToUpdate := models.Producto{
		ID:            21,
		Nombre:        "Actualizado 4",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
	}

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	productoOriginal, _ := service.Get(productToUpdate.ID)

	//Act
	productoActualizado, err := service.Update(productToUpdate)

	//Assert
	assert.Equal(t, productToUpdate, productoActualizado)
	assert.Nil(t, err)

	//Regrear al producto original
	service.Update(productoOriginal)
}

func TestDeleteServiceSQL(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Grove",
		Color:         "Crimson",
		Precio:        "$3470.92",
		Stock:         654,
		Codigo:        "5c62ffa5-a28a-4c08-8edf-b213d4333bb0",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
		Tipo:          models.Tipo{ID: 1},
	}

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	productoCreado, _ := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Stock, newProduct.Precio, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion, newProduct.Tipo.ID)

	//Act
	err := service.Delete(productoCreado.ID)

	//Assert
	assert.Nil(t, err)
}

func TestGetAllFullDataServiceSQL(t *testing.T) {
	//Arrenge
	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	//Act
	productosLeidos, err := service.GetAllFullData()

	//Assert
	assert.Len(t, productosLeidos, 3)
	assert.Nil(t, err)
}

func TestGetWithContextServiceSQL(t *testing.T) {
	//Arrenge
	id := 21

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := db.StorageDB
	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoLeido, err := service.GetWithContext(ctx, id)

	//Assert
	assert.Equal(t, "Grove", productoLeido.Nombre)
	assert.Nil(t, err)
}

func TestUpdateWithContextServiceSQL(t *testing.T) {
	//Arrenge
	productToUpdate := models.Producto{
		ID:            21,
		Nombre:        "Actualizado 4",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
	}

	repository := NewRepositorySQL()
	service := NewServiceSQL(repository)

	productoOriginal, _ := service.Get(productToUpdate.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Act
	productoActualizado, err := service.UpdateWithContext(ctx, productToUpdate)

	//Assert
	assert.Equal(t, productToUpdate, productoActualizado)
	assert.Nil(t, err)

	//Regresar al producto original
	service.Update(productoOriginal)
}

//SQL Mock
func TestGetServiceSQLMock(t *testing.T) {
	//Arrenge
	id := 1

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"})
	rows.AddRow(1, "Grove", "Crismon", "$4673.98", 654, "5c62ffa5-a28a-4c08-8edf-b213d4333bb0", true, "1996-06-01")

	mock.ExpectQuery(regexp.QuoteMeta("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;")).WithArgs(id).WillReturnRows(rows)

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoLeido, err := service.Get(id)

	//Assert
	assert.Equal(t, "Grove", productoLeido.Nombre)
	assert.Equal(t, "Crismon", productoLeido.Color)
	assert.Nil(t, err)
}

func TestStoreServiceSQLMock(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Chair",
		Color:         "Black",
		Precio:        "$560",
		Stock:         44,
		Codigo:        "kbsfdd9b-jbaf8-jkvaf898-jkgs8",
		Publicado:     true,
		FechaCreacion: "2010-01-01",
		Tipo:          models.Tipo{ID: 3},
	}

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	mock.ExpectPrepare("insert into")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(8, 1))

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoLeido, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Stock, newProduct.Precio, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion, newProduct.Tipo.ID)

	//Assert
	assert.Equal(t, 8, productoLeido.ID)
	assert.Equal(t, "Chair", productoLeido.Nombre)
	assert.Equal(t, "Black", productoLeido.Color)
	assert.Nil(t, err)
}

func TestUpdateServiceSQLMock(t *testing.T) {
	//Arrenge
	productToUpdate := models.Producto{
		ID:            21,
		Nombre:        "Actualizado 4",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
	}

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"})
	rows.AddRow(21, "Actualizado 4", "Negro", "$ 50", 2, "123456R", true, "1996-06-01")

	prep := mock.ExpectPrepare(regexp.QuoteMeta("update productos set nombre = ?, color = ?, stock = ?, precio = ?, codigo = ?, publicado = ?, fecha_creacion = ? where id = ?;"))
	prep.ExpectExec().WithArgs(productToUpdate.Nombre, productToUpdate.Color, productToUpdate.Stock, productToUpdate.Precio, productToUpdate.Codigo, productToUpdate.Publicado, productToUpdate.FechaCreacion, productToUpdate.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoActualizado, err := service.Update(productToUpdate)

	//Assert
	assert.Equal(t, productToUpdate, productoActualizado)
	assert.Nil(t, err)
}

func TestDeleteServiceSQLMock(t *testing.T) {
	//Arrenge
	id := 1

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("delete from productos where id = ?;"))
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	err = service.Delete(id)

	//Assert
	assert.Nil(t, err)
}

func TestStoreServiceSQLMockFailed(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Chair",
		Color:         "Black",
		Precio:        "$560",
		Stock:         44,
		Codigo:        "kbsfdd9b-jbaf8-jkvaf898-jkgs8",
		Publicado:     true,
		FechaCreacion: "2010-01-01",
		Tipo:          models.Tipo{ID: 3},
	}

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	mock.ExpectPrepare("insert into")
	mock.ExpectExec("").WillReturnError(errors.New("no se pudo cargar la persona"))

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	_, err = service.Store(newProduct.Nombre, newProduct.Color, newProduct.Stock, newProduct.Precio, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion, newProduct.Tipo.ID)

	//Assert
	assert.Error(t, err)
}

func TestGetServiceSQLMockFailed(t *testing.T) {
	//Arrenge
	id := 1

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from productos where id = ?;")).WithArgs(id).WillReturnError(errors.New("producto no encontrado"))

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	_, err = service.Get(id)

	//Assert
	assert.Error(t, err)
}

//txdb
func TestStoreServiceSQLTxdb(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Chair",
		Color:         "Black",
		Precio:        "$560",
		Stock:         44,
		Codigo:        "kbsfdd9b-jbaf8-jkvaf898-jkgs8",
		Publicado:     true,
		FechaCreacion: "2010-01-01",
		Tipo:          models.Tipo{ID: 3},
	}

	db, err := db.InitDb()
	assert.Nil(t, err)

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoCreado, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Stock, newProduct.Precio, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion, newProduct.Tipo.ID)

	//Assert
	assert.Equal(t, newProduct.Nombre, productoCreado.Nombre)
	assert.Nil(t, err)
}

func TestGetServiceSQLTxdb(t *testing.T) {
	//Arrenge
	id := 21

	db, err := db.InitDb()
	assert.Nil(t, err)

	repository := NewRepositorySQLMock(db)
	service := NewServiceSQL(repository)

	//Act
	productoLeido, err := service.Get(id)

	//Assert
	assert.Equal(t, "Grove", productoLeido.Nombre)
	assert.Equal(t, "Crimson", productoLeido.Color)
	assert.Nil(t, err)
}
