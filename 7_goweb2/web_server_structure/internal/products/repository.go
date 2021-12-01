package internal

import (
	"encoding/json"
	"log"
	"os"
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
	SetLastId() (int, error)
}

type repository struct {
}

var prodList []Product
var lastIDrepo int

// var tokenPrueba string

/////////////// FUNCIONES /////////////////

func NewRepository() Repository {
	return &repository{}
}

func ReadData() {

	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal(readProducts, &prodList); err != nil {
		log.Fatal(err)
	}

}

func (repo *repository) GetAll() ([]Product, error) {
	return prodList, nil
}

func (repo *repository) AddProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	newProduct := Product{id, name, color, price, stock, code, published, created}
	lastIDrepo = id
	prodList = append(prodList, newProduct)
	return newProduct, nil
}

func (repo *repository) SetLastId() (int, error) {
	length := len(prodList) - 1
	lastIDrepo = prodList[length].ID
	return lastIDrepo, nil
}
