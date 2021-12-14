package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/pkg/store"
)

type Transaction struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Date     string `json:"date"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error)
	Update(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error)
	UpdateCodeAndAmount(id int, code string, amount int) (Transaction, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	transactions := []Transaction{}
	err := repo.db.Read(&transactions)
	return transactions, err
}

func (repo *repository) Store(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error) {

	transactions := []Transaction{}
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}

	newTransaction := Transaction{id, code, currency, amount, sender, receiver, date}
	transactions = append(transactions, newTransaction)

	err = repo.db.Write(transactions)
	if err != nil {
		return Transaction{}, err
	}
	return newTransaction, nil
}

func (repo *repository) Update(id int, code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error) {

	newTransaction := Transaction{id, code, currency, amount, sender, receiver, date}

	transactions := []Transaction{}
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	for i, transaction := range transactions {
		if transaction.Id == id {
			transactions[i] = newTransaction

			err := repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return newTransaction, err
		}
	}
	return Transaction{}, fmt.Errorf("Transaction  not found")
}

func (repo *repository) UpdateCodeAndAmount(id int, code string, amount int) (Transaction, error) {

	transactions := []Transaction{}

	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	for i, transaction := range transactions {
		if transaction.Id == id {
			transactions[i].Code = code
			transactions[i].Amount = amount
			err := repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return transactions[i], err
		}
	}
	return Transaction{}, fmt.Errorf("Transaction  not found")
}

func (repo *repository) Delete(id int) error {
	transactions := []Transaction{}

	err := repo.db.Read(&transactions)
	if err != nil {
		return err
	}
	for i, transaction := range transactions {
		if transaction.Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			err := repo.db.Write(transactions)
			return err
		}
	}

	return fmt.Errorf("Transaction with id %v not found", id)
}

func (repo *repository) LastId() (int, error) {

	transactions := []Transaction{}
	err := repo.db.Read(&transactions)

	if err != nil {
		return 0, err
	}

	if len(transactions) == 0 {
		return 0, nil
	}

	return transactions[len(transactions)-1].Id, nil
}
