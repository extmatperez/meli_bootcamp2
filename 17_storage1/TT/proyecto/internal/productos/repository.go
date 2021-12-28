package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/pkg/store"
)

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

var productos []Producto
var lastId int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	LastId() (int, error)
	Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Producto, error) {

	err := r.db.Read(&productos)

	if err != nil {
		return []Producto{}, err
	}

	return productos, nil
}

func (r *repository) Store(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	err := r.db.Read(&productos)

	if err != nil {
		return Producto{}, err
	}

	nuevoProducto := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	productos = append(productos, nuevoProducto)

	err = r.db.Write(&productos)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}

func (r *repository) LastId() (int, error) {

	err := r.db.Read(&productos)

	if err != nil {
		return 0, err
	}

	if len(productos) == 0 {
		return 0, nil
	}

	lastId = productos[len(productos)-1].ID

	return lastId, nil
}

func (r *repository) Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	err := r.db.Read(&productos)

	if err != nil {
		return Producto{}, err
	}

	productoActualizado := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	for i, p := range productos {
		if p.ID == id {
			productos[i] = productoActualizado

			err := r.db.Write(&productos)

			if err != nil {
				return Producto{}, err
			}

			return productoActualizado, nil
		}
	}

	return Producto{}, fmt.Errorf("el producto %v no existe", id)
}

func (r *repository) Delete(id int) error {

	err := r.db.Read(&productos)

	if err != nil {
		return err
	}

	index := 0

	for i, p := range productos {

		if p.ID == id {
			index = i
			productos = append(productos[:index], productos[index+1:]...)

			err := r.db.Write(&productos)

			if err != nil {
				return err
			}

			return nil
		}

	}

	return fmt.Errorf("el producto %v no existe", id)
}

func (r *repository) UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error) {

	err := r.db.Read(&productos)

	if err != nil {
		return Producto{}, err
	}

	for i, p := range productos {

		if p.ID == id {
			productos[i].Nombre = nombre
			productos[i].Precio = precio

			err := r.db.Write(&productos)

			if err != nil {
				return Producto{}, err
			}

			return productos[i], nil
		}
	}

	return Producto{}, fmt.Errorf("el producto %v no existe", id)
}
