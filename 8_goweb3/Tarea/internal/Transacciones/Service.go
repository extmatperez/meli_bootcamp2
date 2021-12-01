package internal

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(Id int, codigo string, moneda string, monto float64, emisor string, receptor string) (Transaccion, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Transaccion, error) {
	transacciones, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	} else {
		return transacciones, nil
	}
}

func (s *service) Store(Id int, codigo string, moneda string, monto float64, emisor string, receptor string) (Transaccion, error) {
	lastID, err := s.repository.GetID()
	if err != nil {
		return Transaccion{}, err
	} else {
		lastID++

		transaccion, err := s.repository.Store(lastID+1, codigo, moneda, monto, emisor, receptor)

		if err != nil {
			return Transaccion{}, err
		}

		return transaccion, nil
	}
}
