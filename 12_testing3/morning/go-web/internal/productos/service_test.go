package internal

import (
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/afternoon/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

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

func TestUptate(t *testing.T) {

	// Se inicializan los datos a usar en el test (input/output)
	myMockStore := store.MockStore{Data: []byte(testStorage), Invoked: false}
	//fmt.Println(string(myMockStore.Data))

	store := store.FileStore{Mock: &myMockStore}

	repo := NewRepository(&store)
	service := NewService(repo)

	productoTestigo := Producto{
		Id:            1,
		Nombre:        "After Update",
		Color:         "rojo",
		Precio:        123,
		Stock:         44,
		Codigo:        "cod 123",
		Publicado:     false,
		FechaCreacion: "12/9/1999",
	}

	// Se ejecuta el test
	respuesta, _ := service.Update(1, "After Update", "rojo", 123, 44, "cod 123", false, "12/9/1999")

	// Se validan los resultados
	//valido que el Read fue utilizado
	assert.True(t, myMockStore.Invoked)
	//valido que Update funciona correctamente comparando mi producto testigo con el que recibi
	assert.Equal(t, productoTestigo, respuesta, "deben ser iguales")
}

func TestDelete(t *testing.T) { //test para cuando el ID si existe

	// Se inicializan los datos a usar en el test (input/output)
	myMockStore := store.MockStore{Data: []byte(testStorage)}
	//fmt.Println(string(myMockStore.Data))

	store := store.FileStore{Mock: &myMockStore}

	repo := NewRepository(&store)
	service := NewService(repo)

	// Se ejecuta el test
	//para ver que efectivamente se borro el producto, hago dos pruebas, intentando borrar el mismo registro dos veces
	respuesta1, err1 := service.Delete(1)
	respuesta2, err2 := service.Delete(1)

	// Se validan los resultados
	//sabiendo las respuestas de Delete, hago las validaciones pertinentes
	//cuando SI existe el producto:
	assert.Equal(t, "Producto eliminado", respuesta1, "deben ser iguales")
	assert.Nil(t, err1, "debe ser nil")

	//cuando NO existe el producto:
	assert.Equal(t, "", respuesta2, "deben ser iguales")
	assert.NotNil(t, err2, "debe ser distinto de nil")

	fmt.Println("1------", respuesta1)
	fmt.Println("2------", respuesta2)
	fmt.Println("3------", err1)
	fmt.Println("4------", err2)
}
