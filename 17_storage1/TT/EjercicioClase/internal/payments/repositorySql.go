package internal

import (
	"context"
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/pkg/db"
)

type RepositorySql interface {
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

type repositorySql struct{}

func NewRepositorySql() RepositorySql {
	return &repositorySql{}
}

func (r *repositorySql) Store(payment models.Payment) (models.Payment, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO Payments(codigo, moneda, monto, emisor, receptor, fecha) VALUES (?,?,?,?,?,?)")
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

func (r *repositorySql) GetById(id int) models.Payment {
	var pay models.Payment
	db := db.StorageDB
	rows, err := db.Query("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?", id)
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

func (r *repositorySql) GetByCode(codigo string) models.Payment {
	var pay models.Payment
	db := db.StorageDB
	rows, err := db.Query("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE codigo = ?", codigo)
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

func (r *repositorySql) GetAllPayments() ([]models.Payment, error) {
	var pays []models.Payment
	db := db.StorageDB
	rows, err := db.Query("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments")
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

func (r *repositorySql) Update(payment models.Payment) (models.Payment, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("UPDATE Payments SET codigo = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha = ? WHERE id = ?")
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

func (r *repositorySql) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM Payments WHERE id = ?")
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

func (r *repositorySql) GetFullDataAllPayments() ([]models.Payment, error) {
	var pays []models.Payment
	db := db.StorageDB
	rows, err := db.Query("SELECT p.id, p.codigo, p.moneda, p.monto, p.emisor, p.receptor, p.fecha, b.id, b.responsable, b.fecha FROM Payments p INNER JOIN BoxClosing b ON p.box_closing_id = b.id")
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

func (r *repositorySql) GetByIdWithContext(ctx context.Context, id int) (models.Payment, error) {
	var pay models.Payment
	db := db.StorageDB
	rows, err := db.QueryContext(ctx, "SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments WHERE id = ?", id)
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

func (r *repositorySql) UpdateWithContext(ctx context.Context, payment models.Payment) (models.Payment, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("UPDATE Payments SET codigo = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha = ? WHERE id = ?")
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
