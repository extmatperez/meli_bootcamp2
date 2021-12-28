package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTT/internal/models"
)

type ServiceSQL interface {
	Store(moneda string, monto float64, emisor, receptor string) (models.Transaccion, error)
	GetOne(id int) models.Transaccion
	Update(transaccion models.Transaccion) (models.Transaccion, error)
	GetByName(name string) []models.Transaccion
	GetAll() ([]models.Transaccion, error)
	Delete(id int) error
	GetFullData() ([]models.Transaccion, error)

	GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(moneda string, monto float64, emisor, receptor string) (models.Transaccion, error) {
	newTransaccion := models.Transaccion{Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor}
	transaccionCreada, err := s.repository.Store(newTransaccion)

	if err != nil {
		return models.Transaccion{}, err
	}

	return transaccionCreada, nil
}

func (s *serviceSQL) GetOne(id int) models.Transaccion {
	return s.repository.GetOne(id)
}

func (s *serviceSQL) Update(transaccion models.Transaccion) (models.Transaccion, error) {
	return s.repository.Update(transaccion)
}

func (s *serviceSQL) GetByName(name string) []models.Transaccion {
	return s.repository.GetByName(name)
}

func (s *serviceSQL) GetAll() ([]models.Transaccion, error) {
	return s.repository.GetAll()
}

func (s *serviceSQL) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *serviceSQL) GetFullData() ([]models.Transaccion, error) {
	return s.repository.GetFullData()
}

func (s *serviceSQL) GetOneWithContext(ctx context.Context, id int) (models.Transaccion, error) {
	return s.repository.GetOneWithContext(ctx, id)
}
