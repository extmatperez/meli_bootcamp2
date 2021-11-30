package internal

import "errors"

type Transaction struct {
	ID                  int     `json:"id"`
	CodigoDeTransaccion string  `json:"codigo_de_transaccion" binding:"required"`
	Moneda              string  `json:"moneda" binding:"required"`
	Monto               float64 `json:"monto" binding:"required"`
	Emisor              string  `json:"emisor" binding:"required"`
	Receptor            string  `json:"receptor" binding:"required"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion" binding:"required"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	GetTransactionByID(id int) (Transaction, error)
	Store(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error)
	LastId() int
	ExistsTransaction(id int) bool
}

type repository struct{}

var transactions []Transaction
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	return transactions, nil // TODO: Manejar errores

}

func (repo *repository) GetTransactionByID(id int) (Transaction, error) {
	for _, t := range transactions {
		if t.ID == id {
			return t, nil
		}
	}
	return Transaction{}, errors.New("transaction not found")
}

func (repo *repository) Store(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {
	transac := Transaction{id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion}
	lastID = id
	transactions = append(transactions, transac)
	return transac, nil // TODO: Manejar errores
}

func (repo *repository) LastId() int {
	return lastID
}

func (repo *repository) ExistsTransaction(id int) bool {
	for _, t := range transactions {
		if t.ID == id {
			return true
		}
	}
	return false
}
