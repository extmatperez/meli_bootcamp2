package internal

import (
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/12_testing3/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllService(t *testing.T) {

	dbMock := store.Mock{Data: []byte(productsData)}
	storeMock := store.FileStore{FileName: "", Mock: &dbMock}

	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(result) == 2)
}

func TestGetAllServiceError(t *testing.T) {

	errorCreated := errors.New("holaaa error!")
	dbMock := store.Mock{Err: errorCreated}
	storeMock := store.FileStore{FileName: "", Mock: &dbMock}

	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	_, err := myService.GetAll()

	assert.ErrorAs(t, errorCreated, err)
}
