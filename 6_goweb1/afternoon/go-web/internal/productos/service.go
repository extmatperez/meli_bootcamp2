package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (ser *service) GetAll() ([]Producto, error) {
	personas, err := ser.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (ser *service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	ultimoId, err := ser.repo.LastId()

	if err != nil {
		return Producto{}, err
	}

	per, err := ser.repo.Store(ultimoId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return per, nil
}
