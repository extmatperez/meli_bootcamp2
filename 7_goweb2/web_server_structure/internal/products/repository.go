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
	getAll() ([]Product, error)
	getProductbyID() (Product, error)
	addProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error)
	setLastId() (int, error)
}

type repository struct {
}

var prodList []Product
var lastID int
var tokenPrueba string

/////////////// FUNCIONES /////////////////

func NewRepository() Repository {
	return &repository{}
}

func readData() {

	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal(readProducts, &prodList); err != nil {
		log.Fatal(err)
	}

}

func (repo *repository) getAll() ([]Product, error) {
	return prodList, nil
}

func (repo *repository) addProduct(id int, name, color string, price float64, stock, code int, published string, created string) (Product, error) {
	product := Product(id, name, color, price, stock, code, published, created)
	lastID = id
	prodList := append(prodList, product)
	return product, nil
}
