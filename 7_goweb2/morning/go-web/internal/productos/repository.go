package internal

import "fmt"

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var productos []Producto
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	LastId() (int, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	UpdateName(id int, nombre string) (Producto, error)
	Delete(id int) (string, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Producto, error) {
	return productos, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	per := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	productos = append(productos, per)
	return per, nil
}

func (repo *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	per := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	for i, p := range productos {
		if p.Id == id {
			productos[i] = per
			return per, nil
		}
	}
	return Producto{}, fmt.Errorf("No se encontro el producto con id %d", id)
}

func (repo *repository) UpdateName(id int, nombre string) (Producto, error) {

	for i, p := range productos {
		if p.Id == id {
			productos[i].Nombre = nombre
			return productos[i], nil
		}
	}
	return Producto{}, fmt.Errorf("No se encontro el producto con id %d", id)
}

func (repo *repository) Delete(id int) (string, error) {

	for i, p := range productos {
		if p.Id == id {
			productos = append(productos[:i], productos[i+1:]...)
			return "Producto eliminado", nil
		}
	}
	return "", fmt.Errorf("No se encontro el producto con id %d", id)
}
