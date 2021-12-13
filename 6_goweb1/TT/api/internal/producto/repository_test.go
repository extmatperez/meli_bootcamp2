package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var prod = `[
	{
	 "id": 0,
	 "name": "Mouse",
	 "color": "Black",
	 "price": 1259.99,
	 "stock": 10,
	 "code": "12abc3",
	 "isPublished": true,
	 "createdAt": "2021/10/03"
	},
	{
	 "id": 1,
	 "name": "Keyboard",
	 "color": "Black",
	 "price": 1759.99,
	 "stock": 10,
	 "code": "12abc4",
	 "isPublished": true,
	 "createdAt": "2021/10/03"
	}
   ]`

type StabStore struct {
}

type MockStore struct {
	nombre string
	read   bool
}

func (stab *StabStore) Write(data interface{}) error {
	return nil
}
func (stab *StabStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(prod), &data)
}

func (mock *MockStore) Write(data interface{}) error {
	return nil
}
func (mock *MockStore) Read(data interface{}) error {
	mock.read = true
	return json.Unmarshal([]byte(prod), &data)
}

func TestGetAll(t *testing.T) {

	db := &StabStore{}

	repo := NewRepository(db)

	data, err := repo.GetAll()

	assert.Nil(t, err)

	var productos []Product

	json.Unmarshal([]byte(prod), &productos)
	assert.Equal(t, data, productos, "Los datos no son iguales")
}

func TestUpdateName(t *testing.T) {

	db := &MockStore{"prueba", false}

	repo := NewRepository(db)

	pro, err := repo.UpdateNombre(0, "prueba")

	assert.Nil(t, err)
	assert.True(t, db.read)
	assert.Equal(t, pro.Name, db.nombre, "Los datos no son iguales")
}
