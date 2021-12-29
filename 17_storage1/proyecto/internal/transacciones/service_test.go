package internal

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/store"
)

type StubRepository struct {
	useGetAll bool
}

func (s *StubRepository) Load() ([]models.Transaccion, error) {
	var salida []models.Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) GetAll() ([]models.Transaccion, error) {
	var salida []models.Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {
	return models.Transaccion{}, nil
}
func (s *StubRepository) FindById(id int) (models.Transaccion, error) {
	return models.Transaccion{}, nil
}
func (s *StubRepository) FilterBy(valores ...string) ([]models.Transaccion, error) {
	var salida []models.Transaccion
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {
	return models.Transaccion{}, nil
}
func (s *StubRepository) UpdateCod(id int, codigotransaccion string) (models.Transaccion, error) {
	return models.Transaccion{}, nil
}
func (s *StubRepository) UpdateMon(id int, monto float64) (models.Transaccion, error) {
	return models.Transaccion{}, nil
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
	var expected []models.Transaccion
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
	var expected []models.Transaccion
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

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	transNueva := models.Transaccion{
		CodigoTransaccion: "15299jdhf",
		Moneda:            "dollarrr",
		Monto:             420.5,
		Emisor:            "rodri",
		Receptor:          "rorrooo",
	}

	repo := NewRepositorySql()

	service := NewServiceSql(repo)

	transCreada, err := service.Store(transNueva.CodigoTransaccion, transNueva.Emisor, transNueva.Receptor, transNueva.Moneda, transNueva.Monto)

	fmt.Println("LEEEEERRR", err)
	assert.Equal(t, transNueva.Emisor, transCreada.Emisor)
	assert.Equal(t, transNueva.Moneda, transCreada.Moneda)
	// assert.Nil(t, misPersonas)

}
