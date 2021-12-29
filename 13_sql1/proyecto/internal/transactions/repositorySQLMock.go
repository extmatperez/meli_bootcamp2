package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/13_sql1/proyecto/internal/models"
)

type RepositorySQLMock interface {
	GetAll() ([]models.Transaction, error)
	GetTransactionByID(id int) (models.Transaction, error)
	Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	UpdateWithContext(ctx context.Context, id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	//UpdateCodigoYMonto(id int, codigo_de_transaccion string, monto float64) (models.Transaction, error)
	Delete(id int) error
	//LastId() (int, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQL {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	transaccion := models.Transaction{
		CodigoDeTransaccion: codigo_de_transaccion,
		Moneda:              moneda,
		Monto:               monto,
		Emisor:              emisor,
		Receptor:            receptor,
		FechaDeTransaccion:  fecha_de_transaccion,
	}

	stmt, err := r.db.Prepare("INSERT INTO transactions(codigo_de_transaccion, moneda, monto, emisor, receptor,fecha_de_transaccion) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)

	if err != nil {
		return models.Transaction{}, err
	}
	idCreado, _ := result.LastInsertId()
	transaccion.ID = int(idCreado)
	return transaccion, nil
}

func (r *repositorySQLMock) GetTransactionByID(id int) (models.Transaction, error) {
	var transaccLeida models.Transaction
	rows, err := r.db.Query("SELECT id,codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion FROM transactions WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return transaccLeida, err
	}
	for rows.Next() {
		err = rows.Scan(&transaccLeida.ID, &transaccLeida.CodigoDeTransaccion, &transaccLeida.Moneda, &transaccLeida.Monto, &transaccLeida.Emisor, &transaccLeida.Receptor, &transaccLeida.FechaDeTransaccion)
		if err != nil {
			log.Fatal(err)
			return models.Transaction{}, err
		}
	}
	return transaccLeida, nil
}

func (r *repositorySQLMock) Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	transaccion := models.Transaction{
		ID:                  id,
		CodigoDeTransaccion: codigo_de_transaccion,
		Moneda:              moneda,
		Monto:               monto,
		Emisor:              emisor,
		Receptor:            receptor,
		FechaDeTransaccion:  fecha_de_transaccion,
	}

	stmt, err := r.db.Prepare("UPDATE transactions SET codigo_de_transaccion = ?, moneda = ? , monto = ? , emisor = ?, receptor = ?,fecha_de_transaccion = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaccion.CodigoDeTransaccion, transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.FechaDeTransaccion, transaccion.ID)

	if err != nil {
		return models.Transaction{}, err
	}
	filasActualizadas, err := result.RowsAffected()

	if err != nil {
		return models.Transaction{}, err
	}

	if filasActualizadas == 0 {
		return models.Transaction{}, errors.New("No se encontró la transaccion")
	}
	return transaccion, nil
}

func (r *repositorySQLMock) GetAll() ([]models.Transaction, error) {
	var transaccLeidas []models.Transaction
	var transaccLeida models.Transaction

	rows, err := r.db.Query("SELECT id,codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion FROM transactions")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&transaccLeida.ID, &transaccLeida.CodigoDeTransaccion, &transaccLeida.Moneda, &transaccLeida.Monto, &transaccLeida.Emisor, &transaccLeida.Receptor, &transaccLeida.FechaDeTransaccion)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		transaccLeidas = append(transaccLeidas, transaccLeida)
	}
	return transaccLeidas, nil
}

func (r *repositorySQLMock) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM transactions WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	filasActualizadas, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if filasActualizadas == 0 {
		return errors.New("No se encontró la transaccion")
	}
	return nil
}

// func (r *repositorySQL) GetByCode(code string) (models.Transaction,error){
//
// 	r.db.Query("SELECT id,codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion FROM transactions WHERE codigo_de_transaccion = ?", code)

// }

func (r *repositorySQLMock) UpdateWithContext(ctx context.Context, id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	//La única diferencia con el Update común es que se usa PrepareContext y ExecContext en vez de Prepare y Exec
	transaccion := models.Transaction{
		ID:                  id,
		CodigoDeTransaccion: codigo_de_transaccion,
		Moneda:              moneda,
		Monto:               monto,
		Emisor:              emisor,
		Receptor:            receptor,
		FechaDeTransaccion:  fecha_de_transaccion,
	}

	stmt, err := r.db.PrepareContext(ctx, "UPDATE transactions SET codigo_de_transaccion = ?, moneda = ? , monto = ? , emisor = ?, receptor = ?,fecha_de_transaccion = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, transaccion.CodigoDeTransaccion, transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.FechaDeTransaccion, transaccion.ID)

	if err != nil {
		return models.Transaction{}, err
	}
	filasActualizadas, err := result.RowsAffected()

	if err != nil {
		return models.Transaction{}, err
	}

	if filasActualizadas == 0 {
		return models.Transaction{}, errors.New("No se encontró la transaccion")
	}
	return transaccion, nil
}
