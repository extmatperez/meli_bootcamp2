package internal

import "time"

type Service interface {
	GetAll() ([]Users, error)
	Store(firstName, lastName string, email string, age int, height float64) (Users, error)
	Update(id int, firstName, lastName string, email string, age int, height float64, active bool, creationDate time.Time) (Users, error)
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

func (ser *service) Store(firstName, lastName string, email string, edad int, height float64) (Users, error) {
	var (
		active       = true
		creationDate = time.Now()
	)
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Users{}, err
	}

	per, err := ser.repository.Store(ultimoId+1, firstName, lastName, email, edad, height, active, creationDate)

	if err != nil {
		return Users{}, err
	}
	return per, nil
}

func (ser *service) Update(id int, firstName, lastName string, email string, age int, height float64, active bool, creationDate time.Time) (Users, error) {
	return ser.repository.Update(id, firstName, lastName, email, age, height, active, creationDate)
}

func (ser *service) UpdateName(id int, firstName string) (Users, error) {
	return ser.repository.UpdateName(id, firstName)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
