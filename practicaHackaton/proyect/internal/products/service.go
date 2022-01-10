package products

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
)

type Service interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (models.Product, error)
	Store(description string, price float64) (models.Product, error)
	Update(ctx context.Context, id int, description string, price float64) (models.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (ser *service) GetAll() ([]models.Product, error) {
	return ser.repo.GetAll()
}

func (ser *service) GetByID(id int) (models.Product, error) {
	return ser.repo.GetByID(id)
}

func (ser *service) Store(description string, price float64) (models.Product, error) {
	return ser.repo.Store(description, price)
}

func (ser *service) Update(ctx context.Context, id int, description string, price float64) (models.Product, error) {
	return ser.repo.Update(ctx, id, description, price)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repo.Delete(ctx, id)
}
