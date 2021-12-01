package internal

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	Search(id string) (Transaccion, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (serv *service) GetAll() ([]Transaccion, error) {
	transacciones, err := serv.repository.getAll()
	if err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (serv *service) Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	lastID, err := serv.repository.LastId()
	if err != nil {
		return Transaccion{}, err
	}

	transaccion, err := serv.repository.Store(lastID+1, codTransaccion, moneda, monto, emisor, receptor, fechaTrans)
	if err != nil {
		return Transaccion{}, err
	}

	return transaccion, err
}

func (serv *service) Search(id string) (Transaccion, error) {
	var transac Transaccion

}
