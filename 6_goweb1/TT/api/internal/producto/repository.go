package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/pkg/store"
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

type repository struct {
	db store.Store
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	GetLastID() (int64, error)
	Delete(id int64) (string, error)
	Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error)
	UpdateNombre(id int64, name string) (Product, error)
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetLastID() (int64, error) {
	err := repo.db.Read(&products)

	if err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].ID, nil
}

func (repo *repository) GetAll() ([]Product, error) {

	err := repo.db.Read(&products)

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, nil
	}

	return products, nil
}

func (repo *repository) Store(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	err := repo.db.Read(&products)

	if err != nil {
		return Product{}, err
	}

	product := Product{id, name, color, price, stock, code, isPublished, createdAt} //creo un productos
	products = append(products, product)                                            //agrego el productor al slice
	err = repo.db.Write(products)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (repo *repository) Delete(id int64) (string, error) {

	err := repo.db.Read(&products)

	if err != nil {
		return "", err
	}
	for k, v := range products {
		if v.ID == id {
			products = append(products[:k], products[k+1:]...)

			err = repo.db.Write(products)

			if err != nil {
				return "", err
			}

			return fmt.Sprintf("Se borro con exito el producto numero %v", id), nil
		}
	}
	return "", fmt.Errorf("No se ha encontrado la transaccion con id %v", id)

}

func (repo *repository) Update(id int64, name string, color string, price float64, stock int64, code string, isPublished bool, createdAt string) (Product, error) {

	prod := Product{id, name, color, price, stock, code, isPublished, createdAt}

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
	return Product{}, fmt.Errorf("No se encontro el producto con id %d", id)

}

func (repo *repository) UpdateNombre(id int64, name string) (Product, error) {
	for i, v := range products {
		if v.ID == id {
			products[i].Name = name
			err := repo.db.Write(products)

			if err != nil {
				return Product{}, err
			}
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("No se encontro el producto con id %d", id)
}
