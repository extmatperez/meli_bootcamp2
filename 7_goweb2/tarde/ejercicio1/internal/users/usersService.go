package internal

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
func (ser *service) GetAll() ([]User, error) {
	products, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (ser *service) Store(nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	lasid, err := ser.repository.LastID()
	if err != nil {
		return User{}, err
	}
	p, err := ser.repository.Store(lasid+1, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return User{}, err
	}
	return p, nil
}
