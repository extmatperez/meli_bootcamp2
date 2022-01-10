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
	return customerLines, nil
}

func NewServiceMockArray() store.SaveFile {
	return &ServiceMockArray{false}
}

func TestServiceImportAllCustomersError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)
	service := NewCustomerService(repo)
	err = service.ImportAllCustomers()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestServiceImportAllCustomersOk(t *testing.T) {
	dataBytes := []byte(customers)
	var expectedCustomers []models.Customer
	json.Unmarshal(dataBytes, &expectedCustomers)

	var c models.Customer
	db := db.StorageDB
	rows, err := db.Query("SELECT id, `lastName`, `firstName`, `condition` FROM Customer WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.LastName, &c.FirstName, &c.Condition)
		if err != nil {
			log.Fatal(err)
		}
	}

	assert.Equal(t, expectedCustomers[0].Id, c.Id)
	assert.Equal(t, expectedCustomers[0].LastName, c.LastName)
	assert.Equal(t, expectedCustomers[0].FirstName, c.FirstName)
	assert.Equal(t, expectedCustomers[0].Condition, c.Condition)
}

func TestServiceStoreCustomerOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newCustomer := models.Customer{
		LastName:  "Vega Gimenez",
		FirstName: "Rodrigo",
		Condition: "Activo",
	}

	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)
	service := NewCustomerService(repo)

	createdCustomer, err := service.StoreCustomer(newCustomer.LastName, newCustomer.FirstName, newCustomer.Condition)

	assert.Equal(t, newCustomer.LastName, createdCustomer.LastName)
	assert.Nil(t, err)
}

func TestServiceUpdateCustomerOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)
	service := NewCustomerService(repo)

	updatedCustomer := models.Customer{
		Id:        51,
		LastName:  "Vega",
		FirstName: "Rodrigo Miguel",
		Condition: "Bloqueado",
	}

	customerUpdated, err := service.UpdateCustomer(updatedCustomer)

	assert.Nil(t, err)
	assert.NotNil(t, customerUpdated)
	assert.Equal(t, updatedCustomer.Id, customerUpdated.Id)
	assert.Equal(t, updatedCustomer.LastName, customerUpdated.LastName)
}

func TestServiceUpdateCustomerError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)
	service := NewCustomerService(repo)

	updatedCustomer := models.Customer{
		Id:        52,
		LastName:  "",
		FirstName: "Rodrigo Miguel",
		Condition: "",
	}

	_, err = service.UpdateCustomer(updatedCustomer)

	assert.NotNil(t, err)
}
