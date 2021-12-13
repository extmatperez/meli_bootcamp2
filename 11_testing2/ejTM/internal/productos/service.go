package internal

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]Producto, error)
	GetById(id int) (Producto, error)
	Store(nombre, color string, precio float64) (Producto, error)
	Update(id int, nombre, color string, precio float64) (Producto, error)
	UpdateNombrePrecio(id int, nombre string, precio float64) (Producto, error)
	Delete(id int) error
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
func (s *service) Store(nombre, color string, precio float64) (Producto, error) {
	lastId, _ := s.repository.GetLastId()
	productToReturn, err := s.repository.Store(lastId+1, nombre, color, precio)
	if err != nil {
		return Producto{}, err
	}
	return productToReturn, nil

}

func (s *service) Update(id int, nombre, color string, precio float64) (Producto, error) {

	product, err := s.repository.Update(id, nombre, color, precio)
	if err != nil {
		return Producto{}, err
	}
	return product, nil
}

func (s *service) UpdateNombrePrecio(id int, nombre string, precio float64) (Producto, error) {

	product, err := s.repository.UpdateNombrePrecio(id, nombre, precio)
	if err != nil {
		return Producto{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
