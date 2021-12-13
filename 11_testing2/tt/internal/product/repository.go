package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/11_testing2/tm/pkg/store"
)

type Product struct {
	ID     int     `json:"id"`
	Color  string  `json:"color"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, color string, price float64, amount int) (Product, error)
	Update(id int, color string, price float64, amount int) (Product, error)
	UpdatePrice(id int, price float64) (Product, error)
	Delete(id int) error
	LastId() (int, error)
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

func (repo *repository) Store(id int, color string, price float64, amount int) (Product, error) {
	repo.db.Read(&products)

	prod := Product{id, color, price, amount}

	products = append(products, prod)
	err := repo.db.Write(products)

	if err != nil {
		return Product{}, err
	}

	return prod, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&products)

	if err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].ID, nil
}

func (repo *repository) Update(id int, color string, price float64, amount int) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	prod := Product{id, color, price, amount}
	for i, v := range products {
		if v.ID == id {
			products[i] = prod
			err := repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return prod, nil
		}
	}
	return Product{}, fmt.Errorf("el producto %d no existe", id)

}
func (repo *repository) UpdatePrice(id int, price float64) (Product, error) {
	for i, v := range products {
		if v.ID == id {
			products[i].Price = price
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("el producto %d no existe", id)

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
			products = append(products[:index], products[index+1:]...)
			err := repo.db.Write(products)

			return err
		}
	}
	return fmt.Errorf("el producto %d no existe", id)

}
