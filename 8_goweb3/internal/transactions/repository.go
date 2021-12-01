package internal

import (
	"fmt"
)

type Transaction struct {
	ID       int   `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Remitter string  `json:"remitter"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error)
	LastID() (int, error)
	Update(id int, code string, currency string, amount float64, remitter string, receptor string, date string) (Transaction, error)
	Delete(id int) error
	ModifyTransactionCode(id int, code string) (Transaction, error)
	ModifyAmount(id int, amount float64) (Transaction, error)
}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) ModifyTransactionCode(id int, transactionCode string) (Transaction, error) {
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Code = transactionCode
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) ModifyAmount(id int, amount float64) (Transaction, error) {
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Amount = amount
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) Delete(id int) error {
	for k, v := range transactions {
		if v.ID == id {
			transactions = append(transactions[:k], transactions[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No se ha encontrado la transaccion con id %v", id)
}

func (repo *repository) GetAll() ([]Transaction, error){
	return transactions, nil
}

func (repo *repository) Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	transact := Transaction{id, code, currency, amount, remitter, receptor, date}
	lastID = id
	transactions = append(transactions, transact)
	return transact, nil
}

func (repo *repository) Update(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	transaction := Transaction{id, code, currency, amount, remitter, receptor, date}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i] = transaction
			return transaction, nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) LastID() (int, error){
	return lastID, nil
}