package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) service {
	return &service{repository: repository}
}
func (ser *service) GetAll() ([]Product, error) {
	products, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ser *service) Store(Id int, Name string, Color string, Price float64, Stock int, Code string, Publish bool, Date string) (Product, error) {
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
