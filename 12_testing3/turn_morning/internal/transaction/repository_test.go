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
	return false, json.Unmarshal([]byte(""), &data)
}

func (s *SpyErrorConfig) Write(data interface{}) (bool, error) {
	s.isReadOnly = true
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(trans, file, 0644)
}

type StubRealStore struct{}

func (s *StubRealStore) Read(data interface{}) (bool, error) {
	return true, json.Unmarshal([]byte(trans), &data)
}

func (s *StubRealStore) Write(data interface{}) (bool, error) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return false, err
	}
	return true, os.WriteFile("./transactiontest.json", file, 0644)
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
	err := json.Unmarshal([]byte(trans), &expected)

	//Assert
	assert.Equal(t, expected, transactions, "Should be equals")
	assert.Nil(t, err)
}

func TestGetAllWithError(t *testing.T) {

	//Arrange
	store := SpyErrorConfig{}
	repository := NewRepository(&store)
	//Act
	_, err := repository.GetAll()

	//Assert
	assert.Error(t, err)
}

func TestGetByID(t *testing.T) {

	//Arrange
	store := StubStore{}
	repository := NewRepository(&store)
	id := 2
	//Act
	transaction, _ := repository.GetByID(id)
	var expected []Transaction
	err := json.Unmarshal([]byte(trans), &expected)

	//Assert
	assert.Equal(t, expected[1], transaction, "Should be equals")
	assert.Nil(t, err)
}

func TestGetByIDWithError(t *testing.T) {

	//Arrange
	store := SpyErrorConfig{}
	repository := NewRepository(&store)
	id := 2
	//Act
	_, err := repository.GetByID(id)

	//Assert
	assert.Error(t, err)
}

func TestLastId(t *testing.T) {

	///Arrange
	store := &StubStore{}
	repository := NewRepository(store)
	idExpected := 2

	//Act
	lastId, err := repository.LastId()

	//Assert
	assert.Equal(t, idExpected, lastId, "Should be equals")
	assert.Nil(t, err)
}

func TestLastIdWithError(t *testing.T) {

	///Arrange
	store := &SpyErrorConfig{}
	repository := NewRepository(store)

	//Act
	_, err := repository.LastId()

	//Assert
	assert.Error(t, err)
}

func TestGetByReceiver(t *testing.T) {

	//Arrange
	store := StubStore{}
	repository := NewRepository(&store)
	receiver := "Digneli"
	//Act
	transaction, _ := repository.GetByReceiver(receiver)
	var expected []Transaction
	err := json.Unmarshal([]byte(trans), &expected)
	//Assert
	assert.Equal(t, expected[0], transaction, "Should be equals")
	assert.Nil(t, err)
}

func TestGetByReceiverWithError(t *testing.T) {

	//Arrange
	store := SpyErrorConfig{}
	repository := NewRepository(&store)
	receiver := "Digneli"
	//Act
	_, err := repository.GetByReceiver(receiver)
	//Assert
	assert.Error(t, err)
}

func TestStoreRepositoryMock(t *testing.T) {
	//Arrange
	newTransaction := Transaction{
		TransactionCode: "2332-3232",
		Currency:        "USD",
		Amount:          10.00,
		Receiver:        "Digneli",
		Sender:          "Rami",
		TransactionDate: "12-12-21",
	}

	store := StubStore{}
	repository := NewRepository(&store)

	service := NewService(repository)

	transactionCreated, _ := service.Store(
		newTransaction.TransactionCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Receiver,
		newTransaction.Sender,
		newTransaction.TransactionDate)

	assert.Equal(t, newTransaction.Receiver, transactionCreated.Receiver)
	assert.Equal(t, newTransaction.Sender, transactionCreated.Sender)
}

func TestStoreRepositoryNewTransMock(t *testing.T) {
	//Arrange
	newTransaction := Transaction{
		TransactionCode: "2332-3232",
		Currency:        "USD",
		Amount:          10.00,
		Receiver:        "Digneli",
		Sender:          "Rami",
		TransactionDate: "12-12-21",
	}

	dbMock := &SpyErrorConfig{}
	repository := NewRepository(dbMock)

	id, _ := repository.LastId()

	_, err := repository.Store(id+1,
		newTransaction.TransactionCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Receiver,
		newTransaction.Sender,
		newTransaction.TransactionDate,
	)

	assert.Error(t, err)
}

func TestSCreateTransactionMock(t *testing.T) {
	//Arrange
	newTransaction := Transaction{
		TransactionCode: "2332-3232",
		Currency:        "USD",
		Amount:          10.00,
		Receiver:        "Digneli",
		Sender:          "Rami",
		TransactionDate: "12-12-21",
	}

	store := StubRealStore{}
	repository := NewRepository(&store)

	service := NewService(repository)

	transactionCreated, _ := service.CreateTransaction(newTransaction)

	file, _ := os.ReadFile("./transactiontest.json")

	var transactions []Transaction
	json.Unmarshal(file, &transactions)
	id, _ := repository.LastId()

	assert.Equal(t, transactions[id].Receiver, transactionCreated.Receiver)
	assert.Equal(t, transactions[id].Sender, transactionCreated.Sender)
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

func TestUpdateTransaction(t *testing.T) {

	///Arrange
	transactionUpdated := Transaction{
		TransactionCode: "2332-3232",
		Currency:        "USD",
		Amount:          10.00,
		Receiver:        "Maria",
		Sender:          "Rami",
		TransactionDate: "12-12-21",
	}

	store := &StubRealStore{}
	repository := NewRepository(store)
	id := 2

	//Act
	transUpdated, err := repository.UpdateTransaction(id,
		transactionUpdated.TransactionCode,
		transactionUpdated.Currency,
		transactionUpdated.Amount,
		transactionUpdated.Receiver,
		transactionUpdated.Sender,
		transactionUpdated.TransactionDate)

	file, _ := os.ReadFile("./transactiontest.json")

	var transactions []Transaction
	json.Unmarshal(file, &transactions)
	//Assert
	assert.Equal(t, transactions[id-1].Amount, transUpdated.Amount, "Should be equals")
	assert.Nil(t, err)
}

func TestUpdateTransactionWithError(t *testing.T) {

	///Arrange
	transactionUpdated := Transaction{
		TransactionCode: "2332-3232",
		Currency:        "USD",
		Amount:          10.00,
		Receiver:        "Maria",
		Sender:          "Rami",
		TransactionDate: "12-12-21",
	}

	store := &StubStore{}
	repository := NewRepository(store)
	id := 4

	//Act
	transUpdated, err := repository.UpdateTransaction(id,
		transactionUpdated.TransactionCode,
		transactionUpdated.Currency,
		transactionUpdated.Amount,
		transactionUpdated.Receiver,
		transactionUpdated.Sender,
		transactionUpdated.TransactionDate)

	//Assert
	assert.Equal(t, "", transUpdated.Sender, "Should be equals")
	assert.Nil(t, err)
}
