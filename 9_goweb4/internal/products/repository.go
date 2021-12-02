package internal

import (
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/store"
)

type Product struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Color      string  `json:"color" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Stock      int     `json:"stock" binding:"required"`
	Code       string  `json:"code" binding:"required"`
	Published  bool    `json:"published" binding:"required"`
	Created_at string  `json:"created_at" binding:"required"`
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int64, name string, color string, price float64, stock int, code string, published bool, createdAt string) (Product, error)
	FindById(id int64) (Product, error)
	LastId() (int64, error)
	Update(id int64, name string, color string, price float64, stock int, code string, published bool, createdAt string) (Product, error)
	Delete(id int64) error
	UpdateNameAndPrice(id int64, name string, price float64) (Product, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) Store(id int64, name string, color string, price float64, stock int, code string, published bool, createdAt string) (Product, error) {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return Product{}, err
	}

	product := Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Price:      price,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}

	products = append(products, product)

	err = r.db.Write(products)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *repository) FindById(id int64) (Product, error) {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return Product{}, err
	}

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
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func (r *repository) Update(id int64, name string, color string, price float64, stock int, code string, published bool, createdAt string) (Product, error) {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return Product{}, err
	}

	product := Product{
		Id:         id,
		Name:       name,
		Color:      color,
		Price:      price,
		Stock:      stock,
		Code:       code,
		Published:  published,
		Created_at: createdAt,
	}

	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products[i] = product

			err = r.db.Write(products)

			if err != nil {
				return Product{}, err
			}

			return product, nil
		}
	}

	return Product{}, fmt.Errorf("Product %d not found", id)
}

func (r *repository) Delete(id int64) error {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return err
	}

	i := 0
	indexToRemoveFound := false

	for i = 0; i < len(products); i++ {
		if products[i].Id == id {
			indexToRemoveFound = true
			break
		}
	}

	if !indexToRemoveFound {
		return fmt.Errorf("Product %d not fount", id)
	}

	products = append(products[:i], products[i+1:]...)

	err = r.db.Write(products)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateNameAndPrice(id int64, name string, price float64) (Product, error) {
	var products []Product
	err := r.db.Read(&products)

	if err != nil {
		return Product{}, err
	}

	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products[i].Name = name
			products[i].Price = price

			err = r.db.Write(products)

			if err != nil {
				return Product{}, err
			}

			return products[i], nil
		}
	}

	return Product{}, fmt.Errorf("Product %d not found", id)
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
