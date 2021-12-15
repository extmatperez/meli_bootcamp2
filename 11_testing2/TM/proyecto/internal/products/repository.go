package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/pkg/store"
)

type Product struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error)
	//Store2(nuevoProduct Product)(Product,err error)
	LastId() (int, error)
	UpdateProd(id int, name string, price float64) (Product, error)
	Delete(id int) error
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

func (repo *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	prod := Product{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}

	products = append(products, prod)
	err2 := repo.db.Write(products)

	if err2 != nil {
		return Product{}, err2
	}
	return prod, nil
}

func (repo *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	prod := Product{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}
	for i := range products {
		if products[i].ID == id {
			products[i] = prod
			return prod, nil
		}
	}
	return Product{}, fmt.Errorf("El ID %d no existe \n", id)
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

func (repo *repository) UpdateProd(id int, name string, price float64) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	for i := range products {
		if products[i].ID == id {
			products[i].Nombre = name
			products[i].Precio = price
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("El ID %d no existe \n", id)
}
func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&products)
	if err != nil {
		return err
	}
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			err := repo.db.Write(products)
			return err
		}
	}
	return fmt.Errorf("El ID %d no existe \n", id)
}
