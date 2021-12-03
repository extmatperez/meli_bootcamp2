package internal

import "fmt"

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Code    string  `json:"code"`
	Publish bool    `json:"publish"`
	Date    string  `json:"date"`
}

var products []Product
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
	LastId() (int, error)
	Update(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}
func (repo *repository) Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
	pro := Product{Id,
		Name,
		Color,
		Price,
		Stock,
		Code,
		Publish,
		Date}
	lastId = Id
	products = append(products, pro)
	return pro, nil
}
func (repo *repository) LastId() (int, error) {
	return lastId, nil
}

func (repo *repository) Update(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
	pr := Product{Name: Name, Color: Color, Price: Price, Stock: Stock, Code: Code, Publish: Publish, Date: Date}
	udapted := false

	for i := range products {
		if products[i].Id == Id {
			pr.Id = Id
			products[i] = pr
			udapted = true
		}
	}
	if !udapted {
		return Product{}, fmt.Errorf("Producto %d no encontrado", Id)
	}
	return pr, nil
}
