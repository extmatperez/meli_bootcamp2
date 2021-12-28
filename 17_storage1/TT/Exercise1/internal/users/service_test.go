package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/Exercise1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/Exercise1/pkg/store"
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

func TestUpdateServiceMock(t *testing.T) {
	newUser := User{
		FirstName:   "Juan",
		LastName:    "Orfali",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	userUpdate, _ := service.Update(1, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CrationDate)

	assert.Equal(t, newUser.FirstName, userUpdate.FirstName)
	assert.Equal(t, newUser.LastName, userUpdate.LastName)
	assert.Equal(t, 1, userUpdate.ID)
}

func TestUpdateServiceMockError(t *testing.T) {
	newUser := User{
		FirstName:   "Juan",
		LastName:    "Orfali",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.Update(22, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CrationDate)

	assert.NotNil(t, err)
}
func TestUpdateLastNameServiceMock(t *testing.T) {
	newLastName := "Prado"

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	userUpdate, _ := service.UpdateLastName(2, newLastName)

	assert.Equal(t, newLastName, userUpdate.LastName)
	assert.Equal(t, 2, userUpdate.ID)
}

func TestUpdateLastNameServiceMockError(t *testing.T) {
	newLastName := "Prado"

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateLastName(22, newLastName)

	assert.NotNil(t, err)
}

func TestUpdateAgeServiceMock(t *testing.T) {
	newAge := 44

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	userUpdate, _ := service.UpdateAge(2, newAge)

	assert.Equal(t, newAge, userUpdate.Age)
	assert.Equal(t, 2, userUpdate.ID)
}

func TestUpdateAgeServiceMockError(t *testing.T) {
	newAge := 44

	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateAge(22, newAge)

	assert.NotNil(t, err)
}

func TestDeleteServiceMock(t *testing.T) {
	dataByte := []byte(usersFakeService)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(2)

	assert.Nil(t, err)

	allUsers, _ := service.GetAll()

	assert.Equal(t, 1, len(allUsers))
}

func TestStoreServiceSQL(t *testing.T) {
	newUser := models.User{
		FirstName:   "Mario",
		LastName:    "Cancino",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	userCreated, _ := service.Store(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CrationDate)

	assert.Equal(t, newUser.FirstName, userCreated.FirstName)
	assert.Equal(t, newUser.LastName, userCreated.LastName)
}

func TestGetOneServiceSQL(t *testing.T) {
	newUser := models.User{
		FirstName:   "Juan",
		LastName:    "Orfali",
		Email:       "Carsan@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	userLoader := service.GetOne(1)

	assert.Equal(t, newUser.FirstName, userLoader.FirstName)
	assert.Equal(t, newUser.LastName, userLoader.LastName)
}

func TestUpdateServiceSQL(t *testing.T) {
	userUpdate := models.User{
		ID:          15,
		FirstName:   "Jaime",
		LastName:    "Martinez",
		Email:       "Martin@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	userLoader, _ := service.Update(userUpdate)

	assert.Equal(t, userUpdate.FirstName, userLoader.FirstName)
	assert.Equal(t, userUpdate.LastName, userLoader.LastName)
}

func TestUpdateServiceSQLFail(t *testing.T) {
	userUpdate := models.User{
		ID:          24,
		FirstName:   "Felipe",
		LastName:    "Martinez",
		Email:       "Martin@cloudflare.com",
		Age:         28,
		Height:      112,
		Active:      true,
		CrationDate: "20/08/2021",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	_, err := service.Update(userUpdate)

	assert.Equal(t, "User not found", err.Error())
}
