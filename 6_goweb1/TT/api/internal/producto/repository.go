package internal

import (
	"encoding/json"
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
	getAll() ([]Product, error)
	store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	getLastID() int64
	loadFile() error
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) loadFile() error {
	data, err := os.ReadFile("./products.json")
	if err != nil {
		return err
	}

	json.Unmarshal(data, &products)
	return nil
}

func (repo *repository) getLastID() int64 {
	return int64(len(products) + 1)
}

func (repo *repository) store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	product := Product{id, name, color, price, stock, code, isPublished, createdAt} //creo un productos
	products = append(products, product)                                            //agrego el productor al slice
	return product, nil
}

func (repo *repository) getAll() ([]Product, error) {

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}
