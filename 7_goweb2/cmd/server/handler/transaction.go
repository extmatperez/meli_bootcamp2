package handler

import (
	transactions "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/transactions"
)

type request struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Remitter string  `json:"remitter"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(ser transactions.Service) *Transaction {
	return &Transaction{service: ser}
}

func (transact *Transaction) GetAll() {

}