package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	Delete(id int) error
	Change(id int, nombre, precio string) (Producto, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Producto, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Producto{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, color, precio, stock, codigo, publicado, creado)
	if err != nil {
		return Producto{}, err
	}

	return producto, nil
}

func (s *service) Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	productoEditado, err := s.repository.Edit(id, nombre, color, precio, stock, codigo, publicado, creado)
	if err != nil {
		return Producto{}, err
	}
	return productoEditado, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Change(id int, nombre, precio string) (Producto, error) {
	cambios, err := s.repository.Change(id, nombre, precio)
	if err != nil {
		return Producto{}, err
	}
	return cambios, err
}
