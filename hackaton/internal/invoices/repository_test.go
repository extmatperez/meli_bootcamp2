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

var inv models.Invoice = models.Invoice{}
var invoices string = `[
	{
		"id": 1,
		"datetime": "2021-12-14 12:24:27",
		"idCustomer": 41,
		"total": 38612606.59
	},
	{
		"id": 2,
		"datetime": "2021-08-24",
		"idCustomer": 7,
		"total": 1789.45
	},
	{
		"id": 3,
		"datetime": "2021-08-25",
		"idCustomer": 8,
		"total": 2789.45
	}
]`

var invoiceLines []string = []string{"1#$%#2021-12-14 12:24:27#$%#41#$%#", "2#$%#2021-12-13 16:52:48#$%#22#$%#", "3#$%#2021-12-21 7:32:40#$%#31#$%#"}

func (m *MockStore) Read(data interface{}) error {
	m.CalledMethod = true
	return json.Unmarshal([]byte(invoices), &data)
}

func (m *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	invoices = string(byteData)
	return nil
}

func (a *MockArray) ReadLines(path string) ([]string, error) {
	a.CalledMethod = true
	return invoiceLines, nil
}

func NewMockStore() store.Store {
	return &MockStore{false}
}

func NewMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestRepoImportAllInvoicesError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)
	err = repo.ImportAllInvoices()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestRepoImportAllInvoicesOk(t *testing.T) {
	dataBytes := []byte(invoices)
	var expectedInvoices []models.Invoice
	json.Unmarshal(dataBytes, &expectedInvoices)

	var i models.Invoice
	db := db.StorageDB
	rows, err := db.Query("SELECT id, `datetime`, `idCustomer`, `total` FROM Invoice WHERE id = 1")
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

func TestRepoStoreInvoiceOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newInvoice := models.Invoice{
		Datetime:   "2021-12-13",
		IdCustomer: 4,
		Total:      6732.31,
	}

	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)

	createdInvoice, err := repo.StoreInvoice(newInvoice)

	assert.Equal(t, newInvoice.Datetime, createdInvoice.Datetime)
	assert.Nil(t, err)
}

func TestRepoStoreInvoiceError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newInvoice := models.Invoice{
		Id:         1,
		Datetime:   "2021-12-13",
		IdCustomer: 4,
		Total:      6732.31,
	}

	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)

	_, err0 := repo.StoreInvoice(newInvoice)

	assert.Error(t, err0)
}

func TestRepoUpdateInvoiceOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)

	updatedInvoice := models.Invoice{
		Id:         106,
		Datetime:   "2021-12-14",
		IdCustomer: 4,
		Total:      6732.31,
	}

	invoiceUpdated, err := repo.UpdateInvoice(updatedInvoice)

	assert.Nil(t, err)
	assert.NotNil(t, invoiceUpdated)
	assert.Equal(t, updatedInvoice.Id, invoiceUpdated.Id)
	assert.Equal(t, updatedInvoice.Datetime, invoiceUpdated.Datetime)
}

func TestRepoUpdateInvoiceError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)

	updatedInvoice := models.Invoice{
		Id:         101,
		Datetime:   "",
		IdCustomer: 8,
		Total:      0,
	}

	_, err = repo.UpdateInvoice(updatedInvoice)

	assert.NotNil(t, err)
}

func TestRepoUpdateInvoiceTotalsError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewInvoiceRepository(&mock_array, db)

	err = repo.UpdateTotalsOfInvoices()
	assert.Equal(t, "No se pudo modificar el total del registro.", err.Error())
}
