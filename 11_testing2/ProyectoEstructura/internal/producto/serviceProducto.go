package internal

import "fmt"

type ServiceProducto interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, nombre string, precio float64) (Producto, error)
}

type serviceProducto struct {
	repository RepositoryProducto
}

func NewService(repository RepositoryProducto) ServiceProducto {
	return &serviceProducto{repository: repository}
}

func (ser *serviceProducto) GetAll() ([]Producto, error) {
	productos, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (ser *serviceProducto) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Producto{}, err
	}

	per, err := ser.repository.Store(ultimoId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return per, nil
}

func (ser *serviceProducto) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	return ser.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (ser *serviceProducto) Delete(id int) error {
	if id > 0 {
		return ser.repository.Delete(id)
	} else {
		return fmt.Errorf("No existe id negativos. Input: %d", id)
	}
}

func (ser *serviceProducto) UpdateNameAndPrice(id int, nombre string, precio float64) (Producto, error) {
	if nombre != "" && precio != 0.0 {
		return ser.repository.UpdateNameAndPrice(id, nombre, precio)
	} else {
		return Producto{}, fmt.Errorf("Error: Nombre no puede ser vacio y precio no puede ser 0. Datos ingresados: nombre %s y precio %f", nombre, precio)
	}
}
