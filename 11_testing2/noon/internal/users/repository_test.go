package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var usersJson string = "[{'id': 1,'name': 'Federico','last_name': 'Archuby','email': '','age': 0,'height': 0,'active': false,'create': ''}, {'id': 2,'name': 'Juan','last_name': 'Ramirez','email': 'juan@perez.com','age': 45,'height': 1.75,'active': true,'created': '01/12/2021'}]"

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(usersJson), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

type StubStoreError struct{}

func (s *StubStoreError) Read(data interface{}) error {
	return errors.New(".")
}

func (s *StubStoreError) Write(data interface{}) error {
	return errors.New(".")
}

func TestGetAll(t *testing.T) {
	store := StubStore{}
	repo := NewRepository(&store)
	var usersExpected []User
	json.Unmarshal([]byte(usersJson), &usersExpected)

	users, _ := repo.GetAll()
	assert.Equal(t, usersExpected, users)
}

func TestGetAllError(t *testing.T) {
	store := StubStoreError{}
	repo := NewRepository(&store)

	_, err := repo.GetAll()
	assert.Error(t, err)
}
