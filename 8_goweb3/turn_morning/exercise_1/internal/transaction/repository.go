package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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

func readTransactions() []Transaction {
	transacionFile := "../../../transaction.json"
	data, err := os.ReadFile(transacionFile)

	if err != nil {
		fmt.Printf("There was a error %v", err)
	}

	return toDeserializer(data)
}

func toDeserializer(data []byte) []Transaction {
	var transactions []Transaction

	if err := json.Unmarshal(data, &transactions); err != nil {
		fmt.Printf("There was a error during deserializer %v", err)
	}

	return transactions
}

func toSaveTransaction(tran *[]Transaction) {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(transactions)
	data := reqBodyBytes.Bytes()
	//data := []byte(fmt.Sprintf("%v", transactions))

	err := os.WriteFile("../../../transaction.json", data, 0644)

	if err != nil {
		panic(err)
	}
}

type Repository interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	GetByReceiver(receiver string) (Transaction, error)
	//CreateTransaction(transaction Transaction) (Transaction, error)
	Store(transactionCode string, currency string, amount float64,
		receiver string, sender string, transactionDate string) (Transaction, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	transactions = readTransactions()
	return transactions, nil
}

func (repo *repository) GetByID(idParam int) (Transaction, error) {
	transactions = readTransactions()
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
	transactions = readTransactions()
	var transaction Transaction
	for i, trans := range transactions {
		if receiver == trans.Receiver {
			transaction = transactions[i]
			break
		}
	}
	return transaction, nil
}

func (repo *repository) Store(transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {

	transactions = readTransactions()
	lastId, err := LastId()
	if err != nil {
		panic(err)
	}
	tran := Transaction{lastId + 1, transactionCode, currency, amount, receiver, sender, transactionDate}

	transactions = append(transactions, tran)

	toSaveTransaction(&transactions)

	return tran, nil
}

func LastId() (int, error) {
	transactions = readTransactions()
	return transactions[len(transactions)-1:][0].ID, nil
}
