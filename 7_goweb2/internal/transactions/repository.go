package internal

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/pkg/store"
)

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
	Update(id int, code string, currency string, amount float64, remitter string, receptor string, date string) (Transaction, error)
	Delete(id int) error
	ModifyTransactionCode(id int, code string) (Transaction, error)
	ModifyAmount(id int, amount float64) (Transaction, error)
	GetByID(id int) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetByID(id int) (Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for _, trans := range transactions {
		if trans.ID == id {
			return trans, nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con el id: %v", id)
}

func (repo *repository) ModifyTransactionCode(id int, transactionCode string) (Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Code = transactionCode
			err = repo.db.Write(&transactions)
			if err != nil {
				return Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) ModifyAmount(id int, amount float64) (Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Amount = amount
			err = repo.db.Write(&transactions)
			if err != nil {
				return Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&transactions)
	if err != nil {
		return fmt.Errorf("Error al leer el store")
	}
	for k, v := range transactions {
		if v.ID == id {
			transactions = append(transactions[:k], transactions[k+1:]...)
			err = repo.db.Write(&transactions)
			if err != nil {
				return fmt.Errorf("Error al escribir el store")
			}
			return nil
		}
	}
	return fmt.Errorf("No se ha encontrado la transaccion con id %v", id)
}

func (repo *repository) GetAll() ([]Transaction, error){
	err := repo.db.Read(&transactions)
	if err != nil {
		return []Transaction{}, fmt.Errorf("Error al leer el store")
	}
	return transactions, nil
}

func (repo *repository) Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	transact := Transaction{id, code, currency, amount, remitter, receptor, date}
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al leer el store")
	}
	lastID = id
	transactions = append(transactions, transact)
	err = repo.db.Write(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al escribir el store")
	}
	return transact, nil
}

func (repo *repository) Update(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	transaction := Transaction{id, code, currency, amount, remitter, receptor, date}
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i] = transaction
			err = repo.db.Write(&transactions)
			if err != nil {
				return Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transaction, nil
		}
	}
	return Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

func (repo *repository) LastID() (int, error){
	err := repo.db.Read(&transactions)
	if err != nil {
		return 0, fmt.Errorf("Error al leer el store")
	}
	return transactions[len(transactions)-1].ID, nil
}