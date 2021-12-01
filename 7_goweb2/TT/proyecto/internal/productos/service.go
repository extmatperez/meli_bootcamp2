/*
Servicio, debe contener la lógica de nuestra aplicación.
a. Se debe crear el archivo service.go.
b. Se debe generar la interface Service con todos sus métodos.
c. Se debe generar la estructura service que contenga el repositorio.
d. Se debe generar una función que devuelva el Servicio.
e. Se deben implementar todos los métodos correspondientes a las operaciones
a realizar (GetAll, Store, etc..).
*/

package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Producto, error) {

	productos, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return productos, nil
}

func (s *service) Store(nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	lastId, err := s.repository.LastId()

	if err != nil {
		return Producto{}, err
	}

	nuevoProducto, err := s.repository.Store(lastId + 1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}
