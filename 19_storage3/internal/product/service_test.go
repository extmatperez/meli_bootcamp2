package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)

	// Act
	result, err := service.GetAll(context.Background())

	// Assert
	assert.NotNil(t, result, "result should not be nil")
	assert.True(t, len(result) > 0, "result should has more than one result")
	assert.Nil(t, err, "error should be nil")
}
func TestGet(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)
	expectedResult := domain.Product{Id: 1, Name: "Mate", Price: 200.5, Description: "Para tomar mate"}

	// Act
	result, err := service.Get(context.Background(), expectedResult.Id)

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")
}

func TestGetByName(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)
	expectedResult := []domain.Product{
		{Id: 1, Name: "Mate", Price: 200.5, Description: "Para tomar mate"},
		{Id: 4, Name: "Mateoli", Price: 200.5, Description: "Para tomar mate"},
	}

	// Act
	result, err := service.GetByName(context.Background(), "mate")

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")
}

func TestStore(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)
	expectedResult := domain.Product{Name: "Pelota", Price: 555.5, Description: "Para tomar jugar al futbol"}

	// Act
	result, err := service.Store(context.Background(), expectedResult.Name, expectedResult.Price, expectedResult.Description)

	// Assert
	assert.Equal(t, expectedResult.Name, result.Name, "result name should be equal to expected result name")
	assert.Equal(t, expectedResult.Price, result.Price, "result price should be equal to expected result price")
	assert.Equal(t, expectedResult.Description, result.Description, "result description should be equal to expected result description")
	assert.Nil(t, err, "error should be nil")

	err = service.Delete(context.Background(), result.Id)
	assert.Nil(t, err, "error should be nil")
}

func TestUpdate(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)

	expectedResult := domain.Product{Id: 1, Name: "Matecito", Price: 999, Description: "El mejor matecito"}
	oldData, _ := service.Get(context.Background(), 1)

	// Act
	result, err := service.Update(context.Background(), expectedResult.Id, expectedResult.Name, expectedResult.Price, expectedResult.Description)

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")

	_, err = service.Update(context.Background(), oldData.Id, oldData.Name, oldData.Price, oldData.Description)
	assert.Nil(t, err, "error should be nil")
}

func TestDelete(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)
	service := NewService(repository)
	productToDelete := domain.Product{Name: "Pelotita a eliminar", Price: 1.1, Description: "se va a eliminar"}

	// Act
	result, _ := service.Store(context.Background(), productToDelete.Name, productToDelete.Price, productToDelete.Description)
	err := service.Delete(context.Background(), result.Id)

	// Assert
	assert.Nil(t, err, "error should be nil")
}
