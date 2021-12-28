package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	UpdateName(id int, nombre string) (Producto, error)
	Delete(id int) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (ser *service) GetAll() ([]Producto, error) {
	return ser.repo.GetAll()
}

func (ser *service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	ultimoId, err := ser.repo.LastId()

	if err != nil {
		return Producto{}, err
	}

	return ser.repo.Store(ultimoId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (ser *service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	return ser.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (ser *service) UpdateName(id int, nombre string) (Producto, error) {
	return ser.repo.UpdateName(id, nombre)
}
func (ser *service) Delete(id int) (string, error) {
	return ser.repo.Delete(id)
}
