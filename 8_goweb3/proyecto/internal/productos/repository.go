package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/pkg/store"
)

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
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Productos, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (repo *repository) Store(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return Productos{}, err
	}
	prod := Productos{id, nombre, color, precio, stock, codigo, publicado, fecha_de_creacion}
	lastId = id
	productos = append(productos, prod)

	err = repo.db.Write(productos)

	if err != nil {
		return Productos{}, err
	}

	return prod, nil
}

func (repo *repository) Modify(id, stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return Productos{}, err
	}
	prod := Productos{id, nombre, color, precio, stock, codigo, publicado, fecha_de_creacion}

	for k, v := range productos {
		if v.Id == id {
			productos[k] = prod
			err := repo.db.Write(&productos)
			if err != nil {
				return Productos{}, err
			}
			return prod, nil
		}
	}
	return Productos{}, fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) ModifyNamePrice(id int, nombre string, precio float64) (Productos, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return Productos{}, err
	}
	for k, v := range productos {
		if v.Id == id {
			productos[k].Nombre = nombre
			productos[k].Precio = precio
			err := repo.db.Write(&productos)
			if err != nil {
				return Productos{}, err
			}
			return productos[k], nil
		}
	}
	return Productos{}, fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&productos)
	if err != nil {
		return err
	}
	for k, v := range productos {
		if v.Id == id {
			index := k
			productos = append(productos[:index], productos[index+1:]...)
			err := repo.db.Write(&productos)
			if err != nil {
				return err
			}
			return fmt.Errorf("success: El producto %d fue eliminado", id)
		}
	}
	return fmt.Errorf("error: El producto %d no existe", id)
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&productos)
	if err != nil {
		return 0, err
	}
	if len(productos) == 0 {
		return 0, nil
	}

	return productos[len(productos)-1].Id, nil
}
