package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/pkg/store"
)

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	UpdateNombre(id int, nombre string) (Product, error)
	UpdatePrecio(id int, precio int) (Product, error)
	Delete(id int) error
	LastID() (int, error)
}
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}
func (repo *repository) GetAll() ([]Product, error) {
	repo.db.Read(&products)
	return products, nil
}
func (repo *repository) Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	repo.db.Read(&products)
	p := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	products = append(products, p)
	err := repo.db.Write(products)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
func (repo *repository) Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	pro := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	for i, p := range products {
		if p.ID == id {
			products[i] = pro
			err = repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return pro, nil
		}
	}

	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (repo *repository) UpdateNombre(id int, nombre string) (Product, error) {
	err := repo.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	for i, p := range products {
		if p.ID == id {
			products[i].Nombre = nombre
			err = repo.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (rep *repository) UpdatePrecio(id int, precio int) (Product, error) {
	err := rep.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	for i, p := range products {
		if p.ID == id {
			products[i].Precio = precio
			err = rep.db.Write(products)
			if err != nil {
				return Product{}, err
			}
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&products)
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			err = repo.db.Write(products)
			return err
		}
	}
	return fmt.Errorf("el prodcuto %d no existe", id)
}
func (repo *repository) LastID() (int, error) {
	err := repo.db.Read(&products)
	if len(products) == 0 {
		return 0, err
	}
	return products[len(products)-1].ID, nil
}
