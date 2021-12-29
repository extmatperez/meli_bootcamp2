package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(nombre string, color string, stock int, precio string, codigo string, publicado bool, fechaCreacion string, idTipo int) (models.Producto, error)
	GetAll() ([]models.Producto, error)
	Get(id int) (models.Producto, error)
	GetByName(nombre string) ([]models.Producto, error)
	Update(producto models.Producto) (models.Producto, error)
	Delete(id int) error
	GetAllFullData() ([]models.Producto, error)
	GetWithContext(ctx context.Context, id int) (models.Producto, error)
	UpdateWithContext(ctx context.Context, producto models.Producto) (models.Producto, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(nombre string, color string, stock int, precio string, codigo string, publicado bool, fechaCreacion string, idTipo int) (models.Producto, error) {

	nuevoProducto := models.Producto{Nombre: nombre, Color: color, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion, Tipo: models.Tipo{ID: idTipo}}

	personaInsertada, err := s.repository.Store(nuevoProducto)

	if err != nil {
		return models.Producto{}, err
	}

	return personaInsertada, nil
}

func (s *serviceSQL) GetAll() ([]models.Producto, error) {
	return s.repository.GetAll()
}

func (s *serviceSQL) Get(id int) (models.Producto, error) {
	return s.repository.Get(id)
}

func (s *serviceSQL) Update(producto models.Producto) (models.Producto, error) {
	return s.repository.Update(producto)
}

func (s *serviceSQL) GetByName(nombre string) ([]models.Producto, error) {
	return s.repository.GetByName(nombre)
}

func (s *serviceSQL) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *serviceSQL) GetAllFullData() ([]models.Producto, error) {
	return s.repository.GetAllFullData()
}

func (s *serviceSQL) GetWithContext(ctx context.Context, id int) (models.Producto, error) {
	return s.repository.GetWithContext(ctx, id)
}

func (s *serviceSQL) UpdateWithContext(ctx context.Context, producto models.Producto) (models.Producto, error) {
	return s.repository.UpdateWithContext(ctx, producto)
}
