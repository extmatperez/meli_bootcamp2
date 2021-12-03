package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
	Update(id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price float64) (Product, error)
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

func (ser *service) Store(Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
	ultimoId, err := ser.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	per, err := ser.repository.Store(ultimoId+1, Name, Color, Price, Stock, Code, Publish, Date)

	if err != nil {
		return Product{}, err
	}
	return per, err
}
func (ser *service) Update(id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {

	return ser.repository.Update(id, Name, Color, Price, Stock, Code, Publish, Date)
}
func (ser *service) Delete(id int) error {

	return ser.repository.Delete(id)
}

func (ser *service) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	return ser.repository.UpdateNamePrice(id, name, price)
}
