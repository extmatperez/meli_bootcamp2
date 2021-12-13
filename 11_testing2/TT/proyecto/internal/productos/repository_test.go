/*
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
	- Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

/*
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del
nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage
para buscar el producto. Para esto:
	- Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
	- El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través
	de un boolean como se observó en la clase.
	- Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado
	y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya
	sido ejecutado durante el test.
*/

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
	json.Unmarshal([]byte(productosTest), &expectedProducts)

	//Act
	result, err := repository.GetAll()

	//Assert
	assert.Equal(t, expectedProducts, result)
	assert.Nil(t, err)
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
