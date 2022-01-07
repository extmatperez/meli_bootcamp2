/*
	Author: AG-Meli - Andrés Ghione
*/

package internal

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/pkg/store"
)

var transactions []domain.Transaction
var lastID int

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error)
	LastID() (int, error)
	Update(id int, code string, currency string, amount float64, remitter string, receptor string, date string) (domain.Transaction, error)
	Delete(id int) error
	ModifyTransactionCode(id int, code string) (domain.Transaction, error)
	ModifyAmount(id int, amount float64) (domain.Transaction, error)
	GetByID(id int) (domain.Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// GetByID
// @Summary Get transaction by id
// @Tags Transaction
// @Description Search a transaction by id in the database
func (repo *repository) GetByID(id int) (domain.Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for _, trans := range transactions {
		if trans.ID == id {
			return trans, nil
		}
	}
	return domain.Transaction{}, fmt.Errorf("No se encontro la transaccion con el id: %v", id)
}

// ModifyTransactionCode
// @Summary Update transaction code
// @Tags Transaction
// @Description Update the transaction code of a transaction indicated by id
func (repo *repository) ModifyTransactionCode(id int, transactionCode string) (domain.Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Code = transactionCode
			err = repo.db.Write(&transactions)
			if err != nil {
				return domain.Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

// ModifyAmount
// @Summary Update amount
// @Tags Transaction
// @Description Update the amount of a transaction indicated by id
func (repo *repository) ModifyAmount(id int, amount float64) (domain.Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Amount = amount
			err = repo.db.Write(&transactions)
			if err != nil {
				return domain.Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transactions[i], nil
		}
	}
	return domain.Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

// Delete
// @Summary Delete transaction
// @Tags Transaction
// @Description Delete the transaction with the indicated id from the database
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

// GetAll
// @Summary Get all transactions
// @Tags Transaction
// @Description Search all transactions in the database
func (repo *repository) GetAll() ([]domain.Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	return transactions, nil
}

// Store
// @Summary Create new transaction
// @Tags Transaction
// @Description Create new transaction in database
func (repo *repository) Store(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error) {
	transact := domain.Transaction{id, code, currency, amount, remitter, receptor, date}
	err := repo.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	lastID = id
	transactions = append(transactions, transact)
	err = repo.db.Write(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al escribir el store")
	}
	return transact, nil
}

// Update
// @Summary Update all fields of a transaction
// @Tags Transaction
// @Description Update all fields of a transaction
func (repo *repository) Update(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error) {
	transaction := domain.Transaction{id, code, currency, amount, remitter, receptor, date}
	err := repo.db.Read(&transactions)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("Error al leer el store")
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i] = transaction
			err = repo.db.Write(&transactions)
			if err != nil {
				return domain.Transaction{}, fmt.Errorf("Error al escribir el store")
			}
			return transaction, nil
		}
	}
	return domain.Transaction{}, fmt.Errorf("No se encontro la transaccion con id %d", id)
}

// LastID
// @Summary Get the last created id
// @Tags Transaction
// @Description Get the last created id
func (repo *repository) LastID() (int, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return 0, fmt.Errorf("Error al leer el store")
	}
	return transactions[len(transactions)-1].ID, nil
}
