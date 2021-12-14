package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/TM/Exercise1/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

var usersFakeService string = `[
	{"id": 1,"first_name": "Andriette","last_name": "Sanchez","email": "jsan@cloudflare.com","age": 28,"height": 112,"active": true,"cration_date": "20/08/2021"},
	{"id": 2,"first_name": "Jose","last_name": "Rios","email": "jrios@cloudflare.com","age": 28,"height": 112,"active": true,"cration_date": "20/08/2021"}]`

func (s *StubRepository) GetAll() ([]User, error) {
	var out []User
	err := json.Unmarshal([]byte(usersFakeService), &out)
	s.useGetAll = true
	return out, err
}

func (s *StubRepository) Store(id int, first_name string, last_name string, email string, age int, heiht int, active bool, create_date string) (User, error) {
	return User{}, nil
}
func (s *StubRepository) Update(id int, first_name string, last_name string, email string, age int, heiht int, active bool, create_date string) (User, error) {
	return User{}, nil
}
func (s *StubRepository) UpdateLastName(id int, last_name string) (User, error) {
	return User{}, nil
}
func (s *StubRepository) UpdateAge(id int, age int) (User, error) {
	return User{}, nil
}
func (s *StubRepository) Delete(id int) error {
	return nil
}
func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	myUsers, _ := service.GetAll()

	assert.Equal(t, 2, len(myUsers))
	assert.True(t, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}

func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(usersFakeService)
	var usersExpected []User
	json.Unmarshal(dataByte, &usersExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myUsers, _ := service.GetAll()

	assert.Equal(t, usersExpected, myUsers)
}

func TestGetAllServiceMockError(t *testing.T) {
	errExpected := errors.New("No data in the mock")

	dbMock := store.Mock{Err: errExpected}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	myUsers, err := service.GetAll()

	assert.Equal(t, errExpected, err)
	assert.Nil(t, myUsers)
}

func TestStoreServiceMock(t *testing.T) {
	newUser := User{
		FirstName:   "Juan",
		LastName:    "Orfali",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	userCreated, _ := service.Store(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CrationDate)

	assert.Equal(t, newUser.FirstName, userCreated.FirstName)
	assert.Equal(t, newUser.LastName, userCreated.LastName)
}

func TestStoreServiceMockError(t *testing.T) {
	newUser := User{
		FirstName:   "Juan",
		LastName:    "Orfali",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	errExpected := errors.New("No data in the mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errExpected}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	userCreated, err := service.Store(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CrationDate)

	assert.Equal(t, errExpected, err)
	assert.Equal(t, User{}, userCreated)
}
