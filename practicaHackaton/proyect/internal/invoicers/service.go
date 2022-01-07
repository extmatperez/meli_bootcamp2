package invoicers

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
)

type Service interface {
	GetAll() ([]models.Invoicer, error)
	GetByID(id int) (models.Invoicer, error)
	Store(date_time string, id_customer int, total float64) (models.Invoicer, error)
	Update(ctx context.Context, id int, date_time string, id_customer int, total float64) (models.Invoicer, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (ser *service) GetAll() ([]models.Invoicer, error) {
	return ser.repo.GetAll()
}

func (ser *service) GetByID(id int) (models.Invoicer, error) {
	return ser.repo.GetByID(id)
}

func (ser *service) Store(date_time string, id_customer int, total float64) (models.Invoicer, error) {
	return ser.repo.Store(date_time, id_customer, total)
}

func (ser *service) Update(ctx context.Context, id int, date_time string, id_customer int, total float64) (models.Invoicer, error) {
	return ser.repo.Update(ctx, id, date_time, id_customer, total)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repo.Delete(ctx, id)
}
