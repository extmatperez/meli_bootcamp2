package internal

import "fmt"

type Transaction struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Date     string `json:"date"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error)
	Update(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (repo *repository) Store(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error) {
	transaction := Transaction{id, code, currency, amount, sender, receiver, date}
	lastID = id
	transactions = append(transactions, transaction)
	return transaction, nil
}

func (repo *repository) Update(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error) {

	for i, transaction := range transactions {
		if transaction.Id == id {
			transactions[i].Code = code
			transactions[i].Currency = currency
			transactions[i].Amount = amount
			transactions[i].Sender = sender
			transactions[i].Receiver = receiver
			transactions[i].Date = date
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("Transaction with id: %v not found", id)
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
