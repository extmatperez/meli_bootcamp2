package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error)
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
func (ser *service) Store(nombre, color string, precio int, stock, codigo string, publicado bool, fechaCreacion string) (Product, error) {
	lasid, err := ser.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	p, err := ser.repository.Store(lasid+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
