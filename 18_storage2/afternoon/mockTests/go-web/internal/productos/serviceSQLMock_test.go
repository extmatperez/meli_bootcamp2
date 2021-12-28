package internal

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/afternoon/mockTests/go-web/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreOkMock(t *testing.T) {

	//arrange
	productoNuevo := models.Producto{
		Nombre:        "Camisa3",
		Color:         "Roja",
		Precio:        10.0,
		Stock:         10,
		Codigo:        "CAM-021",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO productos")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(3, 1))

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQLMock(repo)

	//act
	producto, err := service.Store(productoNuevo)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, 3, producto.ID)
}

func TestGetOneOkMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fechaCreacion"})
	rows.AddRow(6, "Camisa1", "Roja", 10.0, 10, "CAM-001", true, "2020-01-01")
	mock.ExpectQuery(getOne).WithArgs(6).WillReturnRows(rows)

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQLMock(repo)

	producto, err := service.GetOne(6)

	assert.Nil(t, err)
	assert.Equal(t, 6, producto.ID)
}

func TestUpdateOkMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	producto := models.Producto{
		Nombre:        "Camisa1",
		Color:         "Roja",
		Precio:        10.0,
		Stock:         10,
		Codigo:        "CAM-001",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	mock.ExpectPrepare("UPDATE productos")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(3, 1))

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQLMock(repo)

	productoDevuelto, err := service.Update(producto, 3)

	assert.Nil(t, err)
	assert.Equal(t, 3, productoDevuelto.ID)
}

func TestDeleteOkMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM productos")
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepositorySQLMock(db)
	service := NewServiceSQLMock(repo)

	err = service.Delete(1888)

	assert.Nil(t, err)
}
