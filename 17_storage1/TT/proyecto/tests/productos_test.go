package tests

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/internal/models"
	productos "github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/internal/productos"
	"github.com/stretchr/testify/assert"
)

func TestStoreServiceSQL(t *testing.T) {
	//Arrenge
	newProduct := models.Producto{
		Nombre:        "Grove",
		Color:         "Crimson",
		Precio:        "$3470.92",
		Stock:         654,
		Codigo:        "5c62ffa5-a28a-4c08-8edf-b213d4333bb0",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
	}

	repository := productos.NewRepositorySQL()

	//Act
	productoCreado, err := repository.Store(newProduct)

	//Assert
	assert.Equal(t, newProduct.Nombre, productoCreado.Nombre)
	assert.Nil(t, err)
}

func TestGetServiceSQL(t *testing.T) {
	//Arrenge
	id := 3

	repository := productos.NewRepositorySQL()

	//Act
	productoLeido, err := repository.Get(id)

	//Assert
	assert.Equal(t, "Chocolate", productoLeido.Nombre)
	assert.Equal(t, "Negro", productoLeido.Color)
	assert.Nil(t, err)
}

func TestGetByNameServiceSQL(t *testing.T) {
	//Arrenge
	nombre := "Grove"

	repository := productos.NewRepositorySQL()

	//Act
	productosLeidos, err := repository.GetByName(nombre)

	//Assert
	assert.Equal(t, "Crimson", productosLeidos[0].Color)
	assert.Equal(t, "Yellow", productosLeidos[1].Color)
	assert.Len(t, productosLeidos, 2)
	assert.Nil(t, err)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrenge
	productToUpdate := models.Producto{
		ID:            1,
		Nombre:        "Actualizado 4",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
	}

	repository := productos.NewRepositorySQL()

	//Act
	productoActualizado, err := repository.Update(productToUpdate)

	//Assert
	assert.Equal(t, productToUpdate, productoActualizado)
	assert.Nil(t, err)
}
