package models

type Transaction struct {
	ID              int     `form:"id" json:"id"`
	TransactionCode string  `form:"transaction_code" json:"transaction_code" validate:"required,transaction_code"`
	Currency        string  `form:"currency" json:"currency" validate:"required,currency"`
	Amount          float64 `form:"amount" json:"amount" validate:"required,amount"`
	Receiver        string  `form:"receiver" json:"receiver" validate:"required,receiver"`
	Sender          string  `form:"sender" json:"sender" validate:"required,sender"`
	TransactionDate string  `form:"transaction_date" json:"transaction_date" validate:"required,transaction_date"`
}
