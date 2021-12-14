package internal

import (
	"encoding/json"
	"testing"

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
