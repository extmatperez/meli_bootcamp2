package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/models"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/pkg/db"
)

type RepositoryDB interface {
	Store(transaction models.Transaction) (models.Transaction, error)
	GetOne(id int) models.Transaction
	Update(transaction models.Transaction) (models.Transaction, error)
	GetBySender(sender string) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	Delete(id int) error
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
	rows, err := db.Query("SELECT id, transaction_code, currency, amount, receiver, sender, transaction_date FROM transactions WHERE id = ?", id)

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
		log.Fatal("err", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaction.TransactionCode, transaction.Currency, transaction.Amount, transaction.Receiver, transaction.Sender, transaction.TransactionDate, transaction.ID)
	if err != nil {
		fmt.Println("error:", transaction)
		return models.Transaction{}, err
	}
	fmt.Println("inside repository:", transaction)
	filasActualizadas, _ := result.RowsAffected()

	if filasActualizadas == 0 {
		fmt.Println("error:", transaction)
		return models.Transaction{}, errors.New("No se encontro la transaction")
	}
	fmt.Println("inside repository 2:", transaction)
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
		err = rows.Scan(&transactionLeida.TransactionCode, &transactionLeida.Currency, &transactionLeida.Amount, &transactionLeida.Receiver, &transactionLeida.Sender, &transactionLeida.TransactionDate)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		transactions = append(transactions, transactionLeida)
	}
	return transactions, nil
}

func (r *repositoryDB) Delete(id int) error {
	db := db.StorageDB

	stmt, err := db.Prepare("DELETE FROM transactions WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("No se encontro la transaction")
	}
	return nil

}

func (r *repositoryDB) GetOneWithContext(ctx context.Context, id int) (models.Transaction, error) {

	db := db.StorageDB
	//ctx = context.WithValue(ctx, "3", "2")
	var transactionLeida models.Transaction
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := db.QueryContext(ctx, "SELECT transaction_code, currency, amount, receiver, sender, transaction_date FROM transactions WHERE id = ?", id)

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

func (r *repositoryDB) GetFullData() ([]models.Transaction, error) {
	var myTransactions []models.Transaction
	db := db.StorageDB
	var transactionRead models.Transaction
	rows, err := db.Query("SELECT t.id ,t.sender, t.receiver, t.amount, t.currency, p.dni, p.name FROM transactions t INNER JOIN persons p ON t.id_person = p.id")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&transactionRead.ID, &transactionRead.Sender, &transactionRead.Receiver, &transactionRead.Amount, &transactionRead.Currency, &transactionRead.Person.DNI, &transactionRead.Person.Name)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		myTransactions = append(myTransactions, transactionRead)
	}
	return myTransactions, nil
}
