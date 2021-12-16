package internal

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/pkg/store"
)

type StubRepository struct {
	useGetAll bool
}

func (s *StubRepository) Load() ([]Transaccion, error) {
	var salida []Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) GetAll() ([]Transaccion, error) {
	var salida []Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {
	return Transaccion{}, nil
}
func (s *StubRepository) FindById(id int) (Transaccion, error) {
	return Transaccion{}, nil
}
func (s *StubRepository) FilterBy(valores ...string) ([]Transaccion, error) {
	var salida []Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {
	return Transaccion{}, nil
}
func (s *StubRepository) UpdateCod(id int, codigotransaccion string) (Transaccion, error) {
	return Transaccion{}, nil
}
func (s *StubRepository) UpdateMon(id int, monto float64) (Transaccion, error) {
	return Transaccion{}, nil
}
func (s *StubRepository) Delete(id int) error {
	return nil
}

func TestUpdateService(t *testing.T) {
	//arrange
	dataByte := []byte(trans)

	dbMock := store.Mock{Data: dataByte}

	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	nameExpected := "rodri"

	//act
	result, _ := service.Update(1, "cybqf7i0bo", "Baht", 3011534.4, nameExpected, "fodoireidh0", "9/25/2021")

	//assert
	assert.Equal(t, nameExpected, result.Emisor)
	assert.Equal(t, 1, result.ID)

}

func TestUpdateServiceError(t *testing.T) {
	//arrange
	dataByte := []byte(trans)

	dbMock := store.Mock{Data: dataByte}

	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	nameExpected := "rodri"

	//act
	_, err := service.Update(5, "cybqf7i0bo", "Baht", 3011534.4, nameExpected, "fodoireidh0", "9/25/2021")

	//assert
	assert.Error(t, err)

}

func TestDeleteService(t *testing.T) {
	//arrange

	// pre borrado
	var expected []Transaccion
	json.Unmarshal([]byte(trans), &expected)

	//
	dataByte := []byte(trans)

	dbMock := store.Mock{Data: dataByte}

	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	id := 1
	expectedResult := fmt.Errorf("la transaccion %d fue eliminada", id)

	//act
	err := service.Delete(id)
	transEliminadas, _ := service.GetAll()

	//assert
	assert.Equal(t, err, expectedResult)
	assert.Equal(t, len(transEliminadas), len(expected)-1)

}

func TestDeleteServiceError(t *testing.T) {
	//arrange

	// pre borrado
	var expected []Transaccion
	json.Unmarshal([]byte(trans), &expected)

	//
	dataByte := []byte(trans)

	dbMock := store.Mock{Data: dataByte}

	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	id := 10

	//act
	err := service.Delete(id)

	//assert
	assert.Error(t, err)

}
