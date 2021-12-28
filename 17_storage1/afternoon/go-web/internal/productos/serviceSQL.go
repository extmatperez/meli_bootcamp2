package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/afternoon/go-web/internal/models"
)

type ServiceSQL interface {
	Store(product models.Producto) (models.Producto, error)
	GetByName(name string) ([]models.Producto, error)
}

type serviceSQL struct {
	repo RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repo: repo}
}

func (ser *serviceSQL) Store(product models.Producto) (models.Producto, error) {
	return ser.repo.Store(product)
}

func (ser *serviceSQL) GetByName(name string) ([]models.Producto, error) {
	return ser.repo.GetByName(name)
}
