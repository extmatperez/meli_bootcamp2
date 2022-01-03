package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/internal/models"
)

type ServiceSQL interface {
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error)
	GetByName(name string) ([]models.Producto, error)
	GetAll(ctx context.Context) ([]models.Producto, error)
	Update(ctx context.Context, producto models.Producto) (models.Producto, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error) {

	newPersona := models.Producto{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}
	personaCreada, err := ser.repository.Store(newPersona)

	if err != nil {
		return models.Producto{}, err
	}
	return personaCreada, nil
}

func (ser *serviceSQL) GetByName(name string) ([]models.Producto, error) {
	return ser.repository.GetByName(name)
}

func (ser *serviceSQL) GetAll(ctx context.Context) ([]models.Producto, error) {
	return ser.repository.GetAll(ctx)
}
func (ser *serviceSQL) Update(ctx context.Context, producto models.Producto) (models.Producto, error) {
	return ser.repository.Update(ctx, producto)
}
