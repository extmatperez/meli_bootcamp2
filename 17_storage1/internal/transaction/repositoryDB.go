package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"
)

type RepositoryDB interface {
	Store(ntransaction models.Transaction) (models.Transaction, error)
	GetOne(id int) models.Transaction
	Update(transaction models.Transaction) (models.Transaction, error)
}

type repositoryDB struct{}

func NewRepositoryDB() RepositoryDB {
	return &repositoryDB{}
}

func (r *repositoryDB) Store(transaction models.Transaction) (models.Transaction, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO transactions(nombre, apellido, edad) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(transaction.Nombre, transaction.Apellido, transaction.Edad)
	if err != nil {
		return models.Transaction{}, err
	}
	idCreado, _ := result.LastInsertId()
	transaction.ID = int(idCreado)

	return transaction, nil
}

func (r *repositoryDB) GetOne(id int) models.Transaction {
	db := db.StorageDB
	var transactionLeida models.Transaction
	rows, err := db.Query("SELECT id, nombre,apellido, edad FROM transactions WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return transactionLeida
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.Nombre, &transactionLeida.Apellido, &transactionLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return transactionLeida
		}
	}
	return transactionLeida
}

func (r *repositoryDB) Update(transaction models.Transaction) (models.Transaction, error) {

	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE transactions SET nombre = ?, apellido = ?, edad = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaction.Nombre, transaction.Apellido, transaction.Edad, transaction.ID)
	if err != nil {
		return models.Transaction{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Transaction{}, errors.New("No se encontro la transaction")
	}

	return transaction, nil
}
