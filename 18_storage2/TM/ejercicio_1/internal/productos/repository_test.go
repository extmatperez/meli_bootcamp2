package internal

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/18_storage2/TM/ejercicio_1/internal/models"
	"github.com/stretchr/testify/assert"
)

type StubStrore struct {
	useRead bool
}

var prod string = `[
	{"id":1,"name":"Before Update","color":"Crimson","price":12.53,"stock":1,"code":"50438-106","publish":true,"date":"4/4/2021"},
	{"id":2,"name":"Fuel","color":"Purple","price":49.73,"stock":2,"code":"0135-0484","publish":false,"date":"4/26/2021"}]`

func (s *StubStrore) Read(data interface{}) error {
	s.useRead = true
	return json.Unmarshal([]byte(prod), &data)
}
func (s *StubStrore) Write(data interface{}) error {
	return nil
}

func TestGetall(t *testing.T) {
	store := StubStrore{}
	repo := NewRepository(&store)

	misProducts, _ := repo.GetAll()

	var expected []Product
	err := json.Unmarshal([]byte(prod), &expected)
	assert.Nil(t, err)
	assert.Equal(t, misProducts, expected)
}

func TestUpdate(t *testing.T) {
	store := StubStrore{false}
	repo := NewRepository(&store)
	expected := "After Update"
	productUpdate, _ := repo.Update(1, expected, "gris", 12.66, 2, "a1", true, "2020/09")

	assert.Equal(t, productUpdate.Name, expected)
	assert.True(t, store.useRead)
}

func TestStore(t *testing.T) {
	store := StubStrore{false}
	repo := NewRepository(&store)
	expected := "Producto prueba"
	lastId, _ := repo.LastId()
	producto, err := repo.Store(lastId, expected, "negro", 35600.00, 2, "b11", true, "2020/09")
	assert.Nil(t, err)
	assert.Equal(t, expected, producto.Name)

}

func TestUpdateNamePrice(t *testing.T) {
	store := StubStrore{false}
	repo := NewRepository(&store)
	expected := "After Update"
	productUpdate, err := repo.UpdateNamePrice(1, expected, 16.00)

	assert.Nil(t, err)
	assert.Equal(t, productUpdate.Name, expected)
	assert.True(t, store.useRead)
}

func TestStoreSQL(t *testing.T) {
	newProd := models.Product{
		Name:  "play 4",
		Color: "negro",
		Price: 270000,
	}

	repo := NewRepositorySQL()

	prodCreate, _ := repo.Store(newProd)

	assert.Equal(t, newProd.Name, prodCreate.Name)
}

func TestGetOneRepoSQL(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	productsCarge, err := repo.GetOneName("play 4")

	assert.Nil(t, err)
	assert.True(t, len(productsCarge) >= 0)

}

func TestGetAllRepoSQL(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	productsCarge, err := repo.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(productsCarge) >= 0)

}

func TestUpdateSQL(t *testing.T) {
	newProd := models.Product{
		Id:    1,
		Name:  "play 5",
		Color: "blanco",
		Price: 470000,
	}

	repo := NewRepositorySQL()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prodUpdate, _ := repo.Update(ctx, newProd)

	assert.Equal(t, newProd.Name, prodUpdate.Name)
}
