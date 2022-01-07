package customers

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
)

type Service interface {
	GetAll() ([]models.Customer, error)
	GetByID(id int) (models.Customer, error)
	Store(first_name, last_name, condition string) (models.Customer, error)
	Update(ctx context.Context, id int, first_name, last_name, condition string) (models.Customer, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (ser *service) GetAll() ([]models.Customer, error) {
	return ser.repo.GetAll()
}

func (ser *service) GetByID(id int) (models.Customer, error) {
	return ser.repo.GetByID(id)
}

func (ser *service) Store(first_name, last_name, condition string) (models.Customer, error) {
	return ser.repo.Store(first_name, last_name, condition)
}

func (ser *service) Update(ctx context.Context, id int, first_name, last_name, condition string) (models.Customer, error) {
	return ser.repo.Update(ctx, id, first_name, last_name, condition)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repo.Delete(ctx, id)
}
