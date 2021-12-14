package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

type Store struct{}

var produc string = `[
   {
	   "id": 1,	"nombre": "nuevo","color": "123abc","precio": 150.5,"stock": 15,"codigo": "verde","publicado": true,"fecha_de_creacion": "20/11/2021" 
   },
   {
	"id": 2,"nombre": "NUEVO","color": "123abc","precio": 150.5,"stock": 15,"codigo": "verde","publicado": true,"fecha_de_creacion": "20/11/2021"
   }
   ]`

func (s *Store) Read(data interface{}) error {
	return json.Unmarshal([]byte(produc), &data)
}
func (s *Store) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	store := Store{}
	repo := NewRepository(&store)

	productos, _ := repo.GetAll()

	var resultadoEsperado []Productos
	err := json.Unmarshal([]byte(produc), &resultadoEsperado)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, resultadoEsperado, productos)

}

func TestGetAllMock(t *testing.T) {

	dataByte := []byte(produc)                           //* obtengo los bytes de product
	var productosEsperados []Productos                   //* declaro una variable
	err := json.Unmarshal(dataByte, &productosEsperados) //* guardo los datos en la variable
	if err != nil {
		assert.Error(t, err)
	}

	dbMock := store.Mock{Data: dataByte}    //* paso los datos en bytes
	store := store.FileStore{Mock: &dbMock} //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)           //* Creo un repo y le paso mi store que tiene todos los metodos

	productos, _ := repo.GetAll() //* devuelve los productos y un error

	assert.Equal(t, productosEsperados, productos)

}

func TestUpdate(t *testing.T) {

	store := Store{}
	repo := NewRepository(&store)
	nombreEsperado := "nuevoNombre"

	productoActualizado, _ := repo.Modify(1, 10, nombreEsperado, "rojo", "cod", "20/11/2021", 10.0, false)

	assert.Equal(t, nombreEsperado, productoActualizado.Nombre)

}

func TestUpdateMoke(t *testing.T) {

	store := Store{}
	repo := NewRepository(&store)
	nombreEsperado := "nuevoNombre"

	productoActualizado, _ := repo.Modify(1, 10, nombreEsperado, "rojo", "cod", "20/11/2021", 10.0, false)

	assert.Equal(t, nombreEsperado, productoActualizado.Nombre)

}

func TestLastIdMock(t *testing.T) {

	dataByte := []byte(produc)                           //* obtengo los bytes de product
	var productosEsperados []Productos                   //* declaro una variable
	err := json.Unmarshal(dataByte, &productosEsperados) //* guardo los datos en la variable
	if err != nil {
		assert.Error(t, err)
	}

	dbMock := store.Mock{Data: dataByte}    //* paso los datos en bytes
	store := store.FileStore{Mock: &dbMock} //* creo un store y le paso el mock que cree
	repo := NewRepository(&store)           //* Creo un repo y le paso mi store que tiene todos los metodos

	ultimoId, _ := repo.LastId() //* devuelve los productos y un error

	assert.Equal(t, productosEsperados[len(productosEsperados)-1].Id, ultimoId)

}

// func TestLastIdMockErr(t *testing.T) {

// 	var producErr string = `[
//    {
// 	   "id": 1,	"nombre": "nuevo","color": "123abc","precio": 150.5,"stock": 15,"codigo": "verde","publicado": true,"fecha_de_creacion": "20/11/2021"
//    },
//    {
// 	"id": 3,"nombre": "NUEVO","color": "123abc","precio": 150.5,"stock": 15,"codigo": "verde","publicado": true,"fecha_de_creacion": "20/11/2021"
//    }
//    ]`

// 	dataByte := []byte(producErr)                 //* obtengo los bytes de product
// 	var productosEsperados []Productos            //* declaro una variable
// 	json.Unmarshal(dataByte, &productosEsperados) //* guardo los datos en la variable

// 	dbMock := store.Mock{Data: dataByte}    //* paso los datos en bytes
// 	store := store.FileStore{Mock: &dbMock} //* creo un store y le paso el mock que cree
// 	repo := NewRepository(&store)           //* Creo un repo y le paso mi store que tiene todos los metodos

// 	_, err := repo.LastId() //* devuelve los productos y un error

// 	assert.Error(t, err, "Hubo un error al cargar el id")

// }
