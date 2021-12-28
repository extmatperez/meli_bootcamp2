package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
}

func TestStoresServiceOK(t *testing.T) {
	newProduct := models.Product{
		Name:  "Apples",
		Price: 120.50,
		Size:  5,
	}

	repo := NewRepository()
	service := NewService(repo)

	productCreated, _ := service.Store(newProduct.Name, newProduct.Price, newProduct.Size)

	assert.Equal(t, newProduct.Name, productCreated.Name)
	assert.Equal(t, newProduct.Price, productCreated.Price)
}
