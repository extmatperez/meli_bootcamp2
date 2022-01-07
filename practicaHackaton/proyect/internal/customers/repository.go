package customers

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	GetAll() ([]models.Customer, error)
	GetByID(id int) (models.Customer, error)
	Store(last_name, first_name, condition string) (models.Customer, error)
	Update(ctx context.Context, id int, last_name, first_name, condition string) (models.Customer, error)
	Delete(ctx context.Context, id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const (
	getAll = `
		SELECT id, last_name, first_name, conditionn
		FROM customers
	`
	getByID = `
		SELECT id, last_name, first_name, conditionn
		FROM customers 
		WHERE id = ?
	`
	insertCustomer = `
		INSERT INTO customers (last_name, first_name, conditionn)
		VALUES ( ?, ?, ? )
	`
	updateCustomer = `
		UPDATE customers
		SET
		last_name = ?,
		first_name = ?,
		conditionn = ?
		WHERE id = ?
	`
	deleteCustomer = `
		DELETE FROM customers
		WHERE id = ?
	`
)

func (r *repository) GetAll() ([]models.Customer, error) {
	var customer models.Customer
	var customers []models.Customer
	db := db.StorageDB

	rows, err := db.Query(getAll)
	if err != nil {
		log.Println(err)
		return customers, err
	}
	for rows.Next() {
		if err := rows.Scan(&customer.ID, &customer.Last_Name, &customer.First_Name, &customer.Condition); err != nil {
			log.Println(err.Error())
			return []models.Customer{}, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *repository) GetByID(id int) (models.Customer, error) {
	var customer models.Customer
	db := db.StorageDB

	rows, err := db.Query(getByID, id)
	if err != nil {
		log.Println(err)
		return customer, err
	}
	for rows.Next() {
		if err := rows.Scan(&customer.ID, &customer.Last_Name, &customer.First_Name, &customer.Condition); err != nil {
			log.Println(err.Error())
			return models.Customer{}, err
		}
	}
	return customer, nil
}

func (r *repository) Store(last_name, first_name, condition string) (models.Customer, error) {

	db := db.StorageDB
	var customer models.Customer = models.Customer{
		Last_Name:  last_name,
		First_Name: first_name,
		Condition:  condition,
	}

	stmt, err := db.Prepare(insertCustomer)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(last_name, first_name, condition)
	if err != nil {
		return models.Customer{}, err
	}

	insertedId, _ := result.LastInsertId()
	customer.ID = int(insertedId)
	return customer, nil
}

func (r *repository) Update(ctx context.Context, id int, last_name, first_name, condition string) (models.Customer, error) {

	db := db.StorageDB
	var customer models.Customer = models.Customer{
		Last_Name:  last_name,
		First_Name: first_name,
		Condition:  condition,
	}

	stmt, err := db.PrepareContext(ctx, updateCustomer)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, last_name, first_name, condition, id)
	if err != nil {
		return models.Customer{}, err
	}

	customer.ID = id
	return customer, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := db.StorageDB
	stmt, err := db.PrepareContext(ctx, deleteCustomer)
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
