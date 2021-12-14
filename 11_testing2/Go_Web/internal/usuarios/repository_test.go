package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

var users string = `
[{"id":15,"nombre":"Ida","apellido":"Tieman","email":"itieman0@npr.org","edad":82,"altura":187,"activo":true,"fecha_creacion":"06/15/2021"},
{"id":2,"nombre":"Law","apellido":"Lafee","email":"llafee1@barnesandnoble.com","edad":70,"altura":142,"activo":true,"fecha_creacion":"07/12/2021"}]

`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(users), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	//Act
	misUsuarios, _ := repo.GetAll()

	var expected []Usuario
	json.Unmarshal([]byte(users), &expected)

	//Assert
	assert.Equal(t, expected, misUsuarios)

}
