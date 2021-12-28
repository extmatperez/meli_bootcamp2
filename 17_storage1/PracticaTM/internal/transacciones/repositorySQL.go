package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/pkg/db"
)

type RepositorySQL interface {
	Store(transaccion models.Transaccion) (models.Transaccion, error)
	GetOne(id int) models.Transaccion
	Update(transaccion models.Transaccion) (models.Transaccion, error)
	GetByName(name string) []models.Transaccion
	GetAll() ([]models.Transaccion, error)
	Delete(id int) error
	GetFullData() ([]models.Transaccion, error)

	GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(transaccion models.Transaccion) (models.Transaccion, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("INSERT INTO transacciones(moneda, monto, emisor, receptor) VALUES(?,?,?,?)")
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

func (r *repositorySQL) GetOne(id int) models.Transaccion {
	db := db.StorageDB
	var transaccionLeida models.Transaccion
	rows, err := db.Query("SELECT moneda, monto, emisor, receptor FROM transacciones WHERE idtransacciones = ?", id)

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

func (r *repositorySQL) Update(transaccion models.Transaccion) (models.Transaccion, error) {
	db := db.StorageDB
	stmt, err := db.Prepare("UPDATE transacciones SET moneda = ?, monto = ?, emisor = ?, receptor = ? WHERE idtransacciones = ?")
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

func (r *repositorySQL) GetByName(name string) []models.Transaccion {
	db := db.StorageDB
	rows, err := db.Query("SELECT moneda, monto, emisor, receptor FROM transacciones WHERE emisor = ?", name)
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

func (r *repositorySQL) GetAll() ([]models.Transaccion, error) {
	db := db.StorageDB
	var transacciones []models.Transaccion
	var transaccionLeida models.Transaccion
	rows, err := db.Query("SELECT moneda, monto, emisor, receptor FROM transacciones")
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

func (r *repositorySQL) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM transacciones WHERE idtransacciones = ?")
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

func (r *repositorySQL) GetFullData() ([]models.Transaccion, error) {
	db := db.StorageDB
	var transacciones []models.Transaccion
	var transaccionLeida models.Transaccion
	selectQuery := "SELECT t.moneda, t.monto, t.emisor, t.receptor, c.nombre_ciudad, c.nombre_pais FROM transacciones t INNER JOIN ciudad c on t.id_ciudad = c.idciudad"
	rows, err := db.Query(selectQuery)
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

func (r *repositorySQL) GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error) {
	db := db.StorageDB
	var transaccionLeida models.Transaccion
	// rows, err := db.QueryContext(ctx, "SELECT moneda, monto, emisor, receptor FROM transacciones WHERE idtransacciones = ?", id)
	// se utiliza una query que demore 30 segundos en ejecutarse
	getQuery := "SELECT SLEEP(30) FROM DUAL where 0 < ?"
	// ya no se usa db.Query sino db.QueryContext
	rows, err := db.QueryContext(ctx, getQuery, id)
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
