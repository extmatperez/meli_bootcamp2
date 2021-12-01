package internal

import "fmt"

type Productos struct {
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_de_creacion string  `json:"fecha_de_creacion"`
}

var productos []Productos
var lastId int

type Repository interface {
	GetAll() ([]Productos, error)
	Store(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error)
	Modify(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error)
	ModifyNamePrice(id int, nombre string, precio float64) (Productos, error)
	Delete(id int) error
	LastId() (int, error)
}
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Productos, error) {
	return productos, nil
}

func (repo *repository) Store(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error) {
	prod := Productos{id, nombre, color, precio, stock, codigo, publicado, fecha_de_creacion}
	lastId = id

	productos = append(productos, prod)
	return prod, nil
}

func (repo *repository) Modify(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error) {
	prod := Productos{id, nombre, color, precio, stock, codigo, publicado, fecha_de_creacion}

	for k, v := range productos {
		if v.Id == id {
			productos[k] = prod
			return prod, nil
		}
	}
	return Productos{}, fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) ModifyNamePrice(id int, nombre string, precio float64) (Productos, error) {

	for k, v := range productos {
		if v.Id == id {
			productos[k].Nombre = nombre
			productos[k].Precio = precio
			return productos[k], nil
		}
	}
	return Productos{}, fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) Delete(id int) error {

	for k, v := range productos {
		if v.Id == id {
			index := k
			productos = append(productos[:index], productos[index+1:]...)
			return fmt.Errorf("success: El producto %d fue eliminado", id)
		}
	}
	return fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) LastId() (int, error) {
	return lastId, nil
}
