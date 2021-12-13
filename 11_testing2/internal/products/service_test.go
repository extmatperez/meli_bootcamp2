package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/pkg/store"
	"github.com/stretchr/testify/assert"
)

var productsServiceTest = `
[{
	"id": 1,
	"name": "Pelota",
	"color": "Negro",
	"price": 1505.5,
	"stock": 200,
	"code": "#0000000f1",
	"published": true,
	"created_at": "21/11/2021"
},
{
	"id": 2,
	"name": "Botines",
	"color": "Blanco",
	"price": 5020.5,
	"stock": 50,
	"code": "#0000000f2",
	"published": false,
	"created_at": "12/10/2021"
}]`

func TestUpdateMock(t *testing.T) {
	// Arrange
	productBytes := []byte(productsServiceTest)
	dbMock := store.Mock{Data: productBytes, ReadUsed: false}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	expectedResult := Product{1, "Pelota act", "Negra act", 1000.05, 100, "#CodeUpdated", false, "21/11/1997"}

	// Act
	result, _ := service.Update(1, "Pelota act", "Negra act", 1000.05, 100, "#CodeUpdated", false, "21/11/1997")

	// Assert
	assert.Equal(t, expectedResult, result, "should be equal")
	assert.True(t, dbMock.ReadUsed)
}

func TestUpdateMockIdNotExists(t *testing.T) {
	// Arrange
	productBytes := []byte(productsServiceTest)
	dbMock := store.Mock{Data: productBytes}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	// Act
	result, err := service.Update(3, "Pelota act", "Negra act", 1000.05, 100, "#CodeUpdated", false, "21/11/1997")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, Product{}, result, "should be equal")
}

func TestUpdateMockError(t *testing.T) {
	// Arrange
	dbMock := store.Mock{}
	storeStub := store.FileStore{Mock: &dbMock}
	repository := NewRepository(&storeStub)
	service := NewService(repository)

	// Act
	result, err := service.Update(3, "Pelota act", "Negra act", 1000.05, 100, "#CodeUpdated", false, "21/11/1997")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, Product{}, result, "should be equal")
}
