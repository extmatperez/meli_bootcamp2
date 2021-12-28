package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/pkg/store"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

type Repository interface {
	GetAll() ([]Product, error)
	// getProductbyID() (Product, error)
	AddProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error)
	UpdateProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error)
}

type repository struct {
	db store.Store
}

var prodList []Product
var lastIDrepo int

/////////////// FUNCIONES /////////////////

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetAll() ([]Product, error) {
	repo.db.Read(&prodList)
	return prodList, nil
}

func (repo *repository) AddProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	repo.db.Read(&prodList)
	newProduct := Product{id, name, color, price, stock, code, published, created}
	lastIDrepo = id
	prodList = append(prodList, newProduct)

	err := repo.db.Write(&prodList)
	if err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

func (repo *repository) UpdateProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	updatedProduct := Product{id, name, color, price, stock, code, published, created}

	for i, v := range prodList {
		if v.ID == id {
			prodList[i] = updatedProduct
			return updatedProduct, nil
		}
	}
	return Product{}, fmt.Errorf("el producto id: %d no se encontró", id)

}


