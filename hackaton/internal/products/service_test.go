package internal

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
	"github.com/stretchr/testify/assert"
)

type ServiceMockArray struct {
	CalledMethod bool
}

func (a *ServiceMockArray) ReadLines(path string) ([]string, error) {
	a.CalledMethod = true
	return productLines, nil
}

func NewServiceMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestServiceImportAllProductsError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)
	service := NewProductService(repo)
	err = service.ImportAllProducts()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestServiceImportAllProductsOk(t *testing.T) {
	dataBytes := []byte(products)
	var expectedProducts []models.Product
	json.Unmarshal(dataBytes, &expectedProducts)

	var p models.Product
	db := db.StorageDB
	rows, err := db.Query("SELECT id, `description`, price FROM Product WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Description, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
	}

	assert.Equal(t, expectedProducts[0].Id, p.Id)
	assert.Equal(t, expectedProducts[0].Description, p.Description)
}

func TestServiceStoreProductOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newProduct := models.Product{
		Description: "Pepsi 2L",
		Price:       654.20,
	}

	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)
	service := NewProductService(repo)

	createdProduct, err := service.StoreProduct(newProduct.Description, newProduct.Price)

	assert.Equal(t, newProduct.Description, createdProduct.Description)
	assert.Nil(t, err)
}

func TestServiceUpdateProductOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)
	service := NewProductService(repo)

	updatedProduct := models.Product{
		Id:          101,
		Description: "Pepsi 3L",
		Price:       754.20,
	}

	productUpdated, err := service.UpdateProduct(updatedProduct)

	assert.Nil(t, err)
	assert.NotNil(t, productUpdated)
	assert.Equal(t, updatedProduct.Id, productUpdated.Id)
	assert.Equal(t, updatedProduct.Description, productUpdated.Description)
	assert.Equal(t, updatedProduct.Price, updatedProduct.Price)
}

func TestServiceUpdateProductError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)
	service := NewProductService(repo)

	updatedProduct := models.Product{}

	_, err = service.UpdateProduct(updatedProduct)

	assert.NotNil(t, err)
}
