package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var prod = `[{
	"id": 1,
"name": "Strawberry",
"color": "red",
"price": 23,
"stock": 30,
"code": 656,
"published": "5/17/2021",
"created": "8/28/2021"},
{"id": 2,
"name": "Watermelon",
"color": "green",
"price": 23,
"stock": 30,
"code": 656,
"published": "5/17/2021",
"created": "8/28/2021"
}]`

type StubStore struct{}

func (stab *StubStore) Write(data interface{}) error {
	return nil
}

func (stab *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(prod), &data)
}

func TestGetAll(t *testing.T) {

	//Arrange
	db := &StubStore{}
	repo := NewRepository(db)
	data, err := repo.GetAll()

	var productos []Product
	json.Unmarshal([]byte(prod), &productos)
	assert.Equal(t, data, productos, "Los datos son iguales")
	assert.Nil(t, err)

}
