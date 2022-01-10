package internal

type Service interface {
	GetAll() ([]Product, error)
	Save(name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error)
	Update(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error)
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

func (s *service) Save(name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error) {
	ultimoId, err := s.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	prod, err := s.repository.Save(ultimoId+1, name, color, price, stock, code, published, creationDate)

	if err != nil {
		return Product{}, err
	}
	return prod, nil
}

func (s *service) Update(id int, name, color string, price float32, stock int, code string, published bool, creationDate string) (Product, error) {
	prod, err := s.repository.Update(id, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}
	return prod, nil
}
