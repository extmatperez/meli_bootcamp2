package internal

import (
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/8_goweb3/ejTTmodified/pkg/store"
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get all the products stored
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /producto [get]
func (r *repository) GetAll() ([]Producto, error) {
	r.fileStore.Read(&productos)

	return productos, nil
}

// ListProducts godoc
// @Summary Prodcut with id
// @Tags Products
// @Description get the product with the id given
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id url int true "id"
// @Success 200 {object} web.Response
// @Router /producto/:id [get]
func (r *repository) GetById(id int) (Producto, error) {
	r.fileStore.Read(&productos)

	for _, p := range productos {
		if p.Id == id {
			return p, nil
		}
	}
	return Producto{}, errors.New("producto con id no encontrado")
}
func (r *repository) Store(id int, nombre, color string, precio float64) (Producto, error) {
	r.fileStore.Read(&productos)

	p := Producto{id, nombre, color, precio}
	productos = append(productos, p)
	err := r.fileStore.Write(productos)
	if err != nil {
		return Producto{}, err
	}
	return p, nil
}
func (r *repository) GetLastId() (int, error) {
	r.fileStore.Read(&productos)

	if len(productos) == 0 {
		return 0, nil
	}
	return productos[len(productos)-1].Id, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64) (Producto, error) {
	r.fileStore.Read(&productos)

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
		err := r.fileStore.Write(productos)
		if err != nil {
			return Producto{}, err
		}
	}
	return Producto{}, fmt.Errorf("no se encontro el producto con id: %v", id)
}

func (r *repository) Delete(id int) error {
	r.fileStore.Read(&productos)

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
