package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/afternoon/txdbTests/go-web/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/afternoon/txdbTests/go-web/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestStoreOk_txdb(t *testing.T) {

	productoNuevo := models.Producto{
		Nombre:        "Camisa3",
		Color:         "Roja",
		Precio:        10.0,
		Stock:         10,
		Codigo:        "CAM-021",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQLMock(repo)

	productoDevuelto, err := service.Store(productoNuevo)

	assert.Nil(t, err)
	assert.Equal(t, productoNuevo.Color, productoDevuelto.Color)
	assert.Equal(t, productoNuevo.Nombre, productoDevuelto.Nombre)
}

func TestGetOneOk_txdb(t *testing.T) {

	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQLMock(repo)

	productoDevuelto, err := service.GetOne(6)

	assert.Nil(t, err)
	assert.Equal(t, 6, productoDevuelto.ID)
	assert.Equal(t, "Camisa2", productoDevuelto.Nombre)
}

func TestUpdateOk_txdb(t *testing.T) {

	productoNuevo := models.Producto{
		Nombre:        "Camisa3",
		Color:         "Blanca",
		Precio:        10.0,
		Stock:         10,
		Codigo:        "CAM-021",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQLMock(repo)

	productoDevuelto, err := service.Update(productoNuevo, 5)

	assert.Nil(t, err)
	assert.Equal(t, "Blanca", productoDevuelto.Color)
}

func TestDeleteOk_txdb(t *testing.T) {

	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQLMock(repo)

	err = service.Delete(5)

	assert.Nil(t, err)
}
