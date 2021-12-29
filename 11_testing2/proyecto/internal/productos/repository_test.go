package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

var TestStorage string = `[
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

type studStore struct{}

// creamos estructura vacia

func (s *studStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(TestStorage), &data)
	// traemos las funciones del store que son Read y Write, que serian los dos metodos del Store
	// en el caso del Read unmarshaleamos
}

func (s *studStore) Write(data interface{}) error {
	return nil
}

//funcion para testear GetAll()
func TestGetAll(t *testing.T) {
	//creamos variables
	store := studStore{}
	//nuevo store, que va a conectar con mi base de datos falsa, que creamos aca
	repo := NewRepository(&store)
	// creamos un repo con este store. la funcion crea el repo falso

	var studProducts []Producto
	//creamos slice del tipo Productoque traemos del repo (testigo), en la cual vamos a almacenar el unmarshall
	json.Unmarshal([]byte(TestStorage), &studProducts)
	//unmarshal de la base de datos falsa, y almacenamos en el slice

	testProduct, _ := repo.GetAll()
	// llamamos a la func GetAll (conectaod a la bd falsa)
	//testProduct va a tener lo mismo que studProduct, nada masque testProduct va a ser completado luego de
	// pasar por la funcion GetAll

	assert.Equal(t, studProducts, testProduct)
	// comparamos el slice studProduct con el que paso x GetAll, es decir testProduct
}

type mockStore struct {
	usado bool
	// la variable usado es para hacer un spy, que nos devuelva si Read fue usado.
}

func (m *mockStore) Read(data interface{}) error {
	m.usado = true
	return json.Unmarshal([]byte(TestStorage), &data)

}

func (m *mockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {

	store := mockStore{usado: false}
	// partimos de la bse de false y luego lo comprobamos
	repo := NewRepository(&store)
	var mockProducts []Producto

	json.Unmarshal([]byte(TestStorage), &mockProducts)

	primerProducto := mockProducts[0]

	// le cambiamos el nombre al primer producto de nuestra db falsa hardcodeado.
	primerProducto.Nombre = "Despues"

	// usamos la func UpdateName y le almacenamos el valor por el cual queremos cambiar (id, luego el nuevo nombre)
	prod, _ := repo.UpdateName(1, "Despues")

	assert.True(t, store.usado, "deberia haberse invocado el store") // deberia haberse invocado el store (verdadero o falso)
	assert.Equal(t, primerProducto, prod)                            // le pasamos,t, y luego los dos valores comparar

}

func TestStore(t *testing.T) {
	// creamos myMockStore que es un nuevo store que traemos del package store, de la strcut MockStore, a la cual le tenemos que pasar
	// un slice de bytes. En este caso usamos TestStorage que lo creamos arriba.
	myMockStore := store.MockStore{Data: []byte(TestStorage)}

	// creamos una struct de tipo FileStore, la cual teien los campos FileName y MockStore. Creamos MockStore antes para poder mandarsela
	// ahora al repo. El repo necesita una estructura de tipo FileStore, por eso hacemos todo este quilombo.
	//Basicamente estamos cumpliendo con lo que nos pide la funcion NewRepository.

	store := store.FileStore{Mock: &myMockStore}

	//ahora si me puedo crear el repo con store.
	repo := NewRepository(&store)

	// como tmb queresmos testear el service, metemos el repo dentro del service y testeamos.
	service := NewService(repo)

	mockTest, _ := service.Store("After Update", "gris", 999, 12, "sfsdf 444 3 www", true, "12/9/1999")
	// aca testeamos directamente ddesde el service, que , si funciona, va implicar que ambas capas funcionan de manera correcta.

	assert.Equal(t, "After Update", mockTest.Nombre, "deben coincidir")
	//aca lo que hacemos es ver que el 2do parametro "After Update" QUE LE PASO en el assert, coincida con el mockTest.Nombre que fue procesado
	// mockTest va a ser la respuesta del metodo store del service que es una estructura de tipo Producto

}

// el service necesita del repo para funcionar, es por eso que necesariamente le tenemos que pasar el repo como parametro.
// en este caso necesitamos testear service y repo. Entonces nosotros le mandamos la data la service, el service se la manda al repo,
// y luego el repo devuelve una respuesta.
