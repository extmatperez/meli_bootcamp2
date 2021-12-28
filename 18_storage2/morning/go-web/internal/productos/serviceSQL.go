package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/morning/go-web/internal/models"
)

type ServiceSQL interface {
	Store(product models.Producto) (models.Producto, error)
	GetByName(name string) ([]models.Producto, error)
	GetAll() ([]models.Producto, error)
	Update(ctx context.Context, producto models.Producto, id int) (models.Producto, error)
}

type serviceSQL struct {
	repo RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repo: repo}
}

func (ser *serviceSQL) Store(producto models.Producto) (models.Producto, error) {
	return ser.repo.Store(producto)
}

func (ser *serviceSQL) GetByName(name string) ([]models.Producto, error) {
	return ser.repo.GetByName(name)
}

func (ser *serviceSQL) GetAll() ([]models.Producto, error) {
	return ser.repo.GetAll()
}

func (ser *serviceSQL) Update(ctx context.Context, producto models.Producto, id int) (models.Producto, error) {
	return ser.repo.Update(ctx, producto, id)
}
