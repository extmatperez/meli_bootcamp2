package internal

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]Producto, error)
	GetById(id int) (Producto, error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error)
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}
func (s *service) GetAll() ([]Producto, error) {
	return s.repository.GetAll()
}
func (s *service) GetById(id int) (Producto, error) {
	return s.repository.GetById(id)
}
func (s *service) Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fecha_creacion string) (Producto, error) {
	lastId, _ := s.repository.GetLastId()
	productToReturn, err := s.repository.Store(lastId+1, nombre, color, precio, stock, codigo, publicado, fecha_creacion)
	if err != nil {
		return Producto{}, err
	}
	return productToReturn, nil

}
