package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/GoStoragePrac/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error)
	GetOne(id int) models.Producto
	GetAll() ([]models.Producto, error)
	Update(producto models.Producto) (models.Producto, error)
	Delete(id int) error
	GetByName(nombre string) []models.Producto
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (models.Producto, error) {

	nuevoProducto := models.Producto{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreacion: fechaCreacion}

	productoCreado, err := ser.repository.Store(nuevoProducto)

	if err != nil {
		return models.Producto{}, err
	}
	return productoCreado, nil
}

func (ser *serviceSQL) GetOne(id int) models.Producto {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) GetAll() ([]models.Producto, error) {
	return ser.repository.GetAll()
}

func (ser *serviceSQL) Update(producto models.Producto) (models.Producto, error) {
	return ser.repository.Update(producto)
}

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *serviceSQL) GetByName(nombre string) []models.Producto {
	return ser.repository.GetByName(nombre)
}
