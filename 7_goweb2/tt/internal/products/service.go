package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(ID int, Name, Color string, Price float64, Stock, Code int, isPublished bool, CreatedAt string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Product, error) {
	prods, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func (s *service) Store(id int, name, color string, price float64, stock, code int, isPublished bool, createdAt string) (Product, error) {
	//generar id
	prod, err := s.repository.Store(id, name, color, price, stock, code, isPublished, createdAt)
	if err != nil {
		return Product{}, err
	}
	return prod, nil
}
