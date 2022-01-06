package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

type ProductService interface {
	ImportAllProducts() error
	StoreProduct(description string, price float64) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
}

type service_product struct {
	repository_product ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &service_product{repository_product: repo}
}

func (s *service_product) ImportAllProducts() error {
	return s.repository_product.ImportAllProducts()
}

func (s *service_product) StoreProduct(description string, price float64) (models.Product, error) {

	new_product := models.Product{Id: 0,
		Description: description,
		Price:       price,
	}
	c, err := s.repository_product.StoreProduct(new_product)

	if err != nil {
		return models.Product{}, err
	}

	return c, nil
}

func (s *service_product) UpdateProduct(product models.Product) (models.Product, error) {
	return s.repository_product.UpdateProduct(product)
}
