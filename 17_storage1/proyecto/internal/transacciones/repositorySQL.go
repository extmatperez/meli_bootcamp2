package internal

import (
	"errors"
	"log"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
	db "github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/store"
)

const (
	StoreTrans  = "INSERT INTO transaccion (CodigoTransaccion, Moneda, Monto, Emisor, Receptor) values(?,?,?,?,?)"
	GetTrans    = "SELECT ID, CodigoTransaccion, Moneda, Monto, Emisor, Receptor, FechaCreacion FROM transaccion WHERE id = ?"
	UpdateTrans = "UPDATE users SET nombre = ?, apellido = ?, edad = ? WHERE id = ?"
	GetCodTrans = "SELECT ID, CodigoTransaccion, Moneda, Monto, Emisor, Receptor, FechaCreacion FROM transaccion WHERE CodigoTransaccion = ?"
)

type RepositorySql interface {
	Store(models.Transaccion) (models.Transaccion, error)
	GetOne(id int) models.Transaccion
	Update(models.Transaccion) (models.Transaccion, error)
	GetByCode(codeTrans string) ([]models.Transaccion, error)
}

type repositorySql struct {
}

func NewRepositorySql() RepositorySql {
	return &repositorySql{}
}

func (r *repositorySql) Store(trans models.Transaccion) (models.Transaccion, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(StoreTrans)

	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(trans.CodigoTransaccion, trans.Moneda, trans.Monto, trans.Emisor, trans.Receptor)
	if err != nil {
		return models.Transaccion{}, err
	}
	lastId, _ := result.LastInsertId()
	trans.ID = int(lastId)
	trans = r.GetOne(trans.ID)

	return trans, nil
}

func (r *repositorySql) GetOne(id int) models.Transaccion {
	db := db.StorageDB
	var trans models.Transaccion
	rows, err := db.Query(GetTrans, id)

	if err != nil {
		log.Fatal(err)
		return trans
	}

	for rows.Next() {
		err = rows.Scan(&trans.ID, &trans.CodigoTransaccion, &trans.Moneda, &trans.Monto, &trans.Emisor, &trans.Receptor, &trans.FechaCreacion)
		if err != nil {
			log.Fatal(err)
			return trans
		}
	}
	return trans
}

func (r *repositorySql) Update(trans models.Transaccion) (models.Transaccion, error) {

	db := db.StorageDB

	stmt, err := db.Prepare(UpdateTrans)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(trans.ID, trans.CodigoTransaccion, trans.Moneda, trans.Monto, trans.Emisor, trans.Receptor)
	if err != nil {
		return models.Transaccion{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Transaccion{}, errors.New("no se encontro la usuario")
	}

	trans = r.GetOne(trans.ID)

	return trans, nil
}
func (r *repositorySql) GetByCode(codeTrans string) ([]models.Transaccion, error) {
	db := db.StorageDB

	var transac []models.Transaccion
	rows, err := db.Query(GetCodTrans, codeTrans)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var trans models.Transaccion
		if err = rows.Scan(&trans.ID, &trans.CodigoTransaccion, &trans.Moneda, &trans.Monto, &trans.Emisor, &trans.Receptor, &trans.FechaCreacion); err != nil {
			log.Fatal(err)
			return transac, err
		}
		transac = append(transac, trans)
	}

	return transac, nil
}
