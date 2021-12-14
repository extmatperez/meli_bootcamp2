package internal

import (
	"encoding/json"
	"testing"

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
	json.Unmarshal([]byte(prod), &expected)
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
