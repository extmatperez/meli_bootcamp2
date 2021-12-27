package internal

type Service interface {
	GetAll() ([]Payment, error)
	Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error)
	Update(id int, codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error)
	UpdateCodigo(id int, codigo string) (Payment, error)
	UpdateMonto(id int, monto float64) (Payment, error)
	Delete(id int) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *service) Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	ultimoId, err := s.repository.LastId()

	if err != nil {
		return Payment{}, err
	}

	pay, err := s.repository.Store(ultimoId+1, codigo, moneda, monto, emisor, receptor, fecha)

	if err != nil {
		return Payment{}, err
	}

	return pay, nil
}

func (s *service) Update(id int, codigo, moneda, emisor, receptor, fecha string, monto float64) (Payment, error) {
	return s.repository.Update(id, codigo, moneda, monto, emisor, receptor, fecha)
}

func (s *service) UpdateCodigo(id int, codigo string) (Payment, error) {
	return s.repository.UpdateCodigo(id, codigo)
}

func (s *service) UpdateMonto(id int, monto float64) (Payment, error) {
	return s.repository.UpdateMonto(id, monto)
}

func (s *service) Delete(id int) (string, error) {
	return s.repository.Delete(id)
}
