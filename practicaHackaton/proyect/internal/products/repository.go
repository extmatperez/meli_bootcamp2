package products

import (
	"context"
	"database/sql"
	"log"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (models.Product, error)
	Store(description string, price float64) (models.Product, error)
	Update(ctx context.Context, id int, description string, price float64) (models.Product, error)
	Delete(ctx context.Context, id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const (
	getAll = `
		SELECT id, description, price 
		FROM products
	`
	getByID = `
		SELECT id, description, price 
		FROM products 
		WHERE id = ?
	`
	insertProduct = `
		INSERT INTO products (description, price)
		VALUES ( ?, ?)
	`
	updateProduct = `
		UPDATE products
		SET
		description = ?,
		price = ?
		WHERE id = ?
	`
	deleteProduct = `
		DELETE FROM products
		WHERE id = ?
	`
)

func (r *repository) GetAll() ([]models.Product, error) {
	var product models.Product
	var products []models.Product
	db := db.StorageDB

	rows, err := db.Query(getAll)
	if err != nil {
		log.Println(err)
		return products, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Description, &product.Price); err != nil {
			log.Println(err.Error())
			return []models.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) GetByID(id int) (models.Product, error) {
	var product models.Product
	db := db.StorageDB

	rows, err := db.Query(getByID, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Description, &product.Price); err != nil {
			log.Println(err.Error())
			return models.Product{}, err
		}
	}
	return product, nil
}

func (r *repository) Store(description string, price float64) (models.Product, error) {

	db := db.StorageDB
	var product models.Product = models.Product{
		Description: description,
		Price:       price,
	}

	stmt, err := db.Prepare(insertProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(description, price)
	if err != nil {
		return models.Product{}, err
	}

	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)
	return product, nil
}

func (r *repository) Update(ctx context.Context, id int, description string, price float64) (models.Product, error) {

	db := db.StorageDB
	var product models.Product = models.Product{
		Description: description,
		Price:       price,
	}

	stmt, err := db.PrepareContext(ctx, updateProduct)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, description, price, id)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = id
	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := db.StorageDB
	stmt, err := db.PrepareContext(ctx, deleteProduct)
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
