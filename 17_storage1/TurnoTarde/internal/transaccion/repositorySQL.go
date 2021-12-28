package internal

import (
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/17_storage1/TurnoTarde/db"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/17_storage1/TurnoTarde/internal/transaccion/models"
)

type RepositorySql interface {
	Store(transaction models.Transaction) (models.Transaction, error)
	GetById(id int) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	Delete(id int) (error)
}

type repositorySQL struct {
}

func NewRepositorySQL() RepositorySql {
	return &repositorySQL{}
}

const (
	InsertOne = "INSERT INTO transaction(Codigo,Moneda,Monto,Emisor,Receptor,Fecha)" +
		"VALUES(?, ?, ?, ?, ?, ?)"
	GetById = "SELECT Id,Codigo,Moneda,Monto,Emisor,Receptor,Fecha FROM transaction WHERE Id=?"
	GetAll  = "SELECT Id,Codigo,Moneda,Monto,Emisor,Receptor,Fecha FROM transaction"
	Delete = "DELETE FROM transaction WHERE Id=?"
)

func (r *repositorySQL) Store(transaction models.Transaction) (models.Transaction, error) {
	db := db.StorageDB                 // se inicializa la base
	stmt, err := db.Prepare(InsertOne) // se prepara el SQL
	if err != nil {
		return models.Transaction{},err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(transaction.Codigo, transaction.Moneda, transaction.Monto, transaction.Emisor,
		transaction.Receptor, transaction.Fecha) // retorna un sql.Result y un error
	if err != nil {
		return models.Transaction{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	transaction.ID = int(insertedId)
	return transaction, nil

}

func (r *repositorySQL) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	db := db.StorageDB
	rows, err := db.Query(GetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction models.Transaction

		err := rows.Scan(&transaction.ID, &transaction.Codigo, &transaction.Moneda, &transaction.Monto, &transaction.Emisor,
			&transaction.Receptor, &transaction.Fecha)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}
	return transactions, nil

}

func (r *repositorySQL) GetById(id int) (models.Transaction, error) {
	var transaction models.Transaction
	db := db.StorageDB
	rows, err := db.Query(GetById, id)
	if err != nil {
		log.Println(err)
		return transaction, err
	}
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.Codigo, &transaction.Moneda, &transaction.Monto, &transaction.Emisor,
			&transaction.Receptor, &transaction.Fecha); err != nil {
			log.Println(err.Error())
			return transaction, err
		}
	}
	return transaction, nil

}
func (r *repositorySQL) Delete(id int) (error) {
	db := db.StorageDB
	stmt, err := db.Prepare(Delete) // se prepara el SQL
	if err != nil {
		return err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(id) // retorna un sql.Result y un error
	if err != nil {
		return err
	}
	filasAfectadas,_ :=result.RowsAffected() 
	if(filasAfectadas== 0){
		return errors.New("No se encontro el id")
	}

	return  nil

}
