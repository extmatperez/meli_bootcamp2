package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(color string, price float64, amount int) (Product, error)
	Update(id int, color string, price float64, amount int) (Product, error)
	UpdatePrice(id int, price float64) (Product, error)
	Delete(id int) error
	Sum(prices ...float64) float64
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Product, error) {
	products, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ser *service) Store(color string, price float64, amount int) (Product, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	prod, err := ser.repository.Store(ultimoId+1, color, price, amount)

	if err != nil {
		return Product{}, err
	}

	return prod, nil
}

func (ser *service) Update(id int, color string, price float64, amount int) (Product, error) {
	return ser.repository.Update(id, color, price, amount)
}

func (ser *service) UpdatePrice(id int, price float64) (Product, error) {
	return ser.repository.UpdatePrice(id, price)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (s *service) Sum(prices ...float64) float64 {
	var price float64
	for _, p := range prices {
		price += p
	}
	return price
}
