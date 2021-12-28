package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/17_storage1/internal/models"
)

type ServiceSQL interface {
	Store(nombre, tipo string, count int, price float64) (models.Producto, error)
	GetOne(id int) models.Producto
	Update(ctx context.Context, producto models.Producto) (models.Producto, error)
	GetAll() ([]models.Producto, error)
	Delete(id int) error
	GetFullData() ([]models.DTOProducto, error)
	//GetOneWithContext(ctx context.Context, id int) (models.Producto, error)
	//Store2(persona models.Producto) (models.Producto, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(nombre, tipo string, count int, price float64) (models.Producto, error) {

	newProducto := models.Producto{Name: nombre, Type: tipo, Count: count, Price: price}
	personaCreada, err := ser.repository.Store(newProducto)

	if err != nil {
		return models.Producto{}, err
	}
	return personaCreada, nil
}

func (ser *serviceSQL) GetOne(id int) models.Producto {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(ctx context.Context, producto models.Producto) (models.Producto, error) {
	return ser.repository.Update(ctx, producto)
}

func (ser *serviceSQL) GetAll() ([]models.Producto, error) {
	return ser.repository.GetAll()
}

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *serviceSQL) GetFullData() ([]models.DTOProducto, error) {
	return ser.repository.GetFullData()
}
