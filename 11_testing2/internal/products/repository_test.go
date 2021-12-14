package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productsGetAllTest = `
[{
	"id": 1,
	"name": "Pelota",
	"color": "Negro",
	"price": 1505.5,
	"stock": 200,
	"code": "#0000000f1",
	"published": true,
	"created_at": "21/11/2021"
},
{
	"id": 2,
	"name": "Botines",
	"color": "Blanco",
	"price": 5020.5,
	"stock": 50,
	"code": "#0000000f2",
	"published": false,
	"created_at": "12/10/2021"
}]`

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(productsGetAllTest), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

type StubStoreReadError struct{}

func (s *StubStoreReadError) Read(data interface{}) error {
	return fmt.Errorf("error")
}

func (s *StubStoreReadError) Write(data interface{}) error {
	return fmt.Errorf("error")
}

func TestGetAll(t *testing.T) {
	// Arrange
	stubStore := &StubStore{}
	repository := NewRepository(stubStore)
	var expectedResult []Product
	_ = json.Unmarshal([]byte(productsGetAllTest), &expectedResult)

	// Act
	result, err := repository.GetAll()

	// Assert
	assert.Equal(t, result, expectedResult, "deben ser iguales")
	assert.Nil(t, err)
}

func TestFindByID(t *testing.T) {
	// Arrange
	stubStore := &StubStore{}
	repository := NewRepository(stubStore)
	expectedResult := Product{
		Id:         1,
		Name:       "Pelota",
		Color:      "Negro",
		Price:      1505.5,
		Stock:      200,
		Code:       "#0000000f1",
		Published:  true,
		Created_at: "21/11/2021",
	}

	// Act
	result, err := repository.FindById(1)

	// Assert
	assert.Equal(t, result, expectedResult, "deben ser iguales")
	assert.Nil(t, err)
}

func TestFindByIDError(t *testing.T) {
	// Arrange
	stubStoreReadError := &StubStoreReadError{}
	repository := NewRepository(stubStoreReadError)

	// Act
	_, err := repository.FindById(1)

	// Assert
	assert.Error(t, err)
}

func TestLastID(t *testing.T) {
	// Arrange
	stubStore := &StubStore{}
	repository := NewRepository(stubStore)
	expectedResult := int64(2)

	// Act
	result, err := repository.LastId()

	// Assert
	assert.Equal(t, result, expectedResult, "deben ser iguales")
	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	// Arrange
	stubStoreReadError := &StubStoreReadError{}
	repository := NewRepository(stubStoreReadError)

	// Act
	err := repository.Delete(1)

	// Assert
	assert.Error(t, err)
}

type MockStore struct {
	id       int64
	name     string
	price    float64
	readUsed bool
}

func (m *MockStore) Read(data interface{}) error {
	m.readUsed = true
	return json.Unmarshal([]byte(productsGetAllTest), &data)
}

func (m *MockStore) Write(data interface{}) error {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		productsSlice := reflect.ValueOf(data)

		for i := 0; i < productsSlice.Len(); i++ {
			if productsSlice.Index(i).FieldByName("Id").Interface() == m.id {
				m.name = fmt.Sprint(productsSlice.Index(i).FieldByName("Name").Interface())
				m.price, _ = strconv.ParseFloat(fmt.Sprint(productsSlice.Index(i).FieldByName("Price").Interface()), 64)
				return nil
			}
			fmt.Println(productsSlice.Index(i).FieldByName("Name"))
		}
	}
	return nil
}

func TestUpdateNameAndPrice(t *testing.T) {
	// Arrange
	mockStore := &MockStore{1, "Before update", 0.00, false}
	repository := NewRepository(mockStore)

	// Act
	result, err := repository.UpdateNameAndPrice(1, "After update", 100.00)

	// Assert
	assert.Equal(t, result.Name, mockStore.name, "deben ser iguales")
	assert.Equal(t, result.Price, mockStore.price, "deben ser iguales")
	assert.Nil(t, err)
	assert.True(t, mockStore.readUsed)
}

func TestUpdateNameAndPriceNotFound(t *testing.T) {
	// Arrange
	mockStore := &MockStore{1, "Before update", 0.00, false}
	repository := NewRepository(mockStore)

	// Act
	_, err := repository.UpdateNameAndPrice(5, "After update", 100.00)

	// Assert
	assert.Error(t, err)
	assert.True(t, mockStore.readUsed)
}
