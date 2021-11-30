package internal

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Usuario, error) {
	usuarios, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (s *service) Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	ultimoId, err := s.repository.LastId()
	if err != nil {
		return Usuario{}, err
	}

	usr, err := s.repository.Store(ultimoId+1, nombre, apellido, email, edad, altura, activo, fecha_creacion)
	if err != nil {
		return Usuario{}, err
	}
	return usr, nil
}
