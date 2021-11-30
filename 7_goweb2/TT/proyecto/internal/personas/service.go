package internal

type Service interface {
	GetAll() ([]Persona, error)
	Store(nombre, apellido string, edad int) (Persona, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll() ([]Persona, error) {
	personas, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (s *service) Store(nombre, apellido string, edad int) (Persona, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Persona{}, err
	}
	id++
	pers, err := s.repository.Store(id, nombre, apellido, edad)
	if err != nil {
		return Persona{}, nil
	}
	return pers, nil
}
