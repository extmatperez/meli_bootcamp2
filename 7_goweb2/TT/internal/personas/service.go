// ETAPA 2: tiene todos los metodos que va a usar la capa superior del controller.
// Necesita tener internamente a un repository.
// El service realiza la logica del negocio.

package internal

type Service interface {
	GetAll() ([]Persona, error)
	Store(nombre, apellido string, edad int) (Persona, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Persona, error) {
	personas, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (s *service) Store(nombre, apellido string, edad int) (Persona, error) {
	ultimoId, err := s.repository.LastId()

	if err != nil {
		return Persona{}, err
	}

	per, err := s.repository.Store(ultimoId+1, nombre, apellido, edad)
	if err != nil {
		return Persona{}, err
	}

	return per, nil
}
