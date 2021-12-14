package internal

import (
	"encoding/json"
)

type StubRepository struct {
	useUpdate bool
}

func (s *StubRepository) Read(data interface{}) error {
	return json.Unmarshal([]byte(transacciones2), &data)
}

func (s *StubRepository) Write(data interface{}) error {
	return nil
}

var transacciones2 string = `[
	{"id": 1, "transaction_code": "12345", "coin": "USD", "amount": 300.00, "emitor": "Juan", "receptor": "Enrique", "transaction_date": "12/12/21"}, 
	{"id": 2, "transaction_code": "54321", "coin": "Euro", "amount": 150.00, "emitor": "Miguel", "receptor": "Luis", "transaction_date": "13/12/21"}
]`

var transactionExample = Transaction{3, "10001", "Peso", 200, "Ivan", "Geronimo", "10/10/21"}

func (s *StubRepository) GetAll() ([]Transaction, error) {
	return []Transaction{}, nil
}

func (s *StubRepository) Store(transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	return Transaction{}, nil
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
