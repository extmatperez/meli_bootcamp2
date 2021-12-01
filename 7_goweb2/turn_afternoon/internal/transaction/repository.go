package internal

type Transaction struct {
	ID              int     `form:"id" json:"id"`
	TransactionCode string  `form:"transaction_code" json:"transaction_code"`
	Currency        string  `form:"currency" json:"currency"`
	Amount          float64 `form:"amount" json:"amount"`
	Receiver        string  `form:"receiver" json:"receiver"`
	Sender          string  `form:"sender" json:"sender"`
	TransactionDate string  `form:"transaction_date" json:"transaction_date"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	//GetByReceiver(receiver string) ([]Transaction, error)
	//CreateTransaction(transaction Transaction) (Transaction, error)
	Store(ID int, transactionCode string, currency string, amount float64,
		receiver string, sender string, transactionDate string) (Transaction, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (repo *repository) GetByID(id int) (Transaction, error) {
	return transactions[id], nil
}

func (repo *repository) Store(id int, transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {

	tran := Transaction{id, transactionCode, currency, amount, receiver, sender, transactionDate}
	lastID = id
	transactions = append(transactions, tran)

	return tran, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
