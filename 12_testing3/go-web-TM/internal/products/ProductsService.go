package internal

import "fmt"

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
	UpdateNombre(id int, nombre string) (Product, error)
	UpdatePrecio(id int, precio int) (Product, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
func (ser *service) GetAll() ([]Product, error) {
	products, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (ser *service) Store(nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	lasid, err := ser.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	p, err := ser.repository.Store(lasid+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return Product{}, err
	}
	fmt.Println("en service")
	fmt.Println(p)
	fmt.Println("fuera service")
	return p, nil
}
func (ser *service) Update(id int, nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {

	p, err := ser.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
func (ser *service) UpdateNombre(id int, nombre string) (Product, error) {
	p, err := ser.repository.UpdateNombre(id, nombre)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
func (ser *service) UpdatePrecio(id int, precio int) (Product, error) {
	p, err := ser.repository.UpdatePrecio(id, precio)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
