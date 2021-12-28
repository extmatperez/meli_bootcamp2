package internal

import (
	"context"

	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/models"
)

type ServiceSQL interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) (models.Persona, error)
	GetByName(nombre string) []models.Persona
	GetAll() ([]models.Persona, error)
	GetFullDataJoined() ([]models.Persona, error)
	UpdateWithContext(ctx context.Context, persona models.Persona) (models.Persona, error)
	Delete(id int) error
	GetOneWithContext(ctx context.Context, id int) (models.Persona, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(persona models.Persona) (models.Persona, error) {
	per_creada, err := s.repository.Store(persona)
	if err != nil {
		return models.Persona{}, err
	}

	return per_creada, nil
}

func (s *serviceSQL) GetOne(id int) (models.Persona, error) {
	return s.repository.GetOne(id)
}

func (s *serviceSQL) GetByName(nombre string) []models.Persona {
	return s.repository.GetByName(nombre)
}

func (s *serviceSQL) UpdateWithContext(ctx context.Context, persona models.Persona) (models.Persona, error) {
	return s.repository.UpdateWithContext(ctx, persona)
}

func (s *serviceSQL) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *serviceSQL) GetAll() ([]models.Persona, error) {
	return s.repository.GetAll()
}

func (s *serviceSQL) GetOneWithContext(ctx context.Context, id int) (models.Persona, error) {
	return s.repository.GetOneWithContext(ctx, id)
}

func (s *serviceSQL) GetFullDataJoined() ([]models.Persona, error) {
	return s.repository.GetFullDataJoined()
}
