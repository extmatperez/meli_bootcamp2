package internal

import "errors"

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	Search(id string) (Transaccion, error)
	Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error)
	Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
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
	return serv.repository.Search(id)
}

func (serv *service) Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error) {
	filtredTransac, err := serv.repository.Filter(mapEtiquetas, mapRelacionEtiquetas)
	if err != nil {
		return []Transaccion{}, err
	}
	if len(filtredTransac) == 0 {
		err = errors.New("no se encontró ninguna transaccion para el filtro indicado")
	} else {
		err = nil
	}
	return filtredTransac, err
}

func (serv *service) Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	return serv.repository.Update(id, codTransaccion, moneda, monto, emisor, receptor, fechaTrans)
}
