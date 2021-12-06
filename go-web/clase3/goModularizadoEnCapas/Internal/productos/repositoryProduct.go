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
	UpdateName(varID int, nameUpdate string) error
	Delete(varID int) error
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

func (repo *repository) Store(productCTX Product) (Product, error) {

	var productoAux []Product
	repo.db.Read(&productoAux)
	lastID = productCTX.Id
	productoAux = append(productoAux, productCTX)
	repo.db.Write(productoAux)
	return productCTX, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&producto)
	if err != nil {
		return 0, err
	}
	if len(producto) == 0 {
		return 0, nil
	}
	return producto[len(producto)-1].Id, nil
}

func (repo *repository) Update(varID int, productCTX Product) (Product, error) {
	var productoAux []Product
	repo.db.Read(&productoAux)
	for i, _ := range productoAux {
		if productoAux[i].Id == varID {
			productCTX.Id = varID
			productoAux[i] = productCTX
			repo.db.Write(productoAux)
			return productCTX, nil
		}
	}
	return Product{}, errors.New("No se ha encontrado el producto a actualizar")
}

func (repo *repository) UpdateName(varID int, nameUpdate string) error {
	var productoAux []Product
	repo.db.Read(&productoAux)
	for i, _ := range productoAux {
		if productoAux[i].Id == varID {
			productoAux[i].Nombre = nameUpdate
			repo.db.Write(productoAux)
			return nil
		}
	}
	return errors.New("No se ha encontrado el producto a actualizar")
}

func (repo *repository) Delete(varID int) error {
	var productoAux []Product
	var productoEliminado []Product
	var bandera bool = false
	repo.db.Read(&productoAux)
	for i, _ := range productoAux {
		if productoAux[i].Id != varID {
			productoEliminado = append(productoEliminado, productoAux[i])
		} else {
			bandera = true
		}
	}
	if bandera == false {
		return errors.New("No se ha encontrado el producto a eliminar")
	} else {
		repo.db.Write(&productoEliminado)
		return nil
	}
}
