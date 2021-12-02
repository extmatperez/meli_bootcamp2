package internal

import (
	"fmt"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	FindById(id int) (Product, error)
	Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio int) (Product, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (serv *service) GetAll() ([]Product, error) {
	prods, err := serv.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("No se puedo extraer los datos %w", err)
	} else {
		if prods != nil {
			return prods, nil
		} else {
			return []Product{}, nil
		}
	}
}

func (serv *service) Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {

	id, err := serv.repo.LastID()
	if err != nil {
		return Product{}, fmt.Errorf("No se pudo crear el producto %w", err)
	}

	prod, err := serv.repo.Store(id+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Product{}, fmt.Errorf("No se pudo crear el producto %w", err)
	}

	return prod, nil

}

func (serv *service) FindById(id int) (Product, error) {
	prods, err := serv.repo.GetAll()
	if err != nil {
		return Product{}, fmt.Errorf("No se puedo extraer los datos %w", err)
	}

	for _, value := range prods {
		if value.ID == id {
			return value, nil
		}
	}

	return Product{}, fmt.Errorf("Producto id: %d no encontrado", id)

}

func (serv *service) Update(id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	prod, err := serv.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return prod, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil
}

func (serv *service) Delete(id int) error {
	err := serv.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("No se pudo eliminar el producto, %w", err)
	}
	return nil
}

func (serv *service) UpdateNameAndPrice(id int, nombre string, precio int) (Product, error) {

	prod, err := serv.repo.UpdateNameAndPrice(id, nombre, precio)

	if err != nil {
		return Product{}, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil

}
