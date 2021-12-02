package internal

type Service interface {
	GetAll() ([]User, error)
	Store(nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error)
	UpdateNombre(id int, nombre string) (User, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
func (ser *service) GetAll() ([]User, error) {
	users, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (ser *service) Update(id int, nombre, apellido, email string, edad, altura int, activo, fechaCreacion string) (User, error) {
	return ser.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
}

func (ser *service) UpdateNombre(id int, nombre string) (User, error) {
	return ser.repository.UpdateNombre(id, nombre)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
