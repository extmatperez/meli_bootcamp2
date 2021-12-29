package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/18_storage2/PracticaTM/internal/models"
)

type RepositorySQLMock interface {
	Store(transaccion models.Transaccion) (models.Transaccion, error)
	GetOne(id int) models.Transaccion
	Update(transaccion models.Transaccion) (models.Transaccion, error)
	GetByName(name string) []models.Transaccion
	GetAll() ([]models.Transaccion, error)
	Delete(id int) error
	GetFullData() ([]models.Transaccion, error)

	GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error)
	UpdateWithContext(ctx context.Context, transaccion models.Transaccion) (models.Transaccion, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(transaccion models.Transaccion) (models.Transaccion, error) {
	insertQuery := "INSERT INTO transacciones(moneda, monto, emisor, receptor) VALUES(?,?,?,?)"
	stmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor)
	if err != nil {
		return models.Transaccion{}, err
	}
	idCreado, _ := result.LastInsertId()
	transaccion.CodTransaccion = strconv.FormatInt(idCreado, 10)
	return transaccion, nil
}

func (r *repositorySQLMock) GetOne(id int) models.Transaccion {
	var transaccionLeida models.Transaccion
	selectQuery := "SELECT moneda, monto, emisor, receptor FROM transacciones WHERE idtransacciones = ?"
	rows, err := r.db.Query(selectQuery, id)

	if err != nil {
		log.Fatal(err)
		return models.Transaccion{}
	}

	for rows.Next() {
		//Iteramos por cada row que trajo (Igualmente va a traer una sola)
		err = rows.Scan(&transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor)
		//El scan guarda lo que se obtuvo de la consulta en cada variable.
		if err != nil {
			log.Fatal(err)
			return transaccionLeida
		}
	}
	return transaccionLeida
}

func (r *repositorySQLMock) Update(transaccion models.Transaccion) (models.Transaccion, error) {
	updateQuery := "UPDATE transacciones SET moneda = ?, monto = ?, emisor = ?, receptor = ? WHERE idtransacciones = ?"
	stmt, err := r.db.Prepare(updateQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.CodTransaccion)
	if err != nil {
		return models.Transaccion{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Transaccion{}, errors.New("no se encontro la persona")
	}

	return transaccion, nil
}

func (r *repositorySQLMock) GetByName(name string) []models.Transaccion {
	selectQuery := "SELECT moneda, monto, emisor, receptor FROM transacciones WHERE emisor = ?"
	rows, err := r.db.Query(selectQuery, name)
	if err != nil {
		log.Fatal(err)
		return []models.Transaccion{}
	}

	var transacciones []models.Transaccion
	var transaccionLeida models.Transaccion
	for rows.Next() {
		//Iteramos por cada row que trajo
		err = rows.Scan(&transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor)
		transacciones = append(transacciones, transaccionLeida)
		//El scan guarda lo que se obtuvo de la consulta en los campos de la struct.
		if err != nil {
			log.Fatal(err)
			return transacciones
		}
	}
	return transacciones
}

func (r *repositorySQLMock) GetAll() ([]models.Transaccion, error) {
	var transacciones []models.Transaccion
	var transaccionLeida models.Transaccion
	selectQuery := "SELECT moneda, monto, emisor, receptor FROM transacciones"
	rows, err := r.db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
		return []models.Transaccion{}, err
	}

	for rows.Next() {
		//Iteramos por cada row que trajo
		err = rows.Scan(&transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor)
		//El scan guarda lo que se obtuvo de la consulta en los campos de la struct.
		if err != nil {
			log.Fatal(err)
			return transacciones, err
		}
		transacciones = append(transacciones, transaccionLeida)
	}
	return transacciones, nil
}

func (r *repositorySQLMock) Delete(id int) error {
	deleteQuery := "DELETE FROM transacciones WHERE idtransacciones = ?"
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("no se encontro la transaccion")
	}

	return nil
}

func (r *repositorySQLMock) GetFullData() ([]models.Transaccion, error) {
	var transacciones []models.Transaccion
	var transaccionLeida models.Transaccion
	selectQuery := "SELECT t.moneda, t.monto, t.emisor, t.receptor, c.nombre_ciudad, c.nombre_pais FROM transacciones t INNER JOIN ciudad c on t.id_ciudad = c.idciudad"
	rows, err := r.db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
		return []models.Transaccion{}, err
	}

	for rows.Next() {
		//Iteramos por cada row que trajo
		err = rows.Scan(&transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor, &transaccionLeida.Ciudad.Nombre, &transaccionLeida.Ciudad.NombrePais)
		//El scan guarda lo que se obtuvo de la consulta en los campos de la struct.
		if err != nil {
			log.Fatal(err)
			return transacciones, err
		}
		transacciones = append(transacciones, transaccionLeida)
	}
	return transacciones, nil
}

func (r *repositorySQLMock) GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error) {
	var transaccionLeida models.Transaccion
	// rows, err := db.QueryContext(ctx, "SELECT moneda, monto, emisor, receptor FROM transacciones WHERE idtransacciones = ?", id)
	// se utiliza una query que demore 30 segundos en ejecutarse
	getQuery := "SELECT SLEEP(5) FROM DUAL where 0 < ?"
	// ya no se usa db.Query sino db.QueryContext
	rows, err := r.db.QueryContext(ctx, getQuery, id)
	if err != nil {
		log.Fatal(err)
		return models.Transaccion{}, err
	}

	for rows.Next() {
		//Iteramos por cada row que trajo (Igualmente va a traer una sola)
		err = rows.Scan(&transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor)
		//El scan guarda lo que se obtuvo de la consulta en cada variable.
		if err != nil {
			log.Fatal(err)
			return transaccionLeida, err
		}
	}
	return transaccionLeida, nil
}

func (r *repositorySQLMock) UpdateWithContext(ctx context.Context, transaccion models.Transaccion) (models.Transaccion, error) {
	updateQuery := "UPDATE transacciones SET moneda = ?, monto = ?, emisor = ?, receptor = ? WHERE idtransacciones = ?"
	result, err := r.db.ExecContext(ctx, updateQuery, transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.CodTransaccion)
	if err != nil {
		return models.Transaccion{}, err
	}
	filasActualizadas, err := result.RowsAffected()
	fmt.Printf("\n\n%v\n\n", err)
	if filasActualizadas == 0 {
		return models.Transaccion{}, errors.New("no se encontro la transaccion")
	}

	return transaccion, nil
}
