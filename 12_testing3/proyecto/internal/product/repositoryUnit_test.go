package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type emulatedStore struct{}

var productsData string = `[{"id":1, "nombre": "heladera", "precio":15.50}]`

func (s *emulatedStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(productsData), data)
}

func (s *emulatedStore) Write(data interface{}) error {
	return nil
}

func TestGetAllProducts(t *testing.T) {
	myStore := emulatedStore{}
	myRepo := NewRepository(&myStore)

	result, err := myRepo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
}
