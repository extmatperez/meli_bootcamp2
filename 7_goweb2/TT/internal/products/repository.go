package internal

import "errors"

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name" binding:"required"`
	Color      string `json:"color" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Code       string `json:"code" binding:"required"`
	Published  bool   `json:"published" binding:"required"`
	Created_at string `json:"created_at" binding:"required"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int64, name string, color string, stock int, code string, published bool, createdAt string) (Product, error)
	FindById(id int64) (Product, error)
	LastId() (int64, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id int64, name string, color string, stock int, code string, published bool, createdAt string) (Product, error) {
	product := Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}

	products = append(products, product)

	return product, nil
}

func (r *repository) FindById(id int64) (Product, error) {
	var product Product

	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			product = products[i]
			return product, nil
		}
	}

	return Product{}, errors.New("Product not found")
}

func (r *repository) LastId() (int64, error) {
	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func NewRepository() Repository {
	return &repository{}
}
