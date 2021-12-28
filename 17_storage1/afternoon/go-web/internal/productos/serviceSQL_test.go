package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/afternoon/go-web/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreOk(t *testing.T) {
	//arrange
	productoNuevo := models.Producto{
		Nombre:        "Camisa",
		Color:         "Roja",
		Precio:        10.0,
		Stock:         10,
		Codigo:        "CAM-008",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	//act
	producto, err := service.Store(productoNuevo)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productoNuevo.Nombre, producto.Nombre)
}

func TestGetByNameOk(t *testing.T) {
	//arrange
	nombre := "Camisa2"

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	//act
	producto, err := service.GetByName(nombre)

	//assert
	assert.Nil(t, err)
	assert.True(t, len(producto) > 0)
}
