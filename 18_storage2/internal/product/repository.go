package product

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/domain"
)

var (
	GetAllQuery        = "SELECT id, name, price, description FROM products"
	GetQuery           = "SELECT id, name, price, description FROM products WHERE id = ?"
	GetByNameLikeQuery = "SELECT id, name, price, description FROM products WHERE LOWER(name) LIKE (\"%?%\")"
	StoreStatement     = "INSERT INTO products(name, price, description) VALUES(?, ?, ?)"
	UpdateStatement    = "UPDATE products SET name = ?, price = ?, description = ? WHERE id = ?"
	DeleteStatement    = "DELETE FROM products WHERE id = ?"
)

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(ctx, GetAllQuery)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		var auxProduct domain.Product
		if err = rows.Scan(&auxProduct.Id, &auxProduct.Name, &auxProduct.Price, &auxProduct.Description); err != nil {
			return nil, err
		}

		products = append(products, auxProduct)
	}

	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product

	err := r.db.QueryRowContext(ctx, GetQuery, id).Scan(&product.Id, &product.Name, &product.Price, &product.Description)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) GetByName(ctx context.Context, name string) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, price, description FROM products WHERE LOWER(name) LIKE (\"%"+name+"%\")")
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		var auxProduct domain.Product
		if err = rows.Scan(&auxProduct.Id, &auxProduct.Name, &auxProduct.Price, &auxProduct.Description); err != nil {
			return nil, err
		}

		products = append(products, auxProduct)
	}

	return products, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.PrepareContext(ctx, StoreStatement)
	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Description)
	if err != nil {
		return domain.Product{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	product.Id = int(lastId)

	return product, nil
}

func (r *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := r.db.PrepareContext(ctx, UpdateStatement)
	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Description, product.Id)
	if err != nil {
		return domain.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.PrepareContext(ctx, DeleteStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}

	return nil
}
