package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubResposirity struct {
	useGetAll bool
}

var productss string = `[
		{"id": 1, "nombre": "producto1modificadoconpatch","color": "rojo", "precio": 20, "stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"},
   		{"id": 2,"nombre": "producto1","color": "rojo","precio": 20,"stock": "alguno","codigo": "SADFHJK9","publicado": true,"fecha_creacion": "01/12/2021"
	}]`

func (s *StubResposirity) GetAll() ([]Product, error) {
	var salida []Product
	err := json.Unmarshal([]byte(productss), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubResposirity) Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) UpdateNombre(id int, nombre string) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) UpdatePrecio(id int, precio int) (Product, error) {
	return Product{}, nil
}
func (s *StubResposirity) Delete(id int) error {
	return nil
}
func (s *StubResposirity) LastID() (int, error) {
	return 0, nil
}
func TestGetAllS(t *testing.T) {
	stubRepo := StubResposirity{}
	service := NewService(&stubRepo)
	misProductos, _ := service.GetAll()

	assert.Equal(t, 2, len(misProductos))
	assert.True(t, true, stubRepo.useGetAll)
}
