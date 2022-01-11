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

var sale string = `{
	"id": 1,
	"idProduct": 101,
	"idInvoice": 51,
	"quantity": 16
}`

var saless string = `[
	{
		"id": 1,
		"idProduct": 2,
		"idInvoice": 18,
		"quantity": 53
	},
	{
		"id": 2,
		"idProduct": 58,
		"idInvoice": 12,
		"quantity": 45
	},
	{
		"id": 3,
		"idProduct": 13,
		"idInvoice": 35,
		"quantity": 11
	}
]`

var saleLines []string = []string{"1#$%#2#$%#18#$%#52.9", "2#$%#53#$%#72#$%#44618", "3#$%#18#$%#74#$%#93.4"}

func (m *MockStore) Read(data interface{}) error {
	m.CalledMethod = true
	return json.Unmarshal([]byte(saless), &data)
}

func (m *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	saless = string(byteData)
	return nil
}

func (a *MockArray) ReadLines(path string) ([]string, error) {
	a.CalledMethod = true
	return saleLines, nil
}

func NewMockStore() store.Store {
	return &MockStore{false}
}

func NewMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestRepoImportAllSalesError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)
	err = repo.ImportAllSales()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestRepoImportAllSalesOk(t *testing.T) {
	dataBytes := []byte(saless)
	var expectedSales []models.Sale
	json.Unmarshal(dataBytes, &expectedSales)

	var s models.Sale
	db := db.StorageDB
	rows, err := db.Query("SELECT id, idProduct, idInvoice, quantity FROM Sale WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&s.Id, &s.IdProduct, &s.IdInvoice, &s.Quantity)
		if err != nil {
			log.Fatal(err)
		}
	}

	assert.Equal(t, expectedSales[0].Id, s.Id)
	assert.Equal(t, expectedSales[0].IdProduct, s.IdProduct)
	assert.Equal(t, expectedSales[0].Quantity, s.Quantity)
}

func TestRepoStoreSaleOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newSale := models.Sale{
		IdProduct: 45,
		IdInvoice: 11,
		Quantity:  23,
	}

	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)

	createdSale, err := repo.StoreSale(newSale)

	assert.Equal(t, newSale.IdProduct, createdSale.IdProduct)
	assert.Equal(t, newSale.Quantity, createdSale.Quantity)
	assert.Nil(t, err)
}

func TestRepoStoreSaleError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newSale := models.Sale{
		Id:        1,
		IdProduct: 45,
		IdInvoice: 11,
		Quantity:  23,
	}

	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)

	_, err0 := repo.StoreSale(newSale)

	assert.Error(t, err0)
}

func TestRepoUpdateSaleOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)

	updatedSale := models.Sale{
		Id:        1001,
		IdProduct: 45,
		IdInvoice: 14,
		Quantity:  23,
	}

	saleUpdated, err := repo.UpdateSale(updatedSale)

	assert.Nil(t, err)
	assert.NotNil(t, saleUpdated)
	assert.Equal(t, updatedSale.Id, saleUpdated.Id)
	assert.Equal(t, updatedSale.IdProduct, saleUpdated.IdProduct)
}

func TestRepoUpdateSaleError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)

	updatedSale := models.Sale{}

	_, err = repo.UpdateSale(updatedSale)

	assert.NotNil(t, err)
}
