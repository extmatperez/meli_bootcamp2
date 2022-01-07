package invoicers

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	GetAll() ([]models.Invoicer, error)
	GetByID(id int) (models.Invoicer, error)
	Store(date_time string, id_customer int, total float64) (models.Invoicer, error)
	Update(ctx context.Context, id int, date_time string, id_customer int, total float64) (models.Invoicer, error)
	Delete(ctx context.Context, id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const (
	getAll = `
		SELECT id, date_time, id_customer, total
		FROM invoicers
	`
	getByID = `
		SELECT id, date_time, id_customer, total
		FROM invoicers 
		WHERE id = ?
	`
	insertInvoicer = `
		INSERT INTO invoicers (date_time, id_customer, total)
		VALUES ( ?, ?, ? )
	`
	updateInvoicer = `
		UPDATE invoicers
		SET
		date_time = ?,
		id_customer = ?,
		total = ?
		WHERE id = ?
	`
	deleteInvoicer = `
		DELETE FROM invoicers
		WHERE id = ?
	`
)

func (r *repository) GetAll() ([]models.Invoicer, error) {
	var invoicer models.Invoicer
	var invoicers []models.Invoicer
	db := db.StorageDB

	rows, err := db.Query(getAll)
	if err != nil {
		log.Println(err)
		return invoicers, err
	}
	for rows.Next() {
		if err := rows.Scan(&invoicer.ID, &invoicer.Date_Time, &invoicer.ID_Customer, &invoicer.Total); err != nil {
			log.Println(err.Error())
			return []models.Invoicer{}, err
		}
		invoicers = append(invoicers, invoicer)
	}
	return invoicers, nil
}

func (r *repository) GetByID(id int) (models.Invoicer, error) {
	var invoicer models.Invoicer
	db := db.StorageDB

	rows, err := db.Query(getByID, id)
	if err != nil {
		log.Println(err)
		return invoicer, err
	}
	for rows.Next() {
		if err := rows.Scan(&invoicer.ID, &invoicer.Date_Time, &invoicer.ID_Customer, &invoicer.Total); err != nil {
			log.Println(err.Error())
			return models.Invoicer{}, err
		}
	}
	return invoicer, nil
}

func (r *repository) Store(date_time string, id_customer int, total float64) (models.Invoicer, error) {

	db := db.StorageDB
	var invoicer models.Invoicer = models.Invoicer{
		Date_Time:   date_time,
		ID_Customer: id_customer,
		Total:       total,
	}

	stmt, err := db.Prepare(insertInvoicer)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(date_time, id_customer, total)
	if err != nil {
		return models.Invoicer{}, err
	}

	insertedId, _ := result.LastInsertId()
	invoicer.ID = int(insertedId)
	return invoicer, nil
}

func (r *repository) Update(ctx context.Context, id int, date_time string, id_customer int, total float64) (models.Invoicer, error) {

	db := db.StorageDB
	var invoicer models.Invoicer = models.Invoicer{
		Date_Time:   date_time,
		ID_Customer: id_customer,
		Total:       total,
	}

	stmt, err := db.PrepareContext(ctx, updateInvoicer)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, date_time, id_customer, total, id)
	if err != nil {
		return models.Invoicer{}, err
	}

	invoicer.ID = id
	return invoicer, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := db.StorageDB
	stmt, err := db.PrepareContext(ctx, deleteInvoicer)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
