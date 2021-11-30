package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	id++
	prod, err := s.repository.Store(id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion)
	if err != nil {
		return Product{}, nil
	}
	return prod, nil
}
