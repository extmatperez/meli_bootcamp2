package internal

import "fmt"

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

var products []Product
var lastID int

type Repository interface {
	LastId() (int, error)
	GetAll() ([]Product, error)
	Save(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error)
	Update(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repo *repository) Save(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error) {
	newProduct := Product{id, name, color, price, stock, code, published, creationDate}
	lastID = id
	products = append(products, newProduct)
	return newProduct, nil
}

func (repo *repository) Update(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error) {
	updatedProduct := Product{id, name, color, price, stock, code, published, creationDate}

	for i := range products {
		if products[i].Id == id {
			products[i] = updatedProduct
			return updatedProduct, nil
		}
	}
	return Product{}, fmt.Errorf("el ID %v no aparece", id)
}
