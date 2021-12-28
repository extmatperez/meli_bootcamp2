package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Store(ctx context.Context, name string, price float64, description string) (domain.Product, error)
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) GetByName(ctx context.Context, name string) ([]domain.Product, error) {
	products, err := s.repository.GetByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(ctx context.Context, name string, price float64, description string) (domain.Product, error) {
	product := domain.Product{
		Name:        name,
		Price:       price,
		Description: description,
	}

	product, err := s.repository.Store(ctx, product)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
