package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

var productosTest string = `[
	{
		"id": 1,
		"nombre": "Darsie",
		"color": "Turquoise",
		"precio": "$4944.33",
		"stock": 84,
		"codigo": "265b15a6-68dd-4082-ba80-a4b0a16c3d61",
		"publicado": false,
		"fechaCreacion": "21/11/2020"
	},
   {
		"id": 2,
		"nombre": "Grove",
		"color": "Crimson",
		"precio": "$3470.92",
		"stock": 654,
		"codigo": "5c62ffa5-a28a-4c08-8edf-b213d4333bb0",
		"publicado": false,
		"fechaCreacion": "15/08/2020"
	}
]`

func (s *stubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(productosTest), &data)
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	//Arrenge
	store := &stubStore{}
	repository := NewRepository(store)

	var expectedProducts []Producto
	errUnmarshal := json.Unmarshal([]byte(productosTest), &expectedProducts)

	//Act
	result, err := repository.GetAll()

	//Assert
	assert.Equal(t, expectedProducts, result)
	assert.Nil(t, err)
	assert.Nil(t, errUnmarshal)
}

type mockStore struct {
	used bool
}

func (m *mockStore) Read(data interface{}) error {

	m.used = true

	return json.Unmarshal([]byte(`[{
										"id": 1,
										"nombre": "Before update",
										"color": "Turquoise",
										"precio": "$4944.33",
										"stock": 84,
										"codigo": "265b15a6-68dd-4082-ba80-a4b0a16c3d61",
										"publicado": false,
										"fechaCreacion": "21/11/2020"
									}]`), &data)
}

func (m *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {

	//Arrenge
	store := &mockStore{}
	repository := NewRepository(store)
	nameExpected := "After update"

	//Act
	result, err := repository.Update(1, nameExpected, "Turquoise", "$4944.33", 84, "265b15a6-68dd-4082-ba80-a4b0a16c3d61", false, "21/11/2020")

	//Assert
	assert.Equal(t, nameExpected, result.Nombre)
	assert.Nil(t, err)
	assert.True(t, store.used)
}
