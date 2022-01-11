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
	return invoiceLines, nil
}

func NewServiceMockArray() store.SaveFile {
	return &ServiceMockArray{false}
}

func TestServiceImportAllInvoicesError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	service := NewInvoiceService(repo)
	err = service.ImportAllInvoices()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestServiceImportAllInvoicesOk(t *testing.T) {
	dataBytes := []byte(invoices)
	var expectedInvoices []models.Invoice
	json.Unmarshal(dataBytes, &expectedInvoices)

	var i models.Invoice
	db := db.StorageDB
	rows, err := db.Query("SELECT id, `datetime`, idCustomer, total FROM Invoice WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&i.Id, &i.Datetime, &i.IdCustomer, &i.Total)
		if err != nil {
			log.Fatal(err)
		}
	}

	assert.Equal(t, expectedInvoices[0].Id, i.Id)
	assert.Equal(t, expectedInvoices[0].Datetime, i.Datetime)
	assert.Equal(t, expectedInvoices[0].IdCustomer, i.IdCustomer)
	assert.Equal(t, expectedInvoices[0].Total, i.Total)
}

func TestServiceStoreInvoiceOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newInvoice := models.Invoice{
		Datetime:   "2021-12-13",
		IdCustomer: 4,
		Total:      6732.31,
	}

	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	service := NewInvoiceService(repo)

	createdInvoice, err := service.StoreInvoice(newInvoice.Datetime, newInvoice.IdCustomer, newInvoice.Total)

	assert.Equal(t, newInvoice.Datetime, createdInvoice.Datetime)
	assert.Nil(t, err)
}

func TestServiceUpdateInvoiceOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	service := NewInvoiceService(repo)

	updatedInvoice := models.Invoice{
		Id:         106,
		Datetime:   "2021-12-15",
		IdCustomer: 4,
		Total:      6732.31,
	}

	invoiceUpdated, err := service.UpdateInvoice(updatedInvoice)

	assert.Nil(t, err)
	assert.NotNil(t, invoiceUpdated)
	assert.Equal(t, updatedInvoice.Id, invoiceUpdated.Id)
	assert.Equal(t, updatedInvoice.Datetime, invoiceUpdated.Datetime)
}

func TestServiceUpdateInvoiceError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	service := NewInvoiceService(repo)

	updatedInvoice := models.Invoice{
		Id:         101,
		Datetime:   "",
		IdCustomer: 8,
		Total:      0,
	}

	_, err = service.UpdateInvoice(updatedInvoice)

	assert.NotNil(t, err)
}

func TestServiceUpdateInvoiceTotalsError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	service := NewInvoiceService(repo)

	err = service.UpdateTotalsOfInvoices()
	assert.Equal(t, "No se pudo modificar el total del registro.", err.Error())
}
