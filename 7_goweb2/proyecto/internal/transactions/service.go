package internal

type Service interface {
	GetAll() ([]Transaction, error)
	GetTransactionByID(id int) (Transaction, error)
	Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetAll() ([]Transaction, error) {
	resultado, err := s.repository.GetAll()

	return resultado, err //TODO: Manejar errores
}

func (s *service) GetTransactionByID(id int) (Transaction, error) {
	resultado, err := s.repository.GetTransactionByID(id)
	return resultado, err //TODO: Manejar errores
}

func (s *service) Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {
	id := s.repository.LastId() + 1
	resultado, err := s.repository.Store(id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	return resultado, err // TODO: Manejar errores
}
