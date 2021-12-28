package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	// Arrange
	repository := NewRepository()
	service := NewService(repository)
	expectedResult := []domain.Product{
		{Id: 1, Name: "Mate", Price: 200.5, Description: "Para tomar mate"},
		{Id: 4, Name: "Mateoli", Price: 200.5, Description: "Para tomar mate"},
	}

	// Act
	result, err := service.GetByName(context.Background(), "mate")

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err)
}

func TestStore(t *testing.T) {
	// Arrange
	repository := NewRepository()
	service := NewService(repository)
	expectedResult := domain.Product{Name: "Pelota", Price: 555.5, Description: "Para tomar jugar al futbol"}

	// Act
	result, err := service.Store(context.Background(), expectedResult.Name, expectedResult.Price, expectedResult.Description)

	// Assert
	assert.Equal(t, expectedResult.Name, result.Name, "result name should be equal to expected result name")
	assert.Equal(t, expectedResult.Price, result.Price, "result price should be equal to expected result price")
	assert.Equal(t, expectedResult.Description, result.Description, "result description should be equal to expected result description")
	assert.Nil(t, err)
}
