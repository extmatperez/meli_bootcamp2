package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryMongoGetAll(t *testing.T) {
	// Arrange
	mongoRepository := NewMongoRepository(database.MongoDB)
	_ = mongoRepository.SetDatabaseAndCollection("bootcamp_storage", "products")

	// Act
	result, err := mongoRepository.GetAll(context.Background())
	assert.Nil(t, err, "error should be nil")

	// Assert
	assert.True(t, len(result) > 0, "len of products should be more than zero")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryMongoStore(t *testing.T) {
	// Arrange
	mongoRepository := NewMongoRepository(database.MongoDB)
	_ = mongoRepository.SetDatabaseAndCollection("bootcamp_storage", "products")
	expectedResult := domain.ProductMongo{
		Name:        "to-delete",
		Price:       1.5,
		Description: "para eliminar",
	}

	// Act
	result, err := mongoRepository.Store(context.Background(), expectedResult)
	assert.Nil(t, err, "error should be nil")

	// Assert
	assert.NotNil(t, result.Id, "id should not be nil")
	assert.Equal(t, expectedResult.Name, result.Name, "result name should be equal to expected result name")
	assert.Equal(t, expectedResult.Price, result.Price, "result price should be equal to expected result price")
	assert.Equal(t, expectedResult.Description, result.Description, "result description should be equal to expected result description")
	assert.Nil(t, err, "error should be nil")
}
