package product

import (
	"context"
	"testing"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
)

//testing para get de productos
func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misProducts, _ := service.GetAll(context.Background())

	assert.Equal(t, 2, len(misProducts))
	assert.True(t, stubRepo.useGetAll)
}

//Testing para get de productos
func TestGetAllServiceExistent(t *testing.T) {
	//Arrange
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	prodGet, err := service.Get(context.Background(), 1)

	//assert
	assert.Equal(t, "a41", prodGet.ProductCode)
	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

//Testing para get no exiaste producto
func TestGetAllServiceNotExistent(t *testing.T) {
	//Arrange
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	_, err := service.Get(context.Background(), 3)

	//assert
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

//Testing para creado con exito
func TestCreateOk(t *testing.T) {
	//Arrange
	newProd := domain.Product{
		ID:             3,
		Description:    "articulo test",
		ExpirationRate: 3,
		FreezingRate:   5,
		Height:         5,
		Length:         66.2,
		Netweight:      66.3,
		ProductCode:    "a43",
		RecomFreezTemp: 13.3,
		Width:          2.5,
		ProductTypeID:  11,
		SellerID:       0,
	}
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	Idprod, err := service.Save(context.Background(), newProd)

	//Assert
	assert.Equal(t, 3, Idprod)
	assert.Nil(t, err)

}

//Testing para creado con conflicto
//por product code repetido
func TestCreateConflict(t *testing.T) {
	//Arrange
	newProd := domain.Product{
		ID:             1,
		Description:    "articulo test",
		ExpirationRate: 3,
		FreezingRate:   5,
		Height:         5,
		Length:         66.2,
		Netweight:      66.3,
		ProductCode:    "a41",
		RecomFreezTemp: 13.3,
		Width:          2.5,
		ProductTypeID:  11,
		SellerID:       0,
	}
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	Idprod, err := service.Save(context.Background(), newProd)

	//Assert
	assert.Equal(t, 0, Idprod)
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)

}

//Testing para actualizado con exito
func TestUpdateExistent(t *testing.T) {
	//Arrange
	newProd := domain.Product{
		ID:             1,
		Description:    "articulo updated",
		ExpirationRate: 3,
		FreezingRate:   5,
		Height:         5,
		Length:         66.2,
		Netweight:      66.3,
		ProductCode:    "a41",
		RecomFreezTemp: 13.3,
		Width:          2.5,
		ProductTypeID:  11,
		SellerID:       0,
	}
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	prodUpdated, err := service.Update(context.Background(), newProd)

	//Assert
	assert.Equal(t, newProd.Description, prodUpdated.Description)
	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)

}

//Testing para error con el id
func TestUpdateNonExistent(t *testing.T) {
	//Arrange
	newProd := domain.Product{
		ID:             5,
		Description:    "articulo updated",
		ExpirationRate: 3,
		FreezingRate:   5,
		Height:         5,
		Length:         66.2,
		Netweight:      66.3,
		ProductCode:    "a41",
		RecomFreezTemp: 13.3,
		Width:          2.5,
		ProductTypeID:  11,
		SellerID:       0,
	}
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	_, err := service.Update(context.Background(), newProd)

	//Assert
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

//Testing para borrar con exito
func TestDeleteExitent(t *testing.T) {
	//Arange
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 2)

	//Assert
	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func TestDeleteNonExitent(t *testing.T) {
	//Arange
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 5)

	//Assert
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

//Testing para funcion exixt
func TestExists(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	boolean := service.Exists(context.Background(), "a41")
	assert.True(t, boolean)
}
