package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"
)

type Service interface {
	//StoreOpcion2(product models.Product) (models.Product, error)
	Store(name string, price float64, size int) (models.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (ser *service) Store(name string, price float64, size int) (models.Product, error) {
	var newProduct = models.Product{Name: name, Price: price, Size: size}
	productCreated, err := ser.repository.Store(newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return productCreated, nil
}
