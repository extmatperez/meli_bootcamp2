package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/pkg/db"
)

type RepositoryDB interface {
	Store(transaction models.Transaction) (models.Transaction, error)
	GetOne(id int) models.Transaction
	Update(transaction models.Transaction) (models.Transaction, error)
	GetBySender(sender string) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
}

type repositoryDB struct{}

func NewRepositoryDB() RepositoryDB {
	return &repositoryDB{}
}

func (r *repositoryDB) Store(transaction models.Transaction) (models.Transaction, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO transactions(transaction_code, currency, amount, receiver, sender, transaction_date) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(transaction.TransactionCode, transaction.Currency, transaction.Amount, transaction.Receiver, transaction.Sender, transaction.TransactionDate)
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
	rows, err := db.Query("SELECT transaction_code, currency, amount, receiver, sender, transaction_date FROM transactions WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return transactionLeida
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.TransactionCode, &transactionLeida.Currency, &transactionLeida.Amount, &transactionLeida.Receiver, &transactionLeida.Sender, &transactionLeida.TransactionDate)
		if err != nil {
			log.Fatal(err)
			return transactionLeida
		}
	}
	return transactionLeida
}

func (r *repositoryDB) Update(transaction models.Transaction) (models.Transaction, error) {

	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE transactions SET transaction_code = ?, currency = ?, amount = ?, receiver = ?, sender = ?, transaction_date = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaction.TransactionCode, transaction.Currency, transaction.Amount, transaction.Receiver, transaction.Sender, transaction.TransactionDate, transaction.ID)
	if err != nil {
		return models.Transaction{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Transaction{}, errors.New("No se encontro la transaction")
	}

	return transaction, nil
}

func (r *repositoryDB) GetBySender(sender string) (models.Transaction, error) {
	db := db.StorageDB
	var transactionLeida models.Transaction
	rows, err := db.Query("SELECT transaction_code, currency, amount, receiver, sender, transaction_date FROM transactions WHERE sender = ?", sender)

	if err != nil {
		log.Fatal(err)
		return transactionLeida, err
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.TransactionCode, &transactionLeida.Currency, &transactionLeida.Amount, &transactionLeida.Receiver, &transactionLeida.Sender, &transactionLeida.TransactionDate)
		if err != nil {
			log.Fatal(err)
			return transactionLeida, err
		}
	}
	return transactionLeida, nil
}

func (r *repositoryDB) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	db := db.StorageDB
	var transactionLeida models.Transaction
	rows, err := db.Query("SELECT transaction_code, currency, amount, receiver, sender, transaction_date FROM transactions")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.TransactionCode, &transactionLeida.Currency, &transactionLeida.Amount, &transactionLeida.Receiver, &transactionLeida.Sender, &transactionLeida.TransactionDate)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		transactions = append(transactions, transactionLeida)
	}
	return transactions, nil
}
