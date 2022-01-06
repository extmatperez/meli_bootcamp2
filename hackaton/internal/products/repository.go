package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

var products []models.Product

type ProductRepository interface {
	ImportAllProducts() error
	StoreProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
}

type repository_product struct{}

func NewProductRepository() ProductRepository {
	return &repository_product{}
}

func (r *repository_product) ImportAllProducts() error {

}

func (r *repository_product) StoreProduct(product models.Product) (models.Product, error) {

}

func (r *repository_product) UpdateProduct(product models.Product) (models.Product, error) {

}
