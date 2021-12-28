package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/domain"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/pkg/database"
)

var (
	GetByNameLikeQuery = "SELECT id, name, price, description FROM products WHERE LOWER(name) LIKE (\"%?%\")"
	StoreStatement     = "INSERT INTO products(name, price, description) VALUES(?, ?, ?)"
)

func NewRepository() Repository {
	return &repository{}
}

type Repository interface {
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
}

type repository struct{}

func (r *repository) GetByName(ctx context.Context, name string) ([]domain.Product, error) {
	db := database.StorageDB

	rows, err := db.Query("SELECT id, name, price, description FROM products WHERE LOWER(name) LIKE (\"%" + name + "%\")")

	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		var auxProduct domain.Product
		if err = rows.Scan(&auxProduct.Id, &auxProduct.Name, &auxProduct.Price, &auxProduct.Description); err != nil {
			continue
		}

		products = append(products, auxProduct)
	}

	return products, nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	db := database.StorageDB

	stmt, err := db.Prepare(StoreStatement)
	if err != nil {
		return domain.Product{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price, product.Description)

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
