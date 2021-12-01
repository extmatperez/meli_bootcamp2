package internal

type Service interface {
	GetAll() ([]Productos, error)
	Store(stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Productos, error) {
	productos, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (ser *service) Store(stock int, nombre, color, codigo, fecha_de_creacion string, precio float64, publicado bool) (Productos, error) {
	ultimoId, err := ser.repository.LastId()
	if err != nil {
		return Productos{}, err
	}

	prod, err := ser.repository.Store(ultimoId+1, stock, nombre, color, codigo, fecha_de_creacion, precio, publicado)

	if err != nil {
		return Productos{}, err
	}
	return prod, nil
}
