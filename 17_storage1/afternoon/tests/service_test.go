package tests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/store"
	"github.com/stretchr/testify/assert"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/internal/users"
)

func TestServiceGetAll(t *testing.T) {
	user := getMockUser()

	usersExpected := []users.User{}
	usersExpected = append(usersExpected, user)

	jsonData := getJsonData(user)

	dbMock := store.Mock{
		Data: jsonData,
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	usersObtained, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, usersExpected, usersObtained)
	assert.True(t, dbMock.EnterRead)
}

func TestServiceGetAllError(t *testing.T) {

	dbMock := store.Mock{
		Err: errors.New(""),
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	_, err := service.GetAll()

	assert.Error(t, err)
	assert.True(t, dbMock.EnterRead)
}

func TestStore(t *testing.T) {
	user := getMockUser()

	jsonData := getJsonData(user)

	dbMock := store.Mock{
		Data: jsonData,
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	editedUser, err := service.Store(user.Name, user.LastName, user.Email, 42, user.Height, true, user.Created)

	assert.Nil(t, err)
	assert.Equal(t, 2, editedUser.ID)
	assert.Equal(t, user.Name, editedUser.Name)
	assert.Equal(t, 42, editedUser.Age)
	assert.True(t, editedUser.Active)
	assert.True(t, dbMock.EnterRead)
}

func TestStoreError(t *testing.T) {
	user := getMockUser()

	dbMock := store.Mock{
		Err: errors.New(""),
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	editedUser, err := service.Store(user.Name, user.LastName, user.Email, 42, user.Height, true, user.Created)

	assert.Error(t, err)
	assert.Equal(t, users.User{}, editedUser)
	assert.True(t, dbMock.EnterRead)
}

func TestUpdate(t *testing.T) {
	user := getMockUser()

	jsonData := getJsonData(user)

	dbMock := store.Mock{
		Data: jsonData,
	}
	storeStub := store.FileStore{
		Mock: &dbMock,
	}

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

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

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	editedUser, err := service.Update(2, user.Name, user.LastName, user.Email, 42, user.Height, true, user.Created)

	assert.Error(t, err)
	assert.Equal(t, users.User{}, editedUser)
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

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

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

	repo := users.NewRepository(&storeStub)
	service := users.NewService(repo)

	_, err := service.Delete(2)

	assert.Error(t, err)
	assert.True(t, dbMock.EnterRead)
}

func getMockUser() users.User {
	return users.User{
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

func getJsonData(user users.User) []byte {
	users := []users.User{}
	users = append(users, user)

	jsonData, _ := json.Marshal(users)
	return jsonData
}
