package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, stock int, code string, published bool, created_at string) (Product, error)
	FindById(id int64) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(name string, color string, stock int, code string, published bool, created_at string) (Product, error) {
	newId, err := s.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	product, err := s.repository.Store(newId+1, name, color, stock, code, published, created_at)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) FindById(id int64) (Product, error) {
	product, err := s.repository.FindById(id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
