package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

func (s *StubRepository) GetAll() ([]Transaction, error) {
	var salida []Transaction
	err := json.Unmarshal([]byte(trans), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) GetByID(id int) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) GetByReceiver(receiver string) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) Store(id int, transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) CreateTransaction(transaction Transaction) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) UpdateTransaction(id int, transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) UpdateAmount(id int, amount float64) (Transaction, error) {
	return Transaction{}, nil
}
func (s *StubRepository) DeleteTransaction(id int) error {
	return nil
}
func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misPersonas, _ := service.GetAll()

	assert.Equal(t, 2, len(misPersonas))
	assert.True(t, stubRepo.useGetAll)
}

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
