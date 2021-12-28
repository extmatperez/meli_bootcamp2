package internal

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TMan/proyecto/internal/models"
)

type RepositorySQL interface{
	GetByName(name string) models.Transaccion
	Store(transaccion models.Transaccion) (models.Transaccion, error)
	GetAll() []models.Transaccion
	GetFullData() ([]models.Transaccion, error)

	UpdateWithContext(ctx context.Context, transaccion models.Transaccion)(models.Transaccion, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) GetByName(name string) models.Transaccion {
	db := db.StorageDB

	var transaccionLeida models.Transaccion

	rows, err := db.Query("SELECT id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion FROM transacciones WHERE emisor = ?", name)

	if err != nil {
		log.Fatal(err)
		return transaccionLeida
	}

	for rows.Next() {
		err = rows.Scan(&transaccionLeida.ID, &transaccionLeida.CodigoTransaccion, &transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor, &transaccionLeida.FechaTransaccion)
		if err != nil {
			log.Fatal(err)
		}
	}
	return transaccionLeida
}

func (r *repositorySQL) Store(transaccion models.Transaccion) (models.Transaccion, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO transacciones(codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion) VALUES(?,?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result

	result, err = stmt.Exec(transaccion.CodigoTransaccion, transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.FechaTransaccion)

	if err != nil {
		return models.Transaccion{}, err
	}
	idCreado, _ := result.LastInsertId()
	transaccion.ID = int(idCreado)

	return transaccion, nil
}


func (r *repositorySQL) GetAll() []models.Transaccion {
	db := db.StorageDB

	var transaccionLeida models.Transaccion
	var transaccionesLeidas []models.Transaccion

	rows, err := db.Query("SELECT * FROM transacciones")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&transaccionLeida.ID, &transaccionLeida.CodigoTransaccion, &transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor, &transaccionLeida.FechaTransaccion)
		if err != nil {
			log.Fatal(err)
		}
		transaccionesLeidas = append(transaccionesLeidas, transaccionLeida)
	}
	return transaccionesLeidas
}

func (r *repositorySQL) UpdateWithContext(ctx context.Context, transaccion models.Transaccion)(models.Transaccion, error){
	db := db.StorageDB

	stmt, err := db.PrepareContext(ctx, "UPDATE transacciones SET codigo_transaccion = ?, moneda = ?, monto = ?, emisor = ?, receptor = ?, fecha_transaccion = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaccion.CodigoTransaccion, transaccion.Moneda, transaccion.Monto, transaccion.Emisor, transaccion.Receptor, transaccion.FechaTransaccion, transaccion.ID)

	if err != nil {
		return models.Transaccion{}, err
	}

	filasActualizadas, _ := result.RowsAffected()

	if filasActualizadas == 0{
		return models.Transaccion{}, err
	}

	return transaccion, nil
}


func (r *repositorySQL) GetFullData() ([]models.Transaccion, error) {
	db := db.StorageDB

	var misTransacciones []models.Transaccion
	var transaccionLeida models.Transaccion

	rows, err := db.Query("SELECT t.id, t.codigo_transaccion, t.moneda, t.monto, t.emisor, t.receptor, t.fecha_transaccion, p.nombre, p.marca FROM transacciones t INNER JOIN productos p on t.id_producto = p.id")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&transaccionLeida.ID, &transaccionLeida.CodigoTransaccion, &transaccionLeida.Moneda, &transaccionLeida.Monto, &transaccionLeida.Emisor, &transaccionLeida.Receptor, &transaccionLeida.FechaTransaccion, &transaccionLeida.Articulo.Nombre, &transaccionLeida.Articulo.Marca)
		if err != nil {
			log.Fatal(err)
		}
		misTransacciones = append(misTransacciones, transaccionLeida)
	}
	return misTransacciones, nil
}