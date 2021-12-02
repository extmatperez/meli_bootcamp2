package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/afternoon/go-web/pkg/store"
)

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

type Repository interface {
	GetAll() ([]Producto, error)
	LastId() (int, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	UpdateName(id int, nombre string) (Producto, error)
	Delete(id int) (string, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (repo *repository) GetAll() ([]Producto, error) {

	var productos []Producto
	err := repo.db.Read(&productos)

	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (repo *repository) LastId() (int, error) {

	var productos []Producto

	if err := repo.db.Read(&productos); err != nil {
		return 0, err
	}
	if len(productos) == 0 {
		return 0, nil
	}
	return productos[len(productos)-1].Id, nil
}

func (repo *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	var productos []Producto

	err := repo.db.Read(&productos)
	if err != nil {
		return Producto{}, err
	}
	prod := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	productos = append(productos, prod)

	if err := repo.db.Write(productos); err != nil {
		return Producto{}, err
	}
	return prod, nil
}

func (repo *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	var productos []Producto
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
	var productos []Producto
	for i, p := range productos {
		if p.Id == id {
			productos[i].Nombre = nombre
			return productos[i], nil
		}
	}
	return Producto{}, fmt.Errorf("No se encontro el producto con id %d", id)
}

func (repo *repository) Delete(id int) (string, error) {
	var productos []Producto
	for i, p := range productos {
		if p.Id == id {
			productos = append(productos[:i], productos[i+1:]...)
			return "Producto eliminado", nil
		}
	}
	return "", fmt.Errorf("No se encontro el producto con id %d", id)
}
