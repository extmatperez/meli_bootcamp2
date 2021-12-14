package internal

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s *StubStore) Read(data interface{}) (bool, error) {
	return true, json.Unmarshal([]byte(trans), &data)
}

func (s *StubStore) Write(data interface{}) (bool, error) {

	return true, nil
}

type SpyConfig struct {
	isReadOnly bool
}

func (s *SpyConfig) Read(data interface{}) (bool, error) {
	s.isReadOnly = true
	return true, json.Unmarshal([]byte(trans), &data)
}

func (s *SpyConfig) Write(data interface{}) (bool, error) {
	s.isReadOnly = true
	return true, nil
}

type SpyErrorConfig struct {
	isReadOnly bool
}

func (s *SpyErrorConfig) Read(data interface{}) (bool, error) {
	s.isReadOnly = true
	return true, json.Unmarshal([]byte(trans), &data)
}

func (s *SpyErrorConfig) Write(data interface{}) (bool, error) {
	s.isReadOnly = true
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(trans, file, 0644)
}

var trans string = `[
{
 "id": 1,
 "transaction_code": "68828-179",
 "currency": "Bolivares",
 "amount": 600.00,
 "receiver": "Digneli",
 "sender": "Maria",
 "transaction_date": "8/11/2021"
},
{
 "id": 2,
 "transaction_code": "68828-17963",
 "currency": "Bolivares",
 "amount": 250.00,
 "receiver": "Josue",
 "sender": "Sofi",
 "transaction_date": "8/10/2021"
}]`

func TestGetAll(t *testing.T) {

	//Arrange
	store := StubStore{}
	repository := NewRepository(&store)
	//Act
	transactions, _ := repository.GetAll()
	var expected []Transaction
	json.Unmarshal([]byte(trans), &expected)

	//Assert
	assert.Equal(t, expected, transactions, "Should be equals")
}

func TestUpdateAmount(t *testing.T) {

	///Arrange
	store := &SpyConfig{}
	repository := NewRepository(store)
	amountExpected := 100.00
	id := 1

	//Act
	transUpdated, err := repository.UpdateAmount(id, amountExpected)

	//Assert
	assert.Equal(t, amountExpected, transUpdated.Amount, "Should be equals")
	assert.Nil(t, err)
	assert.Equal(t, true, store.isReadOnly)
}

func TestUpdateAmountWithError(t *testing.T) {

	///Arrange
	store := &SpyErrorConfig{}
	repository := NewRepository(store)
	amountExpected := 100.00
	id := 2

	//Act
	_, err := repository.UpdateAmount(id, amountExpected)

	//Assert
	assert.Error(t, err)
}
