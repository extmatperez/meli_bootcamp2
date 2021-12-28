package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type storeEmulado struct{}

/* var dataPersonas string = `
[
	{
		"id": 1,
		"nombre": "donald",
		"apellido": "trump",
		"edad": 27
	},
	{
		"id": 2,
		"nombre": "jair",
		"apellido": "bolsonaro",
		"edad": 55
   }
]` */

var slicePersonas []Persona = []Persona{ {ID: 1, Nombre: "JC", Apellido: "Rossi", Edad: 27}, {ID: 2, Nombre: "JC2", Apellido: "Rossi2", Edad: 28} }

func (s *storeEmulado) Read(data interface{}) error {
	//return json.Unmarshal([]byte(dataPersonas), &data)
	sliceDeBytes, err := json.Marshal(slicePersonas)
	if err != nil {return err}
	return json.Unmarshal(sliceDeBytes, &data)
}

func (s *storeEmulado) Write(data interface{}) error {
	return nil
}

func TestGetAllPersonas(t *testing.T) {
	miStore := storeEmulado{}
	myRepo := NewRepository(&miStore)

	salida, err := myRepo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(salida))
}