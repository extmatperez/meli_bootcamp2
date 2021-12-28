package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(nombre string, color string, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error)
	Get(id int) (models.Producto, error)
	GetByName(nombre string) ([]models.Producto, error)
	Update(producto models.Producto) (models.Producto, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func newServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(nombre string, color string, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error) {

	nuevoProducto := models.Producto{Nombre: nombre, Color: color, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}

	personaInsertada, err := s.repository.Store(nuevoProducto)

	if err != nil {
		return models.Producto{}, err
	}

	return personaInsertada, nil
}

func (s *serviceSQL) Get(id int) (models.Producto, error) {
	return s.repository.Get(id)
}

func (s *serviceSQL) Update(producto models.Producto) (models.Producto, error){
	return s.repository.Update(producto)
}

func (s *serviceSQL) GetByName(nombre string) ([]models.Producto, error){
	return s.repository.GetByName(nombre)
}
