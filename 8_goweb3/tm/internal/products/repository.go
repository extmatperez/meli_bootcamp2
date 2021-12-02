package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/pkg/store"
)

type Product struct {
	ID          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Code        int     `json:"code" binding:"required"`
	IsPublished bool    `json:"ispublished" binding:"required"`
	CreatedAt   string  `json:"createdat" binding:"required"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(ID int, Name, Color string, Price float64, Stock, Code int, isPublished bool, CreatedAt string) (Product, error)
	Update(ID int, Name, Color string, Price float64, Stock, Code int, isPublished bool, CreatedAt string) (Product, error)
	Delete(ID int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetAll() ([]Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *repository) Store(id int, name, color string, price float64, stock, code int, isPublished bool, createdAt string) (Product, error) {
	prod := Product{id, name, color, price, stock, code, isPublished, createdAt}
	products = append(products, prod)
	err := repo.db.Write(products)
	if err != nil {
		return Product{}, err
	}
	return prod, nil
}

func (repo *repository) Update(id int, name, color string, price float64, stock, code int, isPublished bool, createdAt string) (Product, error) {
	prod := Product{id, name, color, price, stock, code, isPublished, createdAt}
	updated := false
	for i := range products {
		if products[i].ID == id {
			products[i] = prod
			err := repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto no encontrado")
	}
	return prod, nil
}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&products)
	if err != nil {
		return err
	}
	index := 0
	for i, v := range products {
		if v.ID == id {
			index = i
			products = append(products[:id], products[id+1:]...)
			err := repo.db.Write(products)
			return err
		}
	}
	return fmt.Errorf("el producto con id %d no existe", id)
}
