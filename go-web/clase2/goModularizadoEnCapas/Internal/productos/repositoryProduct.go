package internal

import "errors"

type Product struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

var producto []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(productoAux Product) (Product, error)
	LastId() (int, error)
	Update(varID int, producto Product) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Product, error) {
	return producto, nil
}

func (repo *repository) Store(productoAux Product) (Product, error) {
	lastID = productoAux.Id
	producto = append(producto, productoAux)
	return productoAux, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(varID int, productoAux Product) (Product, error) {

	for i, _ := range producto {
		if producto[i].Id == varID {
			productoAux.Id = varID
			producto[i] = productoAux
			return producto[i], nil
		}
	}
	return Product{}, errors.New("No se ha encontrado el producto a actualizar")
}
