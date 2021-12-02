package internal

type Service interface {
	GetAll() ([]Users, error)
	Store(firstName, lastName string, age int) (Users, error)
	Update(id int, firstName, lastName string, age int) (Users, error)
	UpdateName(id int, firstName string) (Users, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Users, error) {
	personas, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (ser *service) Store(firstName, lastName string, edad int) (Users, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Users{}, err
	}

	per, err := ser.repository.Store(ultimoId+1, firstName, lastName, edad)

	if err != nil {
		return Users{}, err
	}
	return per, nil
}

func (ser *service) Update(id int, firstName, lastName string, age int) (Users, error) {
	return ser.repository.Update(id, firstName, lastName, age)
}

func (ser *service) UpdateName(id int, firstName string) (Users, error) {
	return ser.repository.UpdateName(id, firstName)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
