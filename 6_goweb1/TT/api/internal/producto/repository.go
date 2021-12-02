package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

var products []Product

type repository struct{}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	GetLastID() int64
	LoadFile() error
	Delete(id int64) (string, error)
	Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	UpdateNombre(id int64, name string) (Product, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) LoadFile() error {
	data, err := os.ReadFile("../../internal/producto/products.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &products)

	return err
}

func (repo *repository) GetLastID() int64 {
	return int64(len(products) + 1)
}

func (repo *repository) Store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	product := Product{id, name, color, price, stock, code, isPublished, createdAt} //creo un productos
	products = append(products, product)                                            //agrego el productor al slice
	return product, nil
}

func (repo *repository) GetAll() ([]Product, error) {

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (repo *repository) Delete(id int64) (string, error) {

	for k, v := range products {
		if v.ID == id {
			products = append(products[:k], products[k+1:]...)
			return "Se borro con exito el producto numero " + string(id), nil
		}
	}
	return "", fmt.Errorf("No se ha encontrado la transaccion con id %v", id)

}

func (repo *repository) Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	prod := Product{id, name, color, price, stock, code, isPublished, createdAt}

	for i, v := range products {
		if v.ID == id {
			products[i] = prod
			return prod, nil
		}
	}
	return Product{}, fmt.Errorf("No se encontro el producto con id %d", id)

}

func (repo *repository) UpdateNombre(id int64, name string) (Product, error) {
	for i, v := range products {
		if v.ID == id {
			products[i].Name = name
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("No se encontro el producto con id %d", id)
}
