package internal

import "fmt"

type Producto struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var productos []Producto
var lastID int

type RepositoryProducto interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	LastId() (int, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio float64) (Producto, error)
}

type repositoryProducto struct{}

func NewRepository() RepositoryProducto {
	return &repositoryProducto{}
}

func (repo *repositoryProducto) GetAll() ([]Producto, error) {
	return productos, nil
}

func (repo *repositoryProducto) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	prod := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = id
	productos = append(productos, prod)
	return prod, nil
}

func (repo *repositoryProducto) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	producto := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	for i, v := range productos {
		if v.ID == id {
			productos[i] = producto
			return producto, nil
		}
	}
	return Producto{}, fmt.Errorf("No se pudo modificar el usuario con id : %d", id)
}

func (repo *repositoryProducto) LastId() (int, error) {
	return lastID, nil
}

func (repo *repositoryProducto) Delete(id int) error {

	for i, v := range productos {
		if v.ID == id {
			productos = append(productos[:i], productos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No se encontro, ni pudo borrar el producto con id: %d", id)
}

func (repo *repositoryProducto) UpdateNameAndPrice(id int, nombre string, precio float64) (Producto, error) {
	for i, v := range productos {
		if v.ID == id {
			productos[i].Nombre = nombre
			productos[i].Precio = precio

			return productos[i], nil
		}
	}
	return Producto{}, fmt.Errorf("Error al modificar campo nombre y precio al producto con ID: %d", id)
}
