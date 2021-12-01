package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
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
