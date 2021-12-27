package internal

import (
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/pkg/db"
)

type RepositorySql interface {
	Store(payment models.Payment) (models.Payment, error)
	GetById(id int) models.Payment
	GetByCode(codigo string) models.Payment
	GetAllPayments() []models.Payment
	Update(payment models.Payment) (models.Payment, error)
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

func (r *repositorySql) GetAllPayments() []models.Payment {
	var pays []models.Payment
	db := db.StorageDB
	rows, err := db.Query("SELECT id, codigo, moneda, monto, emisor, receptor, fecha FROM Payments")
	if err != nil {
		log.Fatal(err)
		return pays
	}

	// Se recorre el resultado de la query.
	for rows.Next() {
		var pay models.Payment
		err := rows.Scan(&pay.Id, &pay.Codigo, &pay.Moneda, &pay.Monto, &pay.Emisor, &pay.Receptor, &pay.Fecha)
		if err != nil {
			log.Fatal(err)
			return pays
		}
		pays = append(pays, pay)
	}
	return pays
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
