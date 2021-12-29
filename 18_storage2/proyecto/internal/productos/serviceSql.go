package internal

import "github.com/extmatperez/meli_bootcamp2/18_storage2/proyecto/internal/models"

type ServiceSql interface {
	Store(stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (models.Productos, error)
	GetByName(name string) models.Productos
}

type serviceSql struct {
	repository RepositorySql
}

func NewServiceSql(r RepositorySql) ServiceSql {
	return &serviceSql{repository: r}
}

func (s *serviceSql) Store(stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (models.Productos, error) {
	newProducto := models.Productos{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha_de_creacion: fecha_de_creacion}
	productoCreado, err := s.repository.Store(newProducto)
	if err != nil {
		return models.Productos{}, err
	}
	return productoCreado, nil
}

func (s *serviceSql) GetByName(name string) models.Productos {
	return s.repository.GetByName(name)
}
