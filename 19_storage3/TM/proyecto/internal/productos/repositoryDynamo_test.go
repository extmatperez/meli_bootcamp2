package internal

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreDynamo(t *testing.T) {
	//Assert
	newProduct := models.ProductoDynamo{
		ID:            "1",
		Nombre:        "Grove",
		Color:         "Crimson",
		Precio:        "$3470.92",
		Stock:         654,
		Codigo:        "5c62ffa5-a28a-4c08-8edf-b213d4333bb0",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
		Tipo:          models.Tipo{ID: 1, Descripcion: "Mueble"},
	}

	db, err := InitDynamo()
	assert.Nil(t, err)

	repository := NewDynamoRepository(db, "Productos")

	//Act
	err = repository.Store(context.Background(), &newProduct)

	//Assert
	assert.Nil(t, err)
}

func TestGetOneDynamo(t *testing.T) {
	//Assert
	id := "1"

	db, err := InitDynamo()
	assert.Nil(t, err)

	repository := NewDynamoRepository(db, "Productos")

	//Act
	productoLeido, err := repository.GetOne(context.Background(), id)

	//Assert
	productoEsperado := models.ProductoDynamo{
		ID:            "1",
		Nombre:        "Grove",
		Color:         "Crimson",
		Precio:        "$3470.92",
		Stock:         654,
		Codigo:        "5c62ffa5-a28a-4c08-8edf-b213d4333bb0",
		Publicado:     true,
		FechaCreacion: "1996-06-01",
		Tipo:          models.Tipo{ID: 1, Descripcion: "Mueble"},
	}

	assert.Equal(t, productoEsperado, *productoLeido)
	assert.Nil(t, err)
}

func TestDeleteDynamo(t *testing.T) {
	//Assert
	id := "1"

	db, err := InitDynamo()
	assert.Nil(t, err)

	repository := NewDynamoRepository(db, "Productos")

	//Act
	err = repository.Delete(context.Background(), id)

	//Assert
	assert.Nil(t, err)
}
