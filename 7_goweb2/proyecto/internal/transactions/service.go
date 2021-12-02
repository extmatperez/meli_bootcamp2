package internal

import "strconv"

type Service interface {
	GetAll(filters map[string]string) ([]Transaction, error)
	GetTransactionByID(id int) (Transaction, error)
	Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error)
	Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error)
	UpdateCodigoYMonto(id int, codigo_de_transaccion string, monto float64) (Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetAll(filters map[string]string) ([]Transaction, error) {
	transacciones, err := s.repository.GetAll()

	for key, val := range filters {
		var resultado []Transaction
		if key == "Codigo" {
			for _, t := range transacciones {
				if t.CodigoDeTransaccion == val {
					resultado = append(resultado, t)
				}
			}
		} else if key == "Moneda" {
			for _, t := range transacciones {
				if t.Moneda == val {
					resultado = append(resultado, t)
				}
			}
		} else if key == "Monto" {
			for _, t := range transacciones {
				if m, _ := strconv.ParseFloat(val, 64); t.Monto == m {
					resultado = append(resultado, t)
				}
			}
		} else if key == "Emisor" {
			for _, t := range transacciones {
				if t.Emisor == val {
					resultado = append(resultado, t)
				}
			}
		} else if key == "Receptor" {
			for _, t := range transacciones {
				if t.Receptor == val {
					resultado = append(resultado, t)
				}
			}
		}
		//TODO fecha desde y fecha hasta
		transacciones = resultado
	}

	return transacciones, err
}

func (s *service) GetTransactionByID(id int) (Transaction, error) {
	resultado, err := s.repository.GetTransactionByID(id)
	return resultado, err
}

func (s *service) Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Transaction{}, err
	}
	id++
	resultado, err := s.repository.Store(id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	return resultado, err
}

func (s *service) Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (Transaction, error) {
	resultado, err := s.repository.Update(id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	return resultado, err
}

func (s *service) UpdateCodigoYMonto(id int, codigo_de_transaccion string, monto float64) (Transaction, error) {
	resultado, err := s.repository.UpdateCodigoYMonto(id, codigo_de_transaccion, monto)
	return resultado, err
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	return err
}
