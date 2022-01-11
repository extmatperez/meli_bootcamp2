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
	return saleLines, nil
}

func NewServiceMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestServiceImportAllSalesError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)
	service := NewSaleService(repo)
	err = service.ImportAllSales()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestServiceImportAllSalesOk(t *testing.T) {
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

func TestServiceStoreSaleOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newSale := models.Sale{
		IdProduct: 45,
		IdInvoice: 11,
		Quantity:  23,
	}

	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)
	service := NewSaleService(repo)

	createdSale, err := service.StoreSale(newSale.IdProduct, newSale.IdInvoice, newSale.Quantity)

	assert.Equal(t, newSale.IdProduct, createdSale.IdProduct)
	assert.Equal(t, newSale.Quantity, createdSale.Quantity)
	assert.Nil(t, err)
}

func TestServiceUpdateSaleOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)
	service := NewSaleService(repo)

	updatedSale := models.Sale{
		Id:        1001,
		IdProduct: 45,
		IdInvoice: 14,
		Quantity:  23,
	}

	saleUpdated, err := service.UpdateSale(updatedSale)

	assert.Nil(t, err)
	assert.NotNil(t, saleUpdated)
	assert.Equal(t, updatedSale.Id, saleUpdated.Id)
	assert.Equal(t, updatedSale.IdProduct, saleUpdated.IdProduct)
}

func TestServiceUpdateSaleError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewSaleRepository(&mock_array, db)
	service := NewSaleService(repo)

	updatedSale := models.Sale{}

	_, err = service.UpdateSale(updatedSale)

	assert.NotNil(t, err)
}
