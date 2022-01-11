package product

import (
	"context"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	Save(ctx context.Context, prod domain.Product) (int, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Update(ctx context.Context, p domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, productCode string) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
func (ser *service) Save(ctx context.Context, prod domain.Product) (int, error) {
	id, err := ser.repository.Save(ctx, prod)
	if err != nil {
		return 0, err
	}
	return id, err
}
func (ser *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	prod, err := ser.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return prod, nil
}
func (ser *service) Update(ctx context.Context, p domain.Product) (domain.Product, error) {
	err := ser.repository.Update(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (ser *service) Delete(ctx context.Context, id int) error {

	err := ser.repository.Delete(ctx, id)
	return err

}
func (ser *service) Get(ctx context.Context, id int) (domain.Product, error) {
	prod, err := ser.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (ser *service) Exists(ctx context.Context, productCode string) bool {
	return ser.repository.Exists(ctx, productCode)
}
