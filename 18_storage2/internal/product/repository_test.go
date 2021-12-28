package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)

	// Act
	products, err := repository.GetAll(context.Background())

	// Assert
	assert.True(t, len(products) > 0, "len of products should be more than zero")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryGet(t *testing.T) {
	// Arrange
	db, _ := database.InitTxSqlDb()
	repository := NewRepository(db)
	defer db.Close()
	expectedResult := domain.Product{
		Id:          1,
		Name:        "Mate",
		Price:       200.5,
		Description: "Para tomar mate",
	}

	// Act
	result, err := repository.Get(context.Background(), expectedResult.Id)

	// Assert
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryGetNotFound(t *testing.T) {
	// Arrange
	db, _ := database.InitTxSqlDb()
	repository := NewRepository(db)
	defer db.Close()
	expectedResult := domain.Product{}

	// Act
	result, err := repository.Get(context.Background(), 0)

	// Assert
	assert.Error(t, err, "error should be nil")
	assert.Equal(t, expectedResult, result, "result should be equal to expected result")
}

func TestRepositoryStore(t *testing.T) {
	// Arrange
	db, _ := database.InitTxSqlDb()
	repository := NewRepository(db)
	defer db.Close()
	expectedResult := domain.Product{
		Name:        "Mate",
		Price:       200.5,
		Description: "Para tomar mate",
	}

	// Act
	result, err := repository.Store(context.Background(), expectedResult)

	// Assert
	assert.Equal(t, expectedResult.Name, result.Name, "result name should be equal to expected result name")
	assert.Equal(t, expectedResult.Price, result.Price, "result price should be equal to expected result price")
	assert.Equal(t, expectedResult.Description, result.Description, "result description should be equal to expected result description")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryUpdate(t *testing.T) {
	// Arrange
	db, _ := database.InitTxSqlDb()
	repository := NewRepository(db)
	defer db.Close()
	expectedResult := domain.Product{
		Id:          1,
		Name:        "Matecito",
		Price:       201.5,
		Description: "Para tomar matecito",
	}

	// Act
	result, err := repository.Update(context.Background(), expectedResult)

	// Assert
	assert.Equal(t, expectedResult.Id, result.Id, "result id should be equal to expected result id")
	assert.Equal(t, expectedResult.Name, result.Name, "result name should be equal to expected result name")
	assert.Equal(t, expectedResult.Price, result.Price, "result price should be equal to expected result price")
	assert.Equal(t, expectedResult.Description, result.Description, "result description should be equal to expected result description")
	assert.Nil(t, err, "error should be nil")

	newExpectedResult, err := repository.Get(context.Background(), expectedResult.Id)

	assert.Equal(t, newExpectedResult.Id, expectedResult.Id, "result id should be equal to expected result id")
	assert.Equal(t, newExpectedResult.Name, expectedResult.Name, "result name should be equal to expected result name")
	assert.Equal(t, newExpectedResult.Price, expectedResult.Price, "result price should be equal to expected result price")
	assert.Equal(t, newExpectedResult.Description, expectedResult.Description, "result description should be equal to expected result description")
	assert.Nil(t, err, "error should be nil")
}

func TestRepositoryDelete(t *testing.T) {
	// Arrange
	db, _ := database.InitTxSqlDb()
	repository := NewRepository(db)
	defer db.Close()

	// Act
	err := repository.Delete(context.Background(), 1)

	// Assert
	assert.Nil(t, err, "error should be nil")

	deletedProduct := domain.Product{}
	resultFindDeletedProduct, err := repository.Get(context.Background(), 1)

	assert.Error(t, err, "error should not be nil")
	assert.Equal(t, deletedProduct, resultFindDeletedProduct, "result should be equal to expected result")

	products, _ := repository.GetAll(context.Background())

	deletedProductFound := false
	for i := 0; i < len(products); i++ {
		if products[i].Id == 1 {
			deletedProductFound = true
		}
	}

	assert.True(t, !deletedProductFound, "deleted product should not found")
}
