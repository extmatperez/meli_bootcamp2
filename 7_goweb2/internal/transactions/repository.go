package internal

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
}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
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

func (repo *repository) LastID() (int, error){
	return lastID, nil
}