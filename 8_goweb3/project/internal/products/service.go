package internal

type Service interface {
	GetAll() ([]Products, error)
	Store(name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error)
	Update(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]Products, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error) {

	lastID, err := s.repository.LastID()

	if err != nil {
		return Products{}, err
	}

	lastID++

	prod, err := s.repository.Store(lastID, name, color, price, stock, code,
		published, creationdate)

	if err != nil {
		return Products{}, err
	}

	return prod, nil
}

func (s *service) Update(id int64, name string, color string, price string, stock string, code string, published bool, creationdate string) (Products, error) {
	return s.repository.Update(id, name, color, price, stock, code,
		published, creationdate)
}
