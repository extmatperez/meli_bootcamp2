/*
	Author: AG-Meli - Andrés Ghione
*/

package internal

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error)
	Update(id int, code string, currency string, amount float64, remitter string, receptor string, date string) (domain.Transaction, error)
	Delete(id int) error
	ModifyTransactionCode(id int, code string) (domain.Transaction, error)
	ModifyAmount(id int, amount float64) (domain.Transaction, error)
	GetByID(id int) (domain.Transaction, error)
}

type service struct {
	repository Repository
}

// GetByID
// @Summary Get transaction by id
// @Tags Transaction
// @Description Search a transaction by id in the repository
func (ser *service) GetByID(id int) (domain.Transaction, error) {
	transactions, err := ser.repository.GetByID(id)
	if err != nil {
		return domain.Transaction{}, err
	} else {
		return transactions, nil
	}
}

// ModifyTransactionCode
// @Summary Update transaction code
// @Tags Transaction
// @Description Update the transaction code of a transaction indicated by id
func (ser *service) ModifyAmount(id int, amount float64) (domain.Transaction, error) {
	if amount > 0 {
		transaction, err := ser.repository.ModifyAmount(id, amount)
		if err != nil {
			return domain.Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return domain.Transaction{}, fmt.Errorf("El codigo de transaccion no puede estar vacio")
	}
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

// ModifyAmount
// @Summary Update amount
// @Tags Transaction
// @Description Update the amount of a transaction indicated by id
func (ser *service) ModifyTransactionCode(id int, transactionCode string) (domain.Transaction, error) {
	if transactionCode != "" {
		transaction, err := ser.repository.ModifyTransactionCode(id, transactionCode)
		if err != nil {
			return domain.Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return domain.Transaction{}, fmt.Errorf("El codigo de transaccion no puede estar vacio")
	}
}

// GetAll
// @Summary Get all transactions
// @Tags Transaction
// @Description Search all transactions in the repository
func (ser *service) GetAll() ([]domain.Transaction, error) {
	transactions, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	} else {
		return transactions, nil
	}
}

// Store
// @Summary Create new transaction
// @Tags Transaction
// @Description Create new transaction in repository
func (ser *service) Store(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error) {
	newID, err := ser.repository.LastID()
	if err != nil {
		return domain.Transaction{}, err
	} else {
		transaction, err := ser.repository.Store(newID+1, code, currency, amount, remitter, receptor, date)
		if err != nil {
			return domain.Transaction{}, err
		} else {
			return transaction, nil
		}
	}
}

// Update
// @Summary Update all fields of a transaction
// @Tags Transaction
// @Description Update all fields of a transaction
func (ser *service) Update(id int, code, currency string, amount float64, remitter, receptor, date string) (domain.Transaction, error) {
	if validateFields(id, code, currency, amount, remitter, receptor, date) {
		transaction, err := ser.repository.Update(id, code, currency, amount, remitter, receptor, date)
		if err != nil {
			return domain.Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return domain.Transaction{}, fmt.Errorf("Alguno de los campos son incorrectos")
	}
}

// Delete
// @Summary Delete transaction
// @Tags Transaction
// @Description Delete the transaction with the indicated id from the repository
func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}

// validateFields
// @Summary Validates the values sent by parameters
// @Tags Transaction
// @Description Validates the values sent by parameters
func validateFields(id int, code string, currency string, amount float64, remitter string, receptor string, date string) bool {
	if id <= 0 || amount <= 0 {
		return false
	}
	if code == "" || currency == "" || remitter == "" || receptor == "" || date == "" {
		return false
	}
	return true
}
