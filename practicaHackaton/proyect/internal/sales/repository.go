package sales

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	GetAll() ([]models.Sales, error)
	GetByID(id int) (models.Sales, error)
	Store(id_invoice, id_product int, quantity float64) (models.Sales, error)
	Update(ctx context.Context, id int, id_invoice, id_product int, quantity float64) (models.Sales, error)
	Delete(ctx context.Context, id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const (
	getAll = `
		SELECT id, id_invoice, id_product, quantity 
		FROM sales
	`
	getByID = `
		SELECT id, id_invoice, id_product, quantity
		FROM sales 
		WHERE id = ?
	`
	insertSales = `
		INSERT INTO sales (id_invoice, id_product, quantity)
		VALUES ( ?, ?, ?)
	`
	updateSales = `
		UPDATE sales
		SET
		id_invoice = ?,
		id_product = ?,
		quantity = ?
		WHERE id = ?
	`
	deleteSales = `
		DELETE FROM sales
		WHERE id = ?
	`
)

func (r *repository) GetAll() ([]models.Sales, error) {
	var sale models.Sales
	var sales []models.Sales
	db := db.StorageDB

	rows, err := db.Query(getAll)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	for rows.Next() {
		if err := rows.Scan(&sale.ID, &sale.ID_Invoice, &sale.ID_Product, &sale.Quantity); err != nil {
			log.Println(err.Error())
			return []models.Sales{}, err
		}
		sales = append(sales, sale)
	}
	return sales, nil
}

func (r *repository) GetByID(id int) (models.Sales, error) {
	var sale models.Sales
	db := db.StorageDB

	rows, err := db.Query(getByID, id)
	if err != nil {
		log.Println(err)
		return sale, err
	}
	for rows.Next() {
		if err := rows.Scan(&sale.ID, &sale.ID_Invoice, &sale.ID_Product, &sale.Quantity); err != nil {
			log.Println(err.Error())
			return models.Sales{}, err
		}
	}
	return sale, nil
}

func (r *repository) Store(id_invoice, id_product int, quantity float64) (models.Sales, error) {

	db := db.StorageDB
	var sale models.Sales = models.Sales{
		ID_Invoice: id_invoice,
		ID_Product: id_product,
		Quantity:   quantity,
	}

	stmt, err := db.Prepare(insertSales)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(id_invoice, id_product, quantity)
	if err != nil {
		return models.Sales{}, err
	}

	insertedId, _ := result.LastInsertId()
	sale.ID = int(insertedId)
	return sale, nil
}

func (r *repository) Update(ctx context.Context, id int, id_invoice, id_product int, quantity float64) (models.Sales, error) {

	db := db.StorageDB
	var sale models.Sales = models.Sales{
		ID_Invoice: id_invoice,
		ID_Product: id_product,
		Quantity:   quantity,
	}

	stmt, err := db.PrepareContext(ctx, updateSales)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id_invoice, id_product, quantity, id)
	if err != nil {
		return models.Sales{}, err
	}

	sale.ID = id
	return sale, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := db.StorageDB
	stmt, err := db.PrepareContext(ctx, deleteSales)
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
