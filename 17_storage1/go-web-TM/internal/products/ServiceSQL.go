package internal

import models "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/17_storage1/go-web-TM/internal/models"

/*
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
*/
type ServiceSQL interface {
	Store(nombre, color string, precio int, stock, codigo string, publicado bool, FechaCreacion string) (models.Product, error)
	GetOne(id int) models.Product
	Update(producto models.Product) (models.Product, error)
	GetByName(nombre string) models.Product
}
type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (s *serviceSQL) Store(nombre, color string, precio int, stock, codigo string, publicado bool, FechaCreacion string) (models.Product, error) {
	newProduct := models.Product{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: FechaCreacion}
	productCreated, err := s.repository.Store(newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return productCreated, nil
}
func (s *serviceSQL) GetOne(id int) models.Product {
	return s.repository.GetOne(id)
}

func (ser *serviceSQL) Update(producto models.Product) (models.Product, error) {
	return ser.repository.Update(producto)
}
func (ser *serviceSQL) GetByName(nombre string) models.Product {
	return ser.repository.GetByName(nombre)
}
