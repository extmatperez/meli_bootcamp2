package internal

type Service interface {
	GetAll() ([]User, error)
	Store(name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	UpdateLastNameAge(id int, lastName string, age int) (User, error)
	Delete(id int) (bool, error)
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

func (ser *service) Store(name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	lastID, err := ser.repository.LastID()

	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := ser.repository.Store(lastID, name, lastName, email, age, height, active, created)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ser *service) Update(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	user, err := ser.repository.Update(id, name, lastName, email, age, height, active, created)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ser *service) UpdateLastNameAge(id int, lastName string, age int) (User, error) {
	user, err := ser.repository.UpdateLastNameAge(id, lastName, age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ser *service) Delete(id int) (bool, error) {
	couldDelete, err := ser.repository.Delete(id)
	if err != nil {
		return couldDelete, err
	}

	return couldDelete, nil
}
