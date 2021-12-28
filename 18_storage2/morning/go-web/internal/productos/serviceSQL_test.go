package internal

import (
	"context"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/morning/go-web/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreOk(t *testing.T) {
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

func TestGetAllOk(t *testing.T) {
	//arrange

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	//act
	producto, err := service.GetAll()

	//assert
	assert.Nil(t, err)
	assert.True(t, len(producto) >= 0)
}

func TestUpdateOk(t *testing.T) {
	//arrange
	productoNuevo := models.Producto{
		Nombre:        "Gorra",
		Color:         "Amarilla",
		Precio:        12.0,
		Stock:         10,
		Codigo:        "CAM-012",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//act
	producto, err := service.Update(ctx, productoNuevo, 10)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productoNuevo.Nombre, producto.Nombre)
}
