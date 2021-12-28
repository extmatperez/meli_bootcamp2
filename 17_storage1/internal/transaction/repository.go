package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/pkg/store"
)

type Transaction struct {
	ID              int     `form:"id" json:"id"`
	TransactionCode string  `form:"transaction_code" json:"transaction_code" validate:"required,transaction_code"`
	Currency        string  `form:"currency" json:"currency" validate:"required,currency"`
	Amount          float64 `form:"amount" json:"amount" validate:"required,amount"`
	Receiver        string  `form:"receiver" json:"receiver" validate:"required,receiver"`
	Sender          string  `form:"sender" json:"sender" validate:"required,sender"`
	TransactionDate string  `form:"transaction_date" json:"transaction_date" validate:"required,transaction_date"`
}

var transactions []Transaction

type Repository interface {
	LastId() (int, error)
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	GetByReceiver(receiver string) (Transaction, error)
	CreateTransaction(transaction Transaction) (Transaction, error)
	Store(id int, transactionCode string, currency string, amount float64,
		receiver string, sender string, transactionDate string) (Transaction, error)
	UpdateTransaction(id int, transactionCode string, currency string, amount float64,
		receiver string, sender string, transactionDate string) (Transaction, error)
	UpdateAmount(id int, amount float64) (Transaction, error)
	DeleteTransaction(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) LastId() (int, error) {
	_, err := repo.db.Read(&transactions)
	if err != nil {
		return 0, err
	}
	return transactions[len(transactions)-1:][0].ID, nil
}

func (repo *repository) GetAll() ([]Transaction, error) {
	_, err := repo.db.Read(&transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repo *repository) GetByID(idParam int) (Transaction, error) {
	_, err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	var transaction Transaction
	for i, trans := range transactions {
		if idParam == trans.ID {
			transaction = transactions[i]
			break
		}
	}
	return transaction, nil
}

func (repo *repository) GetByReceiver(receiver string) (Transaction, error) {
	_, err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	var transaction Transaction
	for i, trans := range transactions {
		if receiver == trans.Receiver {
			transaction = transactions[i]
			break
		}
	}
	return transaction, nil
}

func (repo *repository) Store(id int, transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {

	isExists, err := repo.db.Read(&transactions)

	if !isExists {
		_, err := repo.db.Write(&transactions)

		if err != nil {
			return Transaction{}, err
		}
		_, err = repo.db.Read(&transactions)

		if err != nil {
			return Transaction{}, err
		}
		tran := Transaction{id, transactionCode, currency, amount, receiver, sender, transactionDate}
		transactions = append(transactions, tran)
		_, err = repo.db.Write(&transactions)
		if err != nil {
			return Transaction{}, err
		}

		return tran, nil
	} else {
		if err != nil {
			return Transaction{}, err
		}
		tran := Transaction{id, transactionCode, currency, amount, receiver, sender, transactionDate}
		transactions = append(transactions, tran)
		_, err = repo.db.Write(transactions)
		if err != nil {
			return Transaction{}, err
		}

		return tran, nil
	}
}

func (repo *repository) CreateTransaction(trans Transaction) (Transaction, error) {
	_, err := repo.db.Read(&transactions)

	if err != nil {
		return Transaction{}, err
	}
	lastId, err := repo.LastId()
	if err != nil {
		panic(err)
	}
	trans.ID = lastId + 1

	transactions = append(transactions, trans)

	_, err = repo.db.Write(&transactions)
	if err != nil {
		return Transaction{}, err
	}

	return trans, nil

}

func (repo *repository) UpdateTransaction(id int, transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {
	_, err := repo.db.Read(&transactions)

	if err != nil {
		return Transaction{}, err
	}

	var transaction Transaction
	transaction.TransactionCode = transactionCode
	transaction.Currency = currency
	transaction.Currency = currency
	transaction.Amount = amount
	transaction.Receiver = receiver
	transaction.Sender = sender
	transaction.TransactionDate = transactionDate
	for i, v := range transactions {
		if v.ID == id {
			transaction.ID = i + 1
			transactions[i] = transaction
			_, err = repo.db.Write(&transactions)
			if err != nil {
				return Transaction{}, err
			}
			return transaction, nil
		}
	}

	return Transaction{}, nil

}

func (repo *repository) UpdateAmount(id int, amount float64) (Transaction, error) {
	_, err := repo.db.Read(&transactions)

	if err != nil {
		return Transaction{}, err
	}

	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Amount = amount
			_, err = repo.db.Write(&transactions)
			if err != nil {
				return Transaction{}, err
			}
			return transactions[i], nil
		}
	}

	return Transaction{}, nil

}

func (repo *repository) DeleteTransaction(id int) error {
	_, err := repo.db.Read(&transactions)

	if err != nil {
		return err
	}
	index := 0
	for i, v := range transactions {
		if v.ID == id {
			index = i
			transactions = append(transactions[:index], transactions[index+1:]...)
			_, err = repo.db.Write(&transactions)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("The transaction %d no exists", id)
}
