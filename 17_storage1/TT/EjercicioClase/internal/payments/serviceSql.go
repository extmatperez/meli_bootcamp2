package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/models"
)

type ServiceSql interface {
	Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (models.Payment, error)
	GetById(id int) models.Payment
	GetByCode(codigo string) models.Payment
	GetAllPayments() ([]models.Payment, error)
	Update(payment models.Payment) (models.Payment, error)
	Delete(id int) error
	GetFullDataAllPayments() ([]models.Payment, error)
	GetByIdWithContext(ctx context.Context, id int) (models.Payment, error)
	UpdateWithContext(ctx context.Context, payment models.Payment) (models.Payment, error)
}

type serviceSql struct {
	repositorySql RepositorySql
}

func NewServiceSql(repo RepositorySql) ServiceSql {
	return &serviceSql{repositorySql: repo}
}

func (s *serviceSql) Store(codigo, moneda, emisor, receptor, fecha string, monto float64) (models.Payment, error) {

	new_pay := models.Payment{Id: 0,
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    fecha,
	}
	pay, err := s.repositorySql.Store(new_pay)

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

func (s *serviceSql) GetAllPayments() ([]models.Payment, error) {
	return s.repositorySql.GetAllPayments()
}

func (s *serviceSql) Update(payment models.Payment) (models.Payment, error) {
	return s.repositorySql.Update(payment)
}

func (s *serviceSql) Delete(id int) error {
	return s.repositorySql.Delete(id)
}

func (s *serviceSql) GetFullDataAllPayments() ([]models.Payment, error) {
	return s.repositorySql.GetFullDataAllPayments()
}

func (s *serviceSql) GetByIdWithContext(ctx context.Context, id int) (models.Payment, error) {
	return s.repositorySql.GetByIdWithContext(ctx, id)
}

func (s *serviceSql) UpdateWithContext(ctx context.Context, payment models.Payment) (models.Payment, error) {
	return s.repositorySql.UpdateWithContext(ctx, payment)
}
