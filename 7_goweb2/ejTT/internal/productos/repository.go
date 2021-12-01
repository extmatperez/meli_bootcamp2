package internal

import (
	"errors"
	"fmt"
)

var lastId int
var productos []Producto

type Producto struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color" `
	Precio         float64 `json:"precio" `
	Stock          int     `json:"stock" `
	Codigo         string  `json:"codigo" `
	Publicado      bool    `json:"publicado" `
	Fecha_creacion string  `json:"fecha_creacion"`
}

type Repository interface {
	GetAll() ([]Producto, error)
	GetById(id int) (Producto, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error)
	GetLastId() (int, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error)
	UpdateNombrePrecio(id int, nombre string, precio float64) (Producto, error)
	Delete(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	return productos, nil
}

func (r *repository) GetById(id int) (Producto, error) {
	for _, p := range productos {
		if p.Id == id {
			return p, nil
		}
	}
	return Producto{}, errors.New("Persona con id no encontrado")
}
func (r *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error) {
	p := Producto{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	lastId = id
	productos = append(productos, p)
	return p, nil
}
func (r *repository) GetLastId() (int, error) {
	return lastId, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error) {

	productoNuevo := Producto{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	for i, v := range productos {
		if v.Id == id {
			productos[i] = productoNuevo
			return productoNuevo, nil
		}
	}
	return Producto{}, fmt.Errorf("no se encontro el producto con id: %v", id)
}

func (r *repository) Delete(id int) error {

	for i, v := range productos {
		if v.Id == id {
			productos = append(productos[:i], productos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no se encontro el producto con id: %v", id)
}

func (r *repository) UpdateNombrePrecio(id int, nombre string, precio float64) (Producto, error) {

	for i, v := range productos {
		if v.Id == id {
			v.Nombre = nombre
			v.Precio = precio
			productos[i] = v
			return v, nil
		}
	}
	return Producto{}, fmt.Errorf("no se encontro el producto con id: %v", id)
}
