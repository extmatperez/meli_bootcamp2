package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/project/pkg/store"
)

type Products struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

var products []Products
var lastID int64

type Repository interface {
	GetAll() ([]Products, error)
	Store(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error)
	LastID() (int64, error)
	Update(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Products, error) {
	err := r.db.Read(&products)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) Store(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error) {
	r.db.Read(&products)

	prod := Products{id, name, color, price, stock, code, published, creationdate}

	products = append(products, prod)
	err := r.db.Write(products)

	if err != nil {
		return Products{}, nil
	}
	return prod, nil

}

func (r *repository) LastID() (int64, error) {
	err := r.db.Read(&products)

	if err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].ID, nil
}

func (r *repository) Update(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error) {
	err := r.db.Read(&products)
	if err != nil {
		return Products{}, err
	}

	p := Products{id, name, color, price, stock, code, published, creationdate}
	for i := range products {
		if products[i].ID == id {
			p.ID = id
			products[i] = p
			err := r.db.Write(products)

			if err != nil {
				return Products{}, nil
			}
			return p, nil
		}
	}
	return Products{}, fmt.Errorf("Producto %v no encontrado", id)
}
