/*
Servicio, debe contener la lógica de nuestra aplicación.
- Se debe crear el archivo service.go.
- Se debe generar la interface Service con todos sus métodos.
- Se debe generar la estructura service que contenga el repositorio.
- Se debe generar una función que devuelva el Servicio.
- Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
*/

package internal

type Service interface {
	GetAll() ([]User, error)
	Store(first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error)
	Update(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error)
	Delete(id int) error
	UpdateLastName(id int, last_name string) (User, error)
	UpdateAge(id int, age int) (User, error)
}

type service struct {
	repository Repository
}

func (ser *service) GetAll() ([]User, error) {
	user, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ser *service) Store(first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error) {
	ultimoID, err := ser.repository.LastId()

	if err != nil {
		return User{}, err
	}

	us, err := ser.repository.Store(ultimoID+1, first_name, last_name, email, age, height, active, create_date)
	if err != nil {
		return User{}, err
	}

	return us, nil
}

func (ser *service) Update(id int, first_name string, last_name string, email string, age int, height int, active bool, create_date string) (User, error) {
	return ser.repository.Update(id, first_name, last_name, email, age, height, active, create_date)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *service) UpdateLastName(id int, last_name string) (User, error) {
	return ser.repository.UpdateLastName(id, last_name)
}

func (ser *service) UpdateAge(id int, age int) (User, error) {
	return ser.repository.UpdateAge(id, age)
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
