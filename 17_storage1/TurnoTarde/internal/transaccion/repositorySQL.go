package internal

import (
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/17_storage1/TurnoTarde/db"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/17_storage1/TurnoTarde/internal/transaccion/models"
)

type RepositorySql interface{
	Store(transaction models.Transaction) (models.Transaction,error)
	GetByCodigo(codigo string) (models.Transaction,error)

}


type repositorySQL struct{

}

func NewRepositorySQL() RepositorySql{
	return &repositorySQL{}
}

const (
	InsertOne = "INSERT INTO transaction(Codigo,Moneda,Monto,Emisor,Receptor,Fecha)" +
										"VALUES(?, ?, ?, ?, ?, ?)"
	GetByCodigo = "SELECT (Id, Codigo,Moneda,Monto,Emisor,Receptor,Fecha) FROM transaction WHERE Codigo=?"
)



func (r *repositorySQL) Store(transaction models.Transaction) (models.Transaction,error){
	db := db.StorageDB // se inicializa la base
	stmt, err := db.Prepare(InsertOne) // se prepara el SQL
	if err != nil {
	log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(transaction.Codigo,transaction.Moneda,transaction.Monto,transaction.Emisor,
		transaction.Receptor,transaction.Fecha) // retorna un sql.Result y un error
	if err != nil {
	return models.Transaction{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	transaction.ID = int(insertedId)
	return transaction, nil

}

func (r *repositorySQL) GetByCodigo(codigo string) (models.Transaction,error){
	var transaction models.Transaction
	db := db.StorageDB
	rows, err := db.Query(GetByCodigo, codigo)
	if err != nil {
	log.Println(err)
	return transaction,err
	}
	for rows.Next() {
	if err := rows.Scan(&transaction.Codigo,&transaction.Moneda,&transaction.Monto,&transaction.Emisor,
		&transaction.Receptor,&transaction.Fecha); err != nil {
	log.Println(err.Error())
	return transaction,err
	}
	}
	return transaction,nil
   
}




