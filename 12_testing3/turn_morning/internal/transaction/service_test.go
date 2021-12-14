package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_morning/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTransactionServiceMock(t *testing.T) {

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
	service := NewService(repository)
	id := 2

	//Act
	transUpdated, err := service.UpdateTransaction(id,
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

func TestDeleteTransactionServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(trans)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.DeleteTransaction(2)

	assert.Nil(t, err)

	all, _ := service.GetAll()

	assert.Equal(t, 1, len(all))
}

func TestDeleteTransactionErrorServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(trans)

	id := 6
	expectedErrorMsg := fmt.Sprintf("The transaction %d no exists", id)
	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.DeleteTransaction(id)

	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}
