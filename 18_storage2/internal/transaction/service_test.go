package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/pkg/store"
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

/*func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	newTransaction := models.Transaction{
		TransactionCode: "123-345",
		Currency:        "$",
		Amount:          100.50,
		Receiver:        "Rosa",
		Sender:          "Cele",
		TransactionDate: "12/12/2021",
	}

	repo := NewRepositoryDB()

	service := NewServiceDB(repo)

	transactionCreated, _ := service.Store(newTransaction.Amount, newTransaction.TransactionCode, newTransaction.Currency, newTransaction.Receiver, newTransaction.Sender, newTransaction.TransactionDate)

	assert.Equal(t, newTransaction.Sender, transactionCreated.Sender)
	assert.Equal(t, newTransaction.Receiver, transactionCreated.Receiver)

}*/

func TestGetSenderServiceSQL(t *testing.T) {
	//Arrange

	repo := NewRepositoryDB()

	service := NewServiceDB(repo)

	transaction, _ := service.GetBySender("Dig")

	assert.Equal(t, "Dig", transaction.Sender)
	assert.Equal(t, "Sol", transaction.Receiver)

}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	transactionUpdated := models.Transaction{
		ID:              6,
		TransactionCode: "1234-32332",
		Currency:        "$$",
		Amount:          300.44,
		Receiver:        "Rosita",
		Sender:          "Garcia",
		TransactionDate: "14/05/2020",
	}

	repo := NewRepositoryDB()

	service := NewServiceDB(repo)

	transactionBefore := service.GetOne(transactionUpdated.ID)

	transactionActual, _ := service.Update(transactionUpdated)

	assert.Equal(t, transactionUpdated.Sender, transactionActual.Sender)
	assert.Equal(t, transactionUpdated.Receiver, transactionActual.Receiver)
	// assert.Nil(t, misPersonas)
	_, err := service.Update(transactionBefore)

	assert.Nil(t, err)
}

func TestUpdateServiceSQL_Failed(t *testing.T) {
	//Arrange
	transactionUpdated := models.Transaction{
		ID:              15,
		TransactionCode: "1234-3232",
		Currency:        "$",
		Amount:          25.44,
		Receiver:        "Rivera",
		Sender:          "Soto",
		TransactionDate: "14/05/2020",
	}

	repo := NewRepositoryDB()

	service := NewServiceDB(repo)

	_, err := service.Update(transactionUpdated)

	assert.Equal(t, "No se encontro la transaction", err.Error())
	// assert.Nil(t, misPersonas)
}
