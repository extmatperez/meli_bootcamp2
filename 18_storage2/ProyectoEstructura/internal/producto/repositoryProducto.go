package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/12_testing3/ProyectoEstructura/pkg/store"
)

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

type repositoryProducto struct {
	db store.Store
}

func NewRepository(db store.Store) RepositoryProducto {
	return &repositoryProducto{db}
}

func (repo *repositoryProducto) GetAll() ([]Producto, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (repo *repositoryProducto) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	err := repo.db.Read(&productos)

	if err != nil {
		return Producto{}, err
	} else {
		prod := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
		lastID = id
		productos = append(productos, prod)

		err := repo.db.Write(productos)
		fmt.Println(err)
		if err != nil {
			return Producto{}, err
		}

		return prod, nil
	}

}

func (repo *repositoryProducto) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	err := repo.db.Read(&productos)

	if err != nil {
		return Producto{}, err
	} else {
		producto := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
		for i, v := range productos {
			if v.ID == id {
				productos[i] = producto
				err := repo.db.Write(productos)
				if err != nil {
					return Producto{}, err
				}
				return producto, nil
			}
		}
		return Producto{}, fmt.Errorf("No se pudo modificar el usuario con id : %d", id)
	}
}

func (repo *repositoryProducto) LastId() (int, error) {
	err := repo.db.Read(&productos)

	if err != nil {
		return 0, err
	}

	if len(productos) == 0 {
		return 0, nil
	}

	return productos[len(productos)-1].ID, nil
}

func (repo *repositoryProducto) Delete(id int) error {
	err := repo.db.Read(&productos)
	if err != nil {
		return err
	} else {
		for i, v := range productos {
			if v.ID == id {
				productos = append(productos[:i], productos[i+1:]...)
				err := repo.db.Write(productos)
				if err != nil {
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("No se encontro, ni pudo borrar el producto con id: %d", id)
	}
}

func (repo *repositoryProducto) UpdateNameAndPrice(id int, nombre string, precio float64) (Producto, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return Producto{}, err
	} else {

		for i, v := range productos {
			if v.ID == id {
				productos[i].Nombre = nombre
				productos[i].Precio = precio

				err := repo.db.Write(productos)
				if err != nil {
					return Producto{}, err
				}
				return productos[i], nil
			}
		}
		return Producto{}, fmt.Errorf("Error al modificar campo nombre y precio al producto con ID: %d", id)
	}

}
