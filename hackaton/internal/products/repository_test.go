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

type MockStore struct {
	CalledMethod bool
}

type MockArray struct {
	CalledMethod bool
}

var product string = `{
	"id": 1,
	"description": "Play Station 5 860GB",
	"Price": "689.23"
}`

var products string = `[
	{
		"id": 1,
		"description": "Mountain Dew",
		"Price": "665.85"
	},
	{
		"id": 2,
		"description": "XBOX Series X",
		"Price": "789.23"
	},
	{
		"id": 3,
		"description": "Nintendo Switch Neon",
		"Price": "489.23"
	}
]`

var productLines []string = []string{"1#$%#Mountain Dew#$%#665.85", "2#$%#Dried Apple#$%#663.3", "3#$%#Huck Towels White#$%#510.63"}

func (m *MockStore) Read(data interface{}) error {
	m.CalledMethod = true
	return json.Unmarshal([]byte(products), &data)
}

func (m *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	products = string(byteData)
	return nil
}

func (a *MockArray) ReadLines(path string) ([]string, error) {
	a.CalledMethod = true
	return productLines, nil
}

func NewMockStore() store.Store {
	return &MockStore{false}
}

func NewMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestRepoImportAllProductsError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)
	err = repo.ImportAllProducts()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestRepoImportAllProductsOk(t *testing.T) {
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

func TestRepoStoreProductOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newProduct := models.Product{
		Description: "Pepsi 2L",
		Price:       654.20,
	}

	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)

	createdProduct, err := repo.StoreProduct(newProduct)

	assert.Equal(t, newProduct.Description, createdProduct.Description)
	assert.Nil(t, err)
}

func TestRepoStoreProductError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newProduct := models.Product{
		Id:          1,
		Description: "Pepsi 2L",
		Price:       654.20,
	}

	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)

	_, err0 := repo.StoreProduct(newProduct)

	assert.Error(t, err0)
}

func TestRepoUpdateProductOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)

	updatedProduct := models.Product{
		Id:          101,
		Description: "Pepsi 2L",
		Price:       754.20,
	}

	productUpdated, err := repo.UpdateProduct(updatedProduct)

	assert.Nil(t, err)
	assert.NotNil(t, productUpdated)
	assert.Equal(t, updatedProduct.Id, productUpdated.Id)
	assert.Equal(t, updatedProduct.Description, productUpdated.Description)
	assert.Equal(t, updatedProduct.Price, updatedProduct.Price)
}

func TestRepoUpdateProductError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewProductRepository(&mock_array, db)

	updatedProduct := models.Product{}

	_, err = repo.UpdateProduct(updatedProduct)

	assert.NotNil(t, err)
}
