package internal

import (
	"context"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/models"
)

type ServiceSQL interface {
	GetByName(name string) (models.Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
	GetAll() ([]models.Product, error)
	Update(ctx context.Context, id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
	FindById(id int) (models.Product, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio int) (models.Product, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(r RepositorySQL) *serviceSQL {
	return &serviceSQL{repository: r}
}

func (s *serviceSQL) GetByName(name string) (models.Product, error) {
	prod, err := s.repository.GetByName(name)
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (s *serviceSQL) Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error) {
	prod, err := s.repository.Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (s *serviceSQL) GetAll() ([]models.Product, error) {
	prods, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return prods, nil
}
func (s *serviceSQL) Update(ctx context.Context, id int, nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error) {
	prod, err := s.repository.Update(ctx, id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return prod, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil
}
func (s *serviceSQL) FindById(id int) (models.Product, error) {
	prod, err := s.repository.GetOne(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("No se encontro el producto, %w", err)
	}
	return prod, nil

}
func (s *serviceSQL) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *serviceSQL) UpdateNameAndPrice(id int, nombre string, precio int) (models.Product, error) {
	prod, err := s.repository.UpdateNameAndPrice(id, nombre, precio)

	if err != nil {
		return prod, fmt.Errorf("No se pudo actualizar el producto, %w", err)
	}

	return prod, nil
}
