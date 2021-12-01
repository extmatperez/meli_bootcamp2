package internal

import "errors"

var lastId int
var productos []Producto

type Producto struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre" binding:"required"`
	Color          string  `json:"color" binding:"required"`
	Precio         float64 `json:"precio" binding:"required"`
	Stock          int     `json:"stock" binding:"required"`
	Codigo         string  `json:"codigo" binding:"required"`
	Publicado      bool    `json:"publicado" binding:"required"`
	Fecha_creacion string  `json:"fecha_creacion" binding:"required"`
}

type Repository interface {
	GetAll() ([]Producto, error)
	GetById(id int) (Producto, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error)
	GetLastId() (int, error)
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
