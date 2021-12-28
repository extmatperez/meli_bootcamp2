package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/12_testing3/TT/ejercicio_2/pkg/store"
)

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

//var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
	LastId() (int, error)
	Update(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price float64) (Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (repo *repository) Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
	err1 := repo.db.Read(&products)
	if err1 != nil {
		return Product{}, err1
	}
	pro := Product{Id,
		Name,
		Color,
		Price,
		Stock,
		Code,
		Publish,
		Date}
	//lastId = Id
	products = append(products, pro)
	err := repo.db.Write(products)

	if err != nil {
		return Product{}, err
	}
	return pro, nil
}
func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].Id, nil
}

func (repo *repository) Update(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	pr := Product{Id: Id, Name: Name, Color: Color, Price: Price, Stock: Stock, Code: Code, Publish: Publish, Date: Date}

	for i, v := range products {
		if v.Id == Id {
			// pr.Id = Id
			products[i] = pr
			err := repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return pr, nil
		}
	}
	return Product{}, fmt.Errorf("Producto %d no encontrado", Id)

}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&products)
	if err != nil {
		return err
	}
	index := 0
	for i, v := range products {
		if v.Id == id {
			index = i
			products = append(products[:index], products[index+1:]...)
			err := repo.db.Write(products)
			return err
		}
	}

	return fmt.Errorf("Producto %d no encontrado:", id)

}
func (repo *repository) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	for i, v := range products {
		if v.Id == id {
			products[i].Name = name
			products[i].Price = price
			err := repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return products[i], nil
		}
	}

	return Product{}, fmt.Errorf("Producto %d no encontrado", id)

}
