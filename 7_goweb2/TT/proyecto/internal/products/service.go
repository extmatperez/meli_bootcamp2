package internal

type Service interface {
	GetAll() ([]Product, error)
	Save(newProduct Product) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Save(newProduct Product) (Product, error) {
	prod, err := s.repository.Save(newProduct)
	if err != nil {
		return Product{}, err
	}
	return prod, nil
}
