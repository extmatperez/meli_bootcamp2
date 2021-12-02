package internal

import (
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/pkg/store"
)

var productos []Producto

type Producto struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Color  string  `json:"color" `
	Precio float64 `json:"precio" `
}

type Repository interface {
	GetAll() ([]Producto, error)
	GetById(id int) (Producto, error)
	Store(id int, nombre, color string, precio float64) (Producto, error)
	GetLastId() (int, error)
	Update(id int, nombre, color string, precio float64) (Producto, error)
	UpdateNombrePrecio(id int, nombre string, precio float64) (Producto, error)
	Delete(id int) error
}

type repository struct {
	fileStore store.Store
}

func NewRepository(fileStore store.Store) Repository {
	return &repository{fileStore: fileStore}
}

func (r *repository) GetAll() ([]Producto, error) {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return productos, err
	}
	return productos, nil
}

func (r *repository) GetById(id int) (Producto, error) {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return Producto{}, err
	}
	for _, p := range productos {
		if p.Id == id {
			return p, nil
		}
	}
	return Producto{}, errors.New("producto con id no encontrado")
}
func (r *repository) Store(id int, nombre, color string, precio float64) (Producto, error) {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return Producto{}, err
	}
	p := Producto{id, nombre, color, precio}
	productos = append(productos, p)
	err = r.fileStore.Write(productos)
	if err != nil {
		return Producto{}, err
	}
	return p, nil
}
func (r *repository) GetLastId() (int, error) {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return 0, err
	}
	return productos[len(productos)-1].Id, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64) (Producto, error) {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return Producto{}, err
	}

	productoNuevo := Producto{id, nombre, color, precio}
	encontrado := false
	for i, v := range productos {
		if v.Id == id {
			productos[i] = productoNuevo
			encontrado = true
			break
		}
	}
	if encontrado {
		err = r.fileStore.Write(productos)
		if err != nil {
			return Producto{}, err
		}
	}
	return Producto{}, fmt.Errorf("no se encontro el producto con id: %v", id)
}

func (r *repository) Delete(id int) error {
	err := r.fileStore.Read(&productos)
	if err != nil {
		return err
	}

	for i, v := range productos {
		if v.Id == id {
			productos = append(productos[:i], productos[i+1:]...)
			return r.fileStore.Write(productos)
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
			err := r.fileStore.Write(productos)
			if err != nil {
				return Producto{}, err
			}
			return v, nil
		}
	}
	return Producto{}, fmt.Errorf("no se encontro el producto con id: %v", id)
}
