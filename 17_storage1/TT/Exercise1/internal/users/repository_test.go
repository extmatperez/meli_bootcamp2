package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/TM/Exercise1/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

var usersFake string = `[
	{"id": 1,"first_name": "Andriette","last_name": "Sanchez","email": "jsan@cloudflare.com","age": 28,"height": 112,"active": true,"cration_date": "20/08/2021"},
	{"id": 2,"first_name": "Jose","last_name": "Rios","email": "jrios@cloudflare.com","age": 28,"height": 112,"active": true,"cration_date": "20/08/2021"}]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(usersFake), &users)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAllSuccess(t *testing.T) {
	stubStore := StubStore{}
	repoTest := NewRepository(&stubStore)

	myUsers, _ := repoTest.GetAll()
	var userExpected []User
	err := json.Unmarshal([]byte(usersFake), &userExpected)

	assert.Nil(t, err)
	assert.Equal(t, userExpected, myUsers)
}

// func TestGetAllError(t *testing.T) {
// 	stubStore := StubStore{}
// 	repoTest := NewRepository(&stubStore)

// 	myUsers, err := repoTest.GetAll()

// 	assert.Nil(t, myUsers)
// 	assert.Error(t, err)
// }
func TestLastID(t *testing.T) {
	stubStore := StubStore{}
	repoTest := NewRepository(&stubStore)
	lastIdExpected := 2

	lastId, _ := repoTest.LastId()

	assert.Equal(t, lastIdExpected, lastId)
}

func TestUpdateLastNameSuccess(t *testing.T) {
	stubStore := StubStore{}
	repoTest := NewRepository(&stubStore)
	last_nameExpected := "Golang"

	userAct, _ := repoTest.UpdateLastName(2, last_nameExpected)

	assert.Equal(t, last_nameExpected, userAct.LastName)
}

func TestGetAllRepositoryMock(t *testing.T) {
	dataByte := []byte(usersFake)
	var usersExpected []User
	err := json.Unmarshal(dataByte, &usersExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	myUsers, _ := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, usersExpected, myUsers)
}
