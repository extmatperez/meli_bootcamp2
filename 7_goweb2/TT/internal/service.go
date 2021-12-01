package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, price float64, stock int) (Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (serv *service) GetAll() ([]Product, error) {
	products, err := serv.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (serv *service) Store(name string, price float64, stock int) (Product, error) {
	lastID, _ := serv.repo.LastID()

	lastID++

	product, err := serv.repo.Store(lastID, name, price, stock)
	if err != nil {
		return product, err
	}

	return product, nil
}
