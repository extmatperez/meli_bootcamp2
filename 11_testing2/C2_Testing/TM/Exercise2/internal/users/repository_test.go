package internal

import (
	"encoding/json"
	"testing"

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

func TestGetAll(t *testing.T) {
	stubStore := StubStore{}
	repoTest := NewRepository(&stubStore)

	// fmt.Println("REPO: ", repoTest)
	myUsers, _ := repoTest.GetAll()
	// fmt.Println("USUARIOS", myUsers)

	var userExpected []User

	json.Unmarshal([]byte(usersFake), &userExpected)
	// fmt.Println("USUARIOS ESPERADOS: ", userExpected)
	assert.Equal(t, userExpected, myUsers)
}

func TestUpdateLastName(t *testing.T) {
	store := StubStore{}
	repoTest := NewRepository(&store)
	last_nameExpected := "Golang"

	userAct, _ := repoTest.UpdateLastName(2, last_nameExpected)

	assert.Equal(t, last_nameExpected, userAct.LastName)
}
