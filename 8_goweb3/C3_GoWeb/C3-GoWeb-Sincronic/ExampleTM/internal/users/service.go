package internal

type Service interface {
	GetAll() ([]User, error)
	Store(first_name string, last_name string, age int) (User, error)
	Update(id int, first_name string, last_name string, age int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]User, error) {
	user, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ser *service) Store(first_name string, last_name string, age int) (User, error) {
	ultimoID, err := ser.repository.LastId()

	if err != nil {
		return User{}, err
	}

	us, err := ser.repository.Store(ultimoID+1, first_name, last_name, age)
	if err != nil {
		return User{}, err
	}

	return us, nil
}

func (ser *service) Update(id int, first_name string, last_name string, age int) (User, error) {
	return ser.repository.Update(id, first_name, last_name, age)
}
