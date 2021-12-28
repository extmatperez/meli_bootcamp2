package internal

import (
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/13_sql1/proyecto/pkg/store"
)

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
	Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error)
	UpdateCodigoYMonto(id int, codigo_de_transaccion string, monto float64) (Transaction, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

var transactions []Transaction

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return nil, errors.New("No se pudo leer el archivo")
	}
	return transactions, nil

}

func (repo *repository) GetTransactionByID(id int) (Transaction, error) {
	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, errors.New("No se pudo leer el archivo")
	}
	for _, t := range transactions {
		if t.ID == id {
			return t, nil
		}
	}
	return Transaction{}, errors.New("transaction not found")
}

func (repo *repository) Store(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {

	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	transac := Transaction{id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion}
	transactions = append(transactions, transac)
	err = repo.db.Write(transactions)
	if err != nil {
		return Transaction{}, errors.New("No se pudo escribir el archivo")
	}
	return transac, nil
}

func (repo *repository) Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {

	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, errors.New("No se pudo leer el archivo")
	}
	transac := Transaction{id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion}

	for i, t := range transactions {
		if t.ID == id {
			transactions[i] = transac

			err = repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, errors.New("No se pudo escribir el archivo")
			}

			return transac, nil
		}
	}
	return Transaction{}, fmt.Errorf("transaction %d doesn't exist", id)
}

func (repo *repository) UpdateCodigoYMonto(id int, codigo_de_transaccion string, monto float64) (Transaction, error) {

	err := repo.db.Read(&transactions)
	if err != nil {
		return Transaction{}, errors.New("No se pudo leer el archivo")
	}
	for i, t := range transactions {
		if t.ID == id {
			transactions[i].CodigoDeTransaccion = codigo_de_transaccion
			transactions[i].Monto = monto

			err = repo.db.Write(transactions)
			if err != nil {
				return Transaction{}, errors.New("No se pudo escribir el archivo")
			}

			return transactions[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("transaction %d doesn't exist", id)
}

func (repo *repository) Delete(id int) error {

	err := repo.db.Read(&transactions)
	if err != nil {
		return errors.New("No se pudo leer el archivo")
	}
	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)

			err = repo.db.Write(transactions)
			if err != nil {
				return errors.New("No se pudo escribir el archivo")
			}
			return nil
		}
	}
	return fmt.Errorf("transaction %d doesn't exist", id)

}

//
func (repo *repository) LastId() (int, error) {

	err := repo.db.Read(&transactions)
	if err != nil {
		return 0, errors.New("No se pudo leer el archivo")
	}
	if len(transactions) == 0 {
		return 0, nil
	}

	return transactions[len(transactions)-1].ID, nil
}
