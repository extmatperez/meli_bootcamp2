package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/GoStoragePrac/proyecto/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_store_ok(t *testing.T) {

	productoNuevo := models.Producto{
		Nombre:        "Casa",
		Color:         "Verde",
		Precio:        200000,
		Stock:         2,
		Codigo:        "AJJ035",
		Publicado:     true,
		FechaCreacion: "2021-11-15",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	productoDevuelto, err := service.Store(productoNuevo.Nombre, productoNuevo.Color, productoNuevo.Precio, productoNuevo.Stock, productoNuevo.Codigo, productoNuevo.Publicado, productoNuevo.FechaCreacion)

	assert.Nil(t, err)
	assert.Equal(t, productoNuevo.Nombre, productoDevuelto.Nombre)
}

func Test_getone_ok(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	productoBuscado := service.GetOne(1)

	assert.NotEqual(t, models.Producto{}, productoBuscado)

}

func Test_getall_ok(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	misProductosDB, err := service.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(misProductosDB) >= 0)

}

func Test_update_ok(t *testing.T) {

	productoActualizado := models.Producto{
		Id:     1,
		Nombre: "Avion",
		Color:  "Violeta",
		Precio: 100000000,
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	productoAnterior := service.GetOne(productoActualizado.Id)

	productoCargado, _ := service.Update(productoActualizado)

	assert.Equal(t, productoActualizado.Nombre, productoCargado.Nombre)
	assert.Equal(t, productoActualizado.Color, productoCargado.Color)

	_, err := service.Update(productoAnterior)
	assert.Nil(t, err)
}

func TestDeleteServiceSQL(t *testing.T) {

	productoNuevo := models.Producto{
		Nombre:        "Barco",
		Color:         "Azul",
		Precio:        400000000,
		Stock:         1,
		Codigo:        "AQJ343",
		Publicado:     false,
		FechaCreacion: "2022-11-15",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)
	productoCreado, _ := service.Store(productoNuevo.Nombre, productoNuevo.Color, productoNuevo.Precio, productoNuevo.Stock, productoNuevo.Codigo, productoNuevo.Publicado, productoNuevo.FechaCreacion)

	err := service.Delete(productoCreado.Id)

	assert.Nil(t, err)

}

// func Test_getbyname_ok(t *testing.T) {

// 	repo := NewRepositorySQL()
// 	service := NewServiceSQL(repo)

// 	productoBuscado := service.GetByName("Auto")

// 	assert.Equal(t, models.Producto{}, productoBuscado)
// }
