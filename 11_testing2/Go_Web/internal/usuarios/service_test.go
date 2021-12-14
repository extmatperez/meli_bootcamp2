package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	usedGetAll bool
}

var userss string = `
[{"id":1,"nombre":"Ida","apellido":"Tieman","email":"itieman0@npr.org","edad":82,"altura":187,"activo":true,"fecha_creacion":"06/15/2021"},
{"id":2,"nombre":"Law","apellido":"Lafee","email":"llafee1@barnesandnoble.com","edad":70,"altura":142,"activo":true,"fecha_creacion":"07/12/2021"}]

`

func (s *StubRepository) GetAll() ([]Usuario, error) {
	var salida []Usuario
	err := json.Unmarshal([]byte(userss), &salida)
	s.usedGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Delete(id int) error {
	return nil
}

func (s *StubRepository) EditarNombreEdad(id int, nombre string, edad int) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) LastID() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misUsuarios, _ := service.GetAll()
	assert.Equal(t, 2, len(misUsuarios))
	assert.True(t, stubRepo.usedGetAll)
}
func TestDeleteService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}
