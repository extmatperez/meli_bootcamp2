package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"

type ServiceSql interface {
	Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (models.Payment, error)
	GetById(id int) models.Payment
	GetByCode(codigo string) models.Payment
	GetAllPayments() []models.Payment
	Update(payment models.Payment) (models.Payment, error)
}

type serviceSql struct {
	repositorySql RepositorySql
}

func NewServiceSql(repo RepositorySql) ServiceSql {
	return &serviceSql{repositorySql: repo}
}

func (s *serviceSql) Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (models.Payment, error) {

	pay, err := s.repositorySql.Store(models.Payment{0, codigo, moneda, monto, emisor, receptor, fecha})

	if err != nil {
		return models.Payment{}, err
	}

	return pay, nil
}

func (s *serviceSql) GetById(id int) models.Payment {
	return s.repositorySql.GetById(id)
}

func (s *serviceSql) GetByCode(codigo string) models.Payment {
	return s.repositorySql.GetByCode(codigo)
}

func (s *serviceSql) GetAllPayments() []models.Payment {
	return s.repositorySql.GetAllPayments()
}

func (s *serviceSql) Update(payment models.Payment) (models.Payment, error) {
	return s.repositorySql.Update(payment)
}
