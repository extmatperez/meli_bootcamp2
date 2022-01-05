package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/18_storage2/TM/ejercicio_1/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	UseGetAll bool
}

var produ string = `[
	{"id":1,"name":"Before Update","color":"Crimson","price":12.53,"stock":1,"code":"50438-106","publish":true,"date":"4/4/2021"},
	{"id":2,"name":"Fuel","color":"Purple","price":49.73,"stock":2,"code":"0135-0484","publish":false,"date":"4/26/2021"}]`

func (s *StubRepository) GetAll() ([]Product, error) {
	var salida []Product
	err := json.Unmarshal([]byte(produ), &salida)
	s.UseGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, name string, color string, code int) (Product, error) {
	return Product{}, nil
}
func TestUpdateServiceMock(t *testing.T) {
	//Arrange
	prodNew := Product{
		Name:  "producto nuevo",
		Color: "gris",
		Code:  "a23",
	}

	dataByte := []byte(prod)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productUpdate, _ := service.Update(1, prodNew.Name, prodNew.Color, prodNew.Price, prodNew.Stock, prodNew.Code, prodNew.Publish, prodNew.Date)

	assert.Equal(t, prodNew.Name, productUpdate.Name)
	assert.Equal(t, prodNew.Code, productUpdate.Code)
	assert.Equal(t, 1, productUpdate.Id)
}

func TestDeleteMock(t *testing.T) {

	dataByte := []byte(prod)
	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err, "se borra producto")

	proDeleted, _ := service.GetAll()
	println(proDeleted[0].Name)

	// // println("--", produ, "--", proDeleted)
	// // assert.Equal(t, len(produ), len(proDeleted))

}

func TestStoreService(t *testing.T) {
	//Arrange
	prodNew := Product{
		Id:      1,
		Name:    "Before Update",
		Color:   "Crimson",
		Price:   12.53,
		Stock:   1,
		Code:    "50438-106",
		Publish: true,
		Date:    "4/4/2021",
	}

	dataByte := []byte(prod)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	productStore, err := service.Store(prodNew.Name, prodNew.Color, prodNew.Price, prodNew.Stock, prodNew.Code, prodNew.Publish, prodNew.Date)

	assert.Nil(t, err)
	assert.Equal(t, prodNew.Name, productStore.Name)
	assert.Equal(t, prodNew.Code, productStore.Code)
}
