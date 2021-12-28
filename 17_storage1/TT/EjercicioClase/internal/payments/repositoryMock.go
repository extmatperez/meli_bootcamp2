package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/pkg/db"
)

type RepositorySqlMock interface {
	Store(payment models.Payment) (models.Payment, error)
	GetById(id int) models.Payment
	GetByCode(codigo string) models.Payment
	GetAllPayments() ([]models.Payment, error)
	Update(payment models.Payment) (models.Payment, error)
	Delete(id int) error
	GetFullDataAllPayments() ([]models.Payment, error)
	GetByIdWithContext(ctx context.Context, id int) (models.Payment, error)
	UpdateWithContext(ctx context.Context, payment models.Payment) (models.Payment, error)
}

type repositorySqlMock struct {
	db *sql.DB
}

var r = db.StorageDB

func NewRepositorySqlMock(db *sql.DB) RepositorySqlMock {
	return &repositorySqlMock{db: db}
}

func (r *repositorySqlMock) Store(payment models.Payment) (models.Payment, error) {
	query := "INSERT INTO Payments(codigo, moneda, monto, emisor, receptor, fecha) VALUES (?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Ejecutamos el comando con el payment ingresado.
	result, err := stmt.Exec(payment.Codigo, payment.Moneda, payment.Monto, payment.Emisor, payment.Receptor, payment.Fecha)
	if err != nil {
		return models.Payment{}, err
	}

	// Asi obtenemos el Id para insertarle.
	idCreado, _ := result.LastInsertId()
	payment.Id = int(idCreado)
	return payment, nil
}

func (r *repositorySqlMock) GetById(id int) models.Payment {
	var pay models.Payment
	query := "SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Fatal(err)
		return pay
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha)
		if err != nil {
			log.Fatal(err)
			return pay
		}
	}
	return pay
}

func (r *repositorySqlMock) GetByCode(codigo string) models.Payment {
	var pay models.Payment
	rows, err := r.db.Query("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE codigo = ?", codigo)
	if err != nil {
		log.Fatal(err)
		return pay
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha)
		if err != nil {
			log.Fatal(err)
			return pay
		}
	}
	return pay
}

func (r *repositorySqlMock) GetAllPayments() ([]models.Payment, error) {
	var pays []models.Payment
	query := "SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		var pay models.Payment
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		pays = append(pays, pay)
	}
	return pays, nil
}

func (r *repositorySqlMock) Update(payment models.Payment) (models.Payment, error) {
	query := "UPDATE Payments SET codigo = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(payment.Codigo, payment.Moneda, payment.Monto, payment.Emisor, payment.Receptor, payment.Fecha, payment.Id)
	if err != nil {
		return models.Payment{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Payment{}, errors.New("No se encontró la transacción.")
	}
	return payment, nil
}

func (r *repositorySqlMock) Delete(id int) error {
	query := "DELETE FROM Payments WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return errors.New("No se encontró la transacción.")
	}
	return nil
}

func (r *repositorySqlMock) GetFullDataAllPayments() ([]models.Payment, error) {
	var pays []models.Payment
	query := "SELECT p.id, p.codigo, p.moneda, p.monto, p.emisor, p.receptor, p.fecha, b.id, b.responsable, b.fecha FROM Payments p INNER JOIN BoxClosing b ON p.box_closing_id = b.id"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		var pay models.Payment
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha, &pay.BoxClosing.Id, &pay.BoxClosing.Responsable, &pay.BoxClosing.Fecha)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Aca se hace el append para el listado de elementos.
		pays = append(pays, pay)
	}
	return pays, nil
}

func (r *repositorySqlMock) GetByIdWithContext(ctx context.Context, id int) (models.Payment, error) {
	var pay models.Payment
	query := "SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal(err)
		return pay, err
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha)
		if err != nil {
			log.Fatal(err)
			return pay, err
		}
	}
	return pay, nil
}

func (r *repositorySqlMock) UpdateWithContext(ctx context.Context, payment models.Payment) (models.Payment, error) {
	query := "UPDATE Payments SET codigo = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, payment.Codigo, payment.Moneda, payment.Monto, payment.Emisor, payment.Receptor, payment.Fecha, payment.Id)
	if err != nil {
		return models.Payment{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Payment{}, errors.New("No se encontró la transacción.")
	}
	return payment, nil
}
