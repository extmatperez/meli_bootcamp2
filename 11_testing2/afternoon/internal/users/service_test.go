package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/11_testing2/afternoon/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	user := getMockUser()

	jsonData := getJsonData(user)

	dbMock := store.Mock{
		Data: jsonData,
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	editedUser, err := service.Update(user.ID, user.Name, user.LastName, user.Email, 42, user.Height, true, user.Created)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, editedUser.ID)
	assert.Equal(t, user.Name, editedUser.Name)
	assert.Equal(t, 42, editedUser.Age)
	assert.True(t, editedUser.Active)
	assert.True(t, dbMock.EnterRead)
}

func TestUpdateError(t *testing.T) {
	user := getMockUser()

	dbMock := store.Mock{
		Err: errors.New(""),
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	editedUser, err := service.Update(2, user.Name, user.LastName, user.Email, 42, user.Height, true, user.Created)

	assert.Error(t, err)
	assert.Equal(t, User{}, editedUser)
	assert.True(t, dbMock.EnterRead)
}

func TestDelete(t *testing.T) {
	user := getMockUser()

	jsonData := getJsonData(user)

	dbMock := store.Mock{
		Data: jsonData,
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	couldDelete, err := service.Delete(1)

	assert.Nil(t, err)
	assert.True(t, couldDelete)
	assert.True(t, dbMock.EnterRead)

	couldDelete, err = service.Delete(2)
	assert.Nil(t, err)
	assert.False(t, couldDelete)
}

func TestDeleteError(t *testing.T) {

	dbMock := store.Mock{
		Err: errors.New(""),
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	_, err := service.Delete(2)

	assert.Error(t, err)
	assert.True(t, dbMock.EnterRead)
}

func getMockUser() User {
	return User{
		ID:       1,
		Name:     "Juan",
		LastName: "Carlos",
		Email:    "juan.carlos@gmail.com",
		Age:      32,
		Height:   1.72,
		Active:   false,
		Created:  "11/12/2021",
	}
}

func getJsonData(user User) []byte {
	users := []User{}
	users = append(users, user)

	jsonData, _ := json.Marshal(users)
	return jsonData
}
