package internal

type Transaction struct {
	ID               int     `form:"id", json:"id"`
	Transaction_Code string  `form:"transaction_code", json:"transaction_code"`
	Coin             string  `form:"coin", json:"coin"`
	Amount           float64 `form:"amount", json:"amount"`
	Emitor           string  `form:"emitor", json:"emitor"`
	Receptor         string  `form:"receptor", json:"receptor"`
	Transaction_Date string  `form:"transaction_date", json:"transaction_date"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error)
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (repo *repository) Store(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	trans := Transaction{id, transaction_code, coin, amount, receptor, transaction_date, transaction_date}
	lastID = id
	transactions = append(transactions, trans)
	return trans, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
