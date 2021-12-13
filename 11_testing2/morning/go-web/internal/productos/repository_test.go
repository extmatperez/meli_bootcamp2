package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

var testStorage string = `[
	{
		"id": 1,
		"nombre": "Before Update",
		"color": "gris",
		"precio": 999,
		"stock": 12,
		"codigo": "sfsdf 444 3 www",
		"publicado": true,
		"fechaCreacion": "12/9/1999"
	  },
	  {
		"id": 2,
		"nombre": "termo",
		"color": "gris",
		"precio": 999,
		"stock": 12,
		"codigo": "sfsdf 444 3 www",
		"publicado": true,
		"fechaCreacion": "12/9/1999"
	  }
]`

func (s *stubStore) Read(data interface{}) error {

	return json.Unmarshal([]byte(testStorage), &data)
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	store := stubStore{}
	repo := NewRepository(&store)
	var stubProducts []Producto
	json.Unmarshal([]byte(testStorage), &stubProducts)

	// Se ejecuta el test
	testProd, _ := repo.GetAll()

	// Se validan los resultados
	assert.Equal(t, stubProducts, testProd, "deben ser iguales")
}

type mockStore struct {
	invoked bool
}

func (m *mockStore) Read(data interface{}) error {
	m.invoked = true
	return json.Unmarshal([]byte(testStorage), &data)
}

func (m *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	store := mockStore{false}
	repo := NewRepository(&store)
	var mockProducts Producto
	json.Unmarshal([]byte(testStorage), &mockProducts)
	mockProducts.Nombre = "After Update"

	// Se ejecuta el test
	afterUpdate, _ := repo.UpdateName(1, "After Update")

	// Se validan los resultados
	assert.Equal(t, mockProducts, afterUpdate, "deben ser iguales")
}
