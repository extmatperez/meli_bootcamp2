package internal

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	Load() ([]Transaccion, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Load() ([]Transaccion, error) {
	transaccion, err := s.repository.Load()

	if err != nil {
		return nil, err
	}

	return transaccion, err
}

func (s *service) GetAll() ([]Transaccion, error) {
	transaccion, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return transaccion, nil
}

func (s *service) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {

	if len(transacciones) == 0 {
		id = 1
	} else {
		id = transacciones[len(transacciones)-1].ID + 1
	}

	transaccion, err := s.repository.Store(id, codigotransaccion, moneda, monto, emisor, receptor, fechacreacion)

	if err != nil {
		return Transaccion{}, err
	}

	return transaccion, nil
}
