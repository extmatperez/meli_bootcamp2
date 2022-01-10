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

var customer string = `{
	"id": 1,
	"lastName": "Vega Gimenez",
	"firstName": "Rodrigo",
	"Condition": "Activo"
}`

var customers string = `[
	{
		"id": 1,
		"lastName": "Minshaw",
		"firstName": "Abbie",
		"Condition": "Inactivo"
	},
	{
		"id": 2,
		"lastName": "Caceres",
		"firstName": "Maximiliano",
		"Condition": "Inactivo"
	},
	{
		"id": 3,
		"lastName": "Nieto",
		"firstName": "Juan Pablo",
		"Condition": "Bloqueado"
	}
]`

var customerLines []string = []string{"1#$%#Minshaw#$%#Abbie#$%#Inactivo", "2#$%#Amberger#$%#Barnie#$%#Bloqueado", "3#$%#Close#$%#Cody#$%#Inactivo"}

func (m *MockStore) Read(data interface{}) error {
	m.CalledMethod = true
	return json.Unmarshal([]byte(customers), &data)
}

func (m *MockStore) Write(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	customers = string(byteData)
	return nil
}

func (a *MockArray) ReadLines(path string) ([]string, error) {
	a.CalledMethod = true
	return customerLines, nil
}

func NewMockStore() store.Store {
	return &MockStore{false}
}

func NewMockArray() store.SaveFile {
	return &MockArray{false}
}

func TestRepoImportAllCustomersError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)
	err = repo.ImportAllCustomers()

	assert.Error(t, errors.New("No se pudo guardar elemento en BD."), err)
}

func TestRepoImportAllCustomersOk(t *testing.T) {
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

func TestRepoStoreCustomerOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newCustomer := models.Customer{
		LastName:  "Vega Gimenez",
		FirstName: "Rodrigo",
		Condition: "Activo",
	}

	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)

	createdCustomer, err := repo.StoreCustomer(newCustomer)

	assert.Equal(t, newCustomer.LastName, createdCustomer.LastName)
	assert.Nil(t, err)
}

func TestRepoStoreCustomerError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)

	newCustomer := models.Customer{
		Id:        1,
		LastName:  "Vega Gimenez",
		FirstName: "Rodrigo",
		Condition: "Activo",
	}

	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)

	_, err0 := repo.StoreCustomer(newCustomer)

	assert.Error(t, err0)
}

func TestRepoUpdateCustomerOk(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)

	updatedCustomer := models.Customer{
		Id:        51,
		LastName:  "Vega",
		FirstName: "Rodrigo",
		Condition: "Bloqueado",
	}

	customerUpdated, err := repo.UpdateCustomer(updatedCustomer)

	assert.Nil(t, err)
	assert.NotNil(t, customerUpdated)
	assert.Equal(t, updatedCustomer.Id, customerUpdated.Id)
	assert.Equal(t, updatedCustomer.LastName, customerUpdated.LastName)
}

func TestRepoUpdateCustomerError(t *testing.T) {
	db, err := db.InitDb()
	assert.NoError(t, err)
	mock_array := MockArray{}
	repo := NewCustomerRepository(&mock_array, db)

	updatedCustomer := models.Customer{
		Id:        52,
		LastName:  "",
		FirstName: "Rodrigo Miguel",
		Condition: "",
	}

	_, err = repo.UpdateCustomer(updatedCustomer)

	assert.NotNil(t, err)
}
