package productos

type Service interface {
	GetAll() ([]Producto, error)
	GetOne(id int) (Producto, error)
	AddOne(nombre, color string, precio float64, stock int, codigo string, publicado bool) (Producto, error)
}

type productoService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &productoService{
		repo: r,
	}
}

func (s *productoService) GetAll() ([]Producto, error) {

	return s.repo.GetAll()
}

func (s *productoService) GetOne(id int) (Producto, error) {

	return s.repo.Get(id)
}
func (s *productoService) AddOne(
	nombre, color string,
	precio float64,
	stock int,
	codigo string,
	publicado bool) (Producto, error) {

	return s.repo.Store(nombre, color, precio, stock, codigo, publicado)
}
