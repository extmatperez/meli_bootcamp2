package internal

import "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/internal/models"

type ServiceSQL interface {
	Store(nombre, color string, precio float64) (models.Producto, error)
	GetAll() ([]models.Producto, error)
	GetById(id int) (models.Producto, error)
	// GetLastId() (int, error)
	Update(models.Producto) (models.Producto, error)
	// UpdateNombrePrecio(id int, nombre string, precio float64) (models.Producto, error)
	Delete(id int) error
}
type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(nombre, color string, precio float64) (models.Producto, error) {
	nuevoProducto := models.Producto{Nombre: nombre, Color: color, Precio: precio}
	productoCreado, err := s.repository.Store(nuevoProducto)
	if err != nil {
		return models.Producto{}, err
	}
	return productoCreado, nil

}
func (s *serviceSQL) GetById(id int) (models.Producto, error) {
	producto, err := s.repository.GetById(id)
	if err != nil {
		return models.Producto{}, err
	}
	return producto, nil
}

func (s *serviceSQL) GetAll() ([]models.Producto, error) {
	productos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s *serviceSQL) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *serviceSQL) Update(producto models.Producto) (models.Producto, error) {
	productoActualizado, err := s.repository.Update(producto)
	if err != nil {
		return models.Producto{}, err
	}
	return productoActualizado, nil
}
