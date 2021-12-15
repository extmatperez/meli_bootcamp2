package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
	useUpdate bool
}

// func (s *StubRepository) Read(data interface{}) error {
// 	return json.Unmarshal([]byte(transacciones2), &data)
// }

// func (s *StubRepository) Write(data interface{}) error {
// 	return nil
// }

var transaccionesSlice []Transaction = []Transaction{
	{ID: 1, Transaction_Code: "12345", Coin: "USD", Amount: 300.00, Emitor: "Juan", Receptor: "Enrique", Transaction_Date: "12/12/21"},
	{ID: 2, Transaction_Code: "54321", Coin: "Euro", Amount: 150.00, Emitor: "Miguel", Receptor: "Luis", Transaction_Date: "13/12/21"},
}

var transactionExample = Transaction{3, "10001", "Peso", 200, "Ivan", "Geronimo", "10/10/21"}

func (s *StubRepository) GetAll() ([]Transaction, error) {
	var salida []Transaction
	transactionsByte, _ := json.Marshal(transaccionesSlice)
	err := json.Unmarshal(transactionsByte, &salida)
	s.useGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	var transactionCreated Transaction
	transactionCreated.ID = id
	transactionCreated.Transaction_Code = transaction_code
	transactionCreated.Coin = coin
	transactionCreated.Emitor = emitor
	transactionCreated.Receptor = receptor
	transactionCreated.Transaction_Date = transaction_date

	transactionsAppended := append(transaccionesSlice, transactionCreated)
	return transactionsAppended[len(transactionsAppended)-1], nil
}

func (s *StubRepository) Update(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	transactionExample.ID = id
	transactionExample.Transaction_Code = transaction_code
	transactionExample.Coin = coin
	transactionExample.Emitor = emitor
	transactionExample.Receptor = receptor
	transactionExample.Transaction_Date = transaction_date
	s.useUpdate = true
	return transactionExample, nil
}

func (s *StubRepository) UpdateReceptor(id int, receptor string) (Transaction, error) {
	return Transaction{}, nil
}

func (s *StubRepository) Delete(id int) error {
	return nil
}

func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{useGetAll: false}
	service := NewService(&stubRepo)

	transactionsGetAll, _ := service.GetAll()

	assert.Equal(t, 2, len(transactionsGetAll))
	assert.True(t, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) {

	stubRepo := StubRepository{false, false}
	service := NewService(&stubRepo)

	transactionsGetAll, err := service.GetAll()

	assert.Equal(t, 2, transactionsGetAll[len(transactionsGetAll)-1].ID)
	assert.Nil(t, err)
}
