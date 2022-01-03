package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/models"
)

type RepositoryDBMock interface {
	Store(transaction models.Transaction) (models.Transaction, error)
	GetOne(id int) models.Transaction
	Update(transaction models.Transaction) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	Delete(id int) error
	GetFullData() ([]models.Transaction, error)

	GetOneWithContext(ctx context.Context, id int) (models.Transaction, error)
}

type repositoryDBMock struct {
	db *sql.DB
}

func NewRepositoryDBMock(db *sql.DB) RepositoryDBMock {
	return &repositoryDBMock{db}
}

func (r *repositoryDBMock) Store(transaction models.Transaction) (models.Transaction, error) {

	stmt, err := r.db.Prepare("INSERT INTO transactions(transaction_code, currency, amount, receiver, sender, transaction_date) VALUES(?, ?, ?, ?, ?, ?)")
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

func (r *repositoryDBMock) GetOne(id int) models.Transaction {

	var transactionLeida models.Transaction
	rows, err := r.db.Query("SELECT id, sender, receiver, amount FROM transactions WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return transactionLeida
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.Sender, &transactionLeida.Receiver, &transactionLeida.Amount)
		if err != nil {
			log.Fatal(err)
			return transactionLeida
		}

	}
	return transactionLeida
}
func (r *repositoryDBMock) GetAll() ([]models.Transaction, error) {
	var misTransactions []models.Transaction

	var transactionLeida models.Transaction
	rows, err := r.db.Query("SELECT id, sender, receiver, amount FROM transactions")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.Sender, &transactionLeida.Receiver, &transactionLeida.Amount)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misTransactions = append(misTransactions, transactionLeida)
	}
	return misTransactions, nil
}

func (r *repositoryDBMock) Update(transaction models.Transaction) (models.Transaction, error) {

	stmt, err := r.db.Prepare("UPDATE transactions SET sender = ?, receiver = ?, amount = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaction.Sender, transaction.Receiver, transaction.Amount, transaction.ID)
	if err != nil {
		return models.Transaction{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Transaction{}, errors.New("No se encontro la transaction")
	}

	return transaction, nil
}

func (r *repositoryDBMock) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM transactions WHERE id = ?")
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

func (r *repositoryDBMock) GetFullData() ([]models.Transaction, error) {
	var myTransactions []models.Transaction

	var transactionRead models.Transaction
	rows, err := r.db.Query("SELECT t.id ,t.sender, t.receiver, t.amount, t.currency, p.dni, p.name FROM transactions t INNER JOIN persons p ON t.id_person = p.id")

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

func (r *repositoryDBMock) GetOneWithContext(ctx context.Context, id int) (models.Transaction, error) {

	var transactionLeida models.Transaction
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := r.db.QueryContext(ctx, "SELECT id, sender, receiver, amount FROM transactions WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return transactionLeida, err
	}

	for rows.Next() {
		err = rows.Scan(&transactionLeida.ID, &transactionLeida.Sender, &transactionLeida.Receiver, &transactionLeida.Amount)
		if err != nil {
			log.Fatal(err)
			return transactionLeida, err
		}

	}
	return transactionLeida, nil
}
