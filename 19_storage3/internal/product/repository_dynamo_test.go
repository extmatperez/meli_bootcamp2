package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryDynamoGetAll(t *testing.T) {
	// Arrange
	dynamoDB, err := database.InitDynamo()
	assert.Nil(t, err, "error should be nil")
	dynamoRepository := NewDynamoRepository(dynamoDB, "products")

	// Act
	result, err := dynamoRepository.GetAll(context.Background())

	// Assert
	assert.True(t, len(result) > 0, "len of products should be more than zero")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryDynamoGet(t *testing.T) {
	// Arrange
	dynamoDB, err := database.InitDynamo()
	assert.Nil(t, err, "error should be nil")
	dynamoRepository := NewDynamoRepository(dynamoDB, "products")
	expectedResult := &domain.ProductDynamo{
		Id:          "1",
		Name:        "Mate",
		Price:       200.5,
		Description: "Para tomar mate",
	}

	// Act
	result, err := dynamoRepository.Get(context.Background(), "1")

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryDynamoStore(t *testing.T) {
	// Arrange
	dynamoDB, err := database.InitDynamo()
	assert.Nil(t, err, "error should be nil")
	dynamoRepository := NewDynamoRepository(dynamoDB, "products")
	expectedResult := &domain.ProductDynamo{
		Id:          "to-delete",
		Name:        "para eliminar",
		Price:       1000.5,
		Description: "para eliminar descripcion",
	}

	// Act
	result, err := dynamoRepository.Store(context.Background(), expectedResult)

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryDynamoDelete(t *testing.T) {
	// Arrange
	dynamoDB, err := database.InitDynamo()
	assert.Nil(t, err, "error should be nil")
	dynamoRepository := NewDynamoRepository(dynamoDB, "products")

	// Act
	err = dynamoRepository.Delete(context.Background(), "to-delete")

	// Assert
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryDynamoUpdate(t *testing.T) {
	// Arrange
	dynamoDB, err := database.InitDynamo()
	assert.Nil(t, err, "error should be nil")
	dynamoRepository := NewDynamoRepository(dynamoDB, "products")
	expectedResult := &domain.ProductDynamo{
		Id:          "1",
		Name:        "Matecito",
		Price:       100.5,
		Description: "Para tomar matecito",
	}

	oldProduct, err := dynamoRepository.Get(context.Background(), expectedResult.Id)
	assert.Nil(t, err, "error should be nil")

	// Act
	result, err := dynamoRepository.Update(context.Background(), expectedResult)

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")

	// Put the old data again
	_, err = dynamoRepository.Update(context.Background(), oldProduct)
	assert.Nil(t, err, "error should be nil")
}
