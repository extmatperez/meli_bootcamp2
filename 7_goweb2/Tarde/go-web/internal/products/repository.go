package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Precio        int    `json:"precio" binding:"required"`
	Stock         int    `json:"stock" binding:"required"`
	Codigo        string `json:"codigo" binding:"required"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

var products []Product

var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	LastID() (int, error)
	Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio int) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	loadData()
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	prod := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	products = append(products, prod)

	return prod, nil
}

func (r *repository) LastID() (int, error) {

	var id int

	if len(products) > 0 {
		id = products[len(products)-1].ID
	} else {
		id = 0
	}
	return id, nil
}

func (r *repository) Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	prod := Product{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	for i, _ := range products {
		if products[i].ID == id {
			products[i] = prod
			return prod, nil
		}
	}
	return Product{}, fmt.Errorf("El producto con id: %d no ha sido encontado", id)
}

func (r *repository) Delete(id int) error {
	position := -1
	for i, _ := range products {
		if products[i].ID == id {
			position = i
			break
		}
	}
	if position < 0 {
		return fmt.Errorf("El producto con id: %d no ha sido encontrado", id)
	}
	products = append(products[:position], products[position+1:]...)
	return nil

}

func (r *repository) UpdateNameAndPrice(id int, nombre string, precio int) (Product, error) {

	for i, _ := range products {
		if products[i].ID == id {
			products[i].Nombre = nombre
			products[i].Precio = precio

			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("El producto con id: %d no ha sido encontrado", id)
}

func loadData() {
	content, err := os.ReadFile("../../internal/products/products.json")

	if err != nil {
		fmt.Println(err)
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}

	p := []Product{}

	json.Unmarshal(content, &p)

	products = p
}
