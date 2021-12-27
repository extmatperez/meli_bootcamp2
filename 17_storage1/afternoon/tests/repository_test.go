package tests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/internal/users"
)

var usersJson string = `
[
	{
		"id": 1,
		"name": "Federico",
		"last_name": "Archuby",
		"email": "",
		"age": 0,
		"height": 0,
		"active": false,
		"create": ""
	}, 
	{
		"id": 2,
		"name": "Juan",
		"last_name": "BeforeUpdate",
		"email": "juan@perez.com",
		"age": 45,
		"height": 1.75,
		"active": true,
		"created": "01/12/2021"
	}
]`

type stubStore struct{}

func (s *stubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(usersJson), &data)
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

type stubStoreError struct{}

func (s *stubStoreError) Read(data interface{}) error {
	return errors.New("")
}

func (s *stubStoreError) Write(data interface{}) error {
	return errors.New("")
}

type mockStore struct {
	readCalled bool
}

func (s *mockStore) Read(data interface{}) error {
	s.readCalled = true
	return json.Unmarshal([]byte(usersJson), &data)
}

func (s *mockStore) Write(data interface{}) error {
	return nil
}

func TestLastId(t *testing.T) {
	store := stubStore{}
	repo := users.NewRepository(&store)
	var idExpected int = 2

	id, err := repo.LastID()
	assert.Equal(t, idExpected, id)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	store := stubStore{}
	repo := users.NewRepository(&store)
	var usersExpected []users.User
	err := json.Unmarshal([]byte(usersJson), &usersExpected)

	users, _ := repo.GetAll()
	assert.Equal(t, usersExpected, users)
	assert.Nil(t, err)
}

func TestGetAllError(t *testing.T) {
	store := stubStoreError{}
	repo := users.NewRepository(&store)

	_, err := repo.GetAll()
	assert.Error(t, err)
}

func TestUpdateName(t *testing.T) {
	store := mockStore{}
	repo := users.NewRepository(&store)

	var lastNameExpected string = "After Update"

	user, err := repo.UpdateLastNameAge(2, lastNameExpected, 10)
	assert.Equal(t, lastNameExpected, user.LastName)
	assert.Equal(t, 2, user.ID)
	assert.True(t, store.readCalled)
	assert.Nil(t, err)
}

func TestUpdateNameError(t *testing.T) {
	store := mockStore{}
	repo := users.NewRepository(&store)
	var userExpected users.User
	var lastName string = "After Update"

	user, err := repo.UpdateLastNameAge(5, lastName, 10)
	assert.Equal(t, userExpected, user)
	assert.Nil(t, err)
}
