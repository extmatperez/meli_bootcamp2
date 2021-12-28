package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Store(ctx context.Context, name string, price float64, description string) (domain.Product, error)
	Update(ctx context.Context, id int, name string, price float64, description string) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	product, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.Get(ctx, id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
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

func (s *service) Update(ctx context.Context, id int, name string, price float64, description string) (domain.Product, error) {
	product := domain.Product{
		Id:          id,
		Name:        name,
		Price:       price,
		Description: description,
	}

	product, err := s.repository.Update(ctx, product)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
