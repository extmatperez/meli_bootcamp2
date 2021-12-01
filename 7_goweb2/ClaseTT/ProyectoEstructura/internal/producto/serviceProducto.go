package internal

type ServiceProducto interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
}

type serviceProducto struct {
	repository RepositoryProducto
}

func NewService(repository RepositoryProducto) ServiceProducto {
	return &serviceProducto{repository: repository}
}

func (ser *serviceProducto) GetAll() ([]Producto, error) {
	productos, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (ser *serviceProducto) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Producto{}, err
	}

	per, err := ser.repository.Store(ultimoId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return per, nil
}
