package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/pkg/store"
)

type Transaction struct {
	ID               int     `form:"id", json:"id"`
	Transaction_Code string  `form:"transaction_code", json:"transaction_code"`
	Coin             string  `form:"coin", json:"coin"`
	Amount           float64 `form:"amount", json:"amount"`
	Emitor           string  `form:"emitor", json:"emitor"`
	Receptor         string  `form:"receptor", json:"receptor"`
	Transaction_Date string  `form:"transaction_date", json:"transaction_date"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error)
	Update(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error)
	UpdateReceptor(id int, receptor string) (Transaction, error)
	Delete(id int) error
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	err := repo.db.Read(transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repo *repository) Store(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	var transactions []Transaction
	repo.db.Read(transactions)

	trans := Transaction{id, transaction_code, coin, amount, receptor, transaction_date, transaction_date}
	transactions = append(transactions, trans)

	err := repo.db.Write(transactions)

	if err != nil {
		return Transaction{}, err
	}

	return trans, nil
}

func (repo *repository) LastId() (int, error) {
	var transactions []Transaction
	err := repo.db.Read(transactions)

	if err != nil {
		return 0, err
	}

	if len(transactions) == 0 {
		return 0, nil
	}
	return transactions[len(transactions)-1].ID, nil
}

func (repo *repository) Update(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	var transactions []Transaction
	err := repo.db.Read(transactions)

	if err != nil {
		return Transaction{}, err
	}

	trans := Transaction{id, transaction_code, coin, amount, receptor, transaction_date, transaction_date}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i] = trans
			err := repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return trans, nil
		}
	}
	return Transaction{}, fmt.Errorf("La transacción %d no existe", id)

}

func (repo *repository) UpdateReceptor(id int, receptor string) (Transaction, error) {
	var transactions []Transaction
	err := repo.db.Read(transactions)

	if err != nil {
		return Transaction{}, err
	}
	for i, v := range transactions {
		if v.ID == id {
			transactions[i].Receptor = receptor
			err := repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("La transacción %d no existe", id)

}

func (repo *repository) Delete(id int) error {
	var transactions []Transaction
	err := repo.db.Read(transactions)

	if err != nil {
		return err
	}

	index := 0
	for i, v := range transactions {
		if v.ID == id {
			index = i
			transactions = append(transactions[:index], transactions[index+1:]...)
			err := repo.db.Write(transactions)
			return err
		}
	}
	return fmt.Errorf("La transacción %d no existe", id)

}
