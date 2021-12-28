package product

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	// Arrange
	repository := NewRepository(database.StorageDB)

	// Act
	products, err := repository.GetAll(context.Background())

	// Assert
	assert.True(t, len(products) > 0, "len of products should be more than zero")
	assert.Nil(t, err, "eeror should be nil")
}
