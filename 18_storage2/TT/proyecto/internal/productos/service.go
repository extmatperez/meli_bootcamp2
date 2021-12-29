package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error)
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

	nuevoProducto, err := s.repository.Store(lastId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}

func (s *service) Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error) {
	return s.repository.UpdateNombrePrecio(id, nombre, precio)
}
