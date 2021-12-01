package internal

import (
	"fmt"
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
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	UpdateNombre(id int, nombre string) (Product, error)
	UpdatePrecio(id int, precio int) (Product, error)
	Delete(id int) error
	LastID() (int, error)
}
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
func (repo *repository) GetAll() ([]Product, error) {
	return products, nil
}
func (repo *repository) Store(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	p := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	products = append(products, p)
	return p, nil
}
func (repo *repository) Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	pro := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	for i, p := range products {
		if p.ID == id {
			products[i] = pro
			return pro, nil
		}
	}
	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (repo *repository) UpdateNombre(id int, nombre string) (Product, error) {
	for i, p := range products {
		if p.ID == id {
			products[i].Nombre = nombre
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (rep *repository) UpdatePrecio(id int, precio int) (Product, error) {
	for i, p := range products {
		if p.ID == id {
			products[i].Precio = precio
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("el prodcuto %d no existe", id)
}
func (repo *repository) Delete(id int) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("el prodcuto %d no existe", id)

}
func (repo *repository) LastID() (int, error) {
	return lastID, nil
}
