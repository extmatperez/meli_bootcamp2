package internal

import (
	"errors"

	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase3/goModularizadoEnCapas/pkg/store"
)

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

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Product, error) {
	var productoAux []Product
	repo.db.Read(&productoAux)
	return productoAux, nil
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
