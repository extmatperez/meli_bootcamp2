package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

var perso string = ` [
	{	"id": 1,	"nombre": "Matias",	"apellido": "Perez",	"edad": 27   },
   {	"id": 2,	"nombre": "Jose",	"apellido": "Manolo",	"edad": 22   }]`

func (s *stubStore) Read(data interface{}) error {

	return json.Unmarshal([]byte(perso), &data)
}
func (s *stubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	store := stubStore{}
	repo := NewRepository(&store)

	misPersonas, _ := repo.GetAll()
	var expected []Persona
	json.Unmarshal([]byte(perso), &expected)

	assert.Equal(t, expected, misPersonas)
}
