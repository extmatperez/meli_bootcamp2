package sales

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
)

type Service interface {
	GetAll() ([]models.Sales, error)
	GetByID(id int) (models.Sales, error)
	Store(id_invoice, id_product int, quantity float64) (models.Sales, error)
	Update(ctx context.Context, id int, id_invoice, id_product int, quantity float64) (models.Sales, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (ser *service) GetAll() ([]models.Sales, error) {
	return ser.repo.GetAll()
}

func (ser *service) GetByID(id int) (models.Sales, error) {
	return ser.repo.GetByID(id)
}

func (ser *service) Store(id_invoice, id_product int, quantity float64) (models.Sales, error) {
	return ser.repo.Store(id_invoice, id_product, quantity)
}

func (ser *service) Update(ctx context.Context, id int, id_invoice, id_product int, quantity float64) (models.Sales, error) {
	return ser.repo.Update(ctx, id, id_invoice, id_product, quantity)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repo.Delete(ctx, id)
}
