package internal

import "github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/models"

type ServiceSQL interface {
	GetByName(name string) (models.Product, error)
	Store(nombre string, color string, precio int, stock int, codigo string, publicado bool, fechaCreacion string) (models.Product, error)
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
