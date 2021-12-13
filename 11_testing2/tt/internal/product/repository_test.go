package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}
type MockStorage struct {
	IsRead bool
}

var prod string = `[ 
	{	"id": 1,	"color": "white",	"price": "7.35",	"amount": 11   },
   	{	"id": 2,	"color": "yellow",	"price": "8.20",	"amount": 25   }]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(prod), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.IsRead = true
	return json.Unmarshal([]byte(prod), &data)
}

func (m *MockStorage) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//Act
	products, _ := repo.GetAll()
	var expected []Product
	json.Unmarshal([]byte(prod), &expected)

	//Assert
	assert.Equal(t, expected, products)
}

func TestUpdatePrice(t *testing.T) {
	//Arrange
	store := MockStorage{}
	repo := NewRepository(&store)

	store.IsRead = false

	//Act
	updatedProduct, _ := repo.UpdatePrice(2, 3.50)
	var expected []Product
	json.Unmarshal([]byte(prod), &expected)

	//Assert
	assert.Equal(t, true, store.IsRead)
	assert.Contains(t, expected, updatedProduct)
}
