package internal

import "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/internal/models"

type ServiceSQL interface {
	Store(nombre, tipo string, cantidad int, precio float64) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(id int) (models.Product, error)
	Update(models.Product) (models.Product, error)
	Delete(id int) error
}
type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(nombre, tipo string, cantidad int, precio float64) (models.Product, error) {
	nuevoProduct := models.Product{Name: nombre, Type: tipo, Count: cantidad, Price: precio}
	ProductCreado, err := s.repository.Store(nuevoProduct)
	if err != nil {
		return models.Product{}, err
	}
	return ProductCreado, nil

}
func (s *serviceSQL) GetById(id int) (models.Product, error) {
	Product, err := s.repository.GetById(id)
	if err != nil {
		return models.Product{}, err
	}
	return Product, nil
}

func (s *serviceSQL) GetAll() ([]models.Product, error) {
	Products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (s *serviceSQL) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *serviceSQL) Update(Product models.Product) (models.Product, error) {
	ProductActualizado, err := s.repository.Update(Product)
	if err != nil {
		return models.Product{}, err
	}
	return ProductActualizado, nil
}
