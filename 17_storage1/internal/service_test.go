package internal

import (
	"context"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/17_storage1/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	productoNuevo := models.Producto{
		Name:        "producto 1",
		Type:        "Varios",
		Count:       27,
		Price:       89.92,
		WarehouseID: 1,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	productoCreado, err := service.Store(productoNuevo.Name, productoNuevo.Type, productoNuevo.Count, productoNuevo.Price)

	assert.Nil(t, err)
	assert.Equal(t, productoNuevo.Name, productoCreado.Name)
	assert.Equal(t, productoNuevo.Type, productoCreado.Type)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	productoNuevo := models.Producto{
		Name:        "producto 1",
		Type:        "Varios",
		Count:       27,
		Price:       89.92,
		WarehouseID: 1,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	productoCargado := service.GetOne(1)

	assert.Equal(t, productoNuevo.Name, productoCargado.Name)
	assert.Equal(t, productoNuevo.Type, productoCargado.Type)
	// assert.Nil(t, misPersonas)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	productoUpdate := models.Producto{
		ID:          1,
		Name:        "producto 3",
		Type:        "Otros",
		Count:       22,
		Price:       21.92,
		WarehouseID: 1,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	productoAnterior := service.GetOne(productoUpdate.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productoCargado, _ := service.Update(ctx, productoUpdate)

	assert.Equal(t, productoUpdate.Name, productoCargado.Name)
	assert.Equal(t, productoUpdate.Type, productoCargado.Type)
	// assert.Nil(t, misPersonas)
	_, err := service.Update(ctx, productoAnterior)
	assert.Nil(t, err)
}

func TestUpdateServiceSQL_Failed(t *testing.T) {
	//Arrange
	productoUpdate := models.Producto{
		ID:          15,
		Name:        "producto 3",
		Type:        "Otros",
		Count:       22,
		Price:       21.92,
		WarehouseID: 1,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := service.Update(ctx, productoUpdate)

	assert.Equal(t, "no se encontro la persona", err.Error())
	// assert.Nil(t, misPersonas)
}

func TestGetAllServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misPersonasDB, err := service.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(misPersonasDB) >= 0)
	// assert.Nil(t, misPersonas)
}

func TestDeleteServiceSQL(t *testing.T) {
	//Arrange

	productoNuevo := models.Producto{
		Name:        "producto 1",
		Type:        "Varios",
		Count:       27,
		Price:       89.92,
		WarehouseID: 1,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCreada, _ := service.Store(productoNuevo.Name, productoNuevo.Type, productoNuevo.Count, productoNuevo.Price)

	err := service.Delete(personaCreada.ID)

	assert.Nil(t, err)
	// assert.Nil(t, misPersonas)
}

func TestDeleteServiceSQL_NotFound(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	err := service.Delete(0)

	assert.Equal(t, "no se encontro la persona", err.Error())
	// assert.Nil(t, misPersonas)
}

func TestGetFullDataServiceSQL(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misProductos, err := service.GetFullData()

	assert.Nil(t, err)
	assert.True(t, len(misProductos) >= 0)
	assert.Equal(t, "Main Warehouse", misProductos[0].Warehouse.Name)
	// fmt.Printf("\n%+v", misPersonas)
	// assert.Nil(t, misPersonas)
}
