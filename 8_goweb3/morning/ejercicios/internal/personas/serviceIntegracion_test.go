package internal

import (
	"encoding/json"
	"testing"

	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllServices(t *testing.T) {

	sliceDeBytes, _ := json.Marshal(slicePersonas) //viene del archivo repositoryUnit_test.go
	dbMock := store.Mock{Data: sliceDeBytes}

	storeMock := store.FileStore{FileName: "", Mock: &dbMock}

	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	res, err := myService.GetAll()

	assert.Nil(t, err)
	//assert.Error(t, err)
	assert.True(t, len(res) == 2)

}