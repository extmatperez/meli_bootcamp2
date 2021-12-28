package internal

import (
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/internal/models"
)

type ServiceSQL interface {
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error)
	GetByName(name string) (models.Producto, error)
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

func (ser *serviceSQL) GetByName(name string) (models.Producto, error) {
	return ser.repository.GetByName(name)
}
