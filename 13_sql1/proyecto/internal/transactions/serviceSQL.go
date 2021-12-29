package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/13_sql1/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	UpdateWithContext(ctx context.Context, id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error)
	Delete(id int) error
	GetTransactionByID(id int) (models.Transaction, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	transaccionCreada, err := ser.repository.Store(codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaccionCreada, nil
}

func (ser *serviceSQL) Update(id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	transaccionActualizada, err := ser.repository.Update(id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaccionActualizada, nil
}

func (ser *serviceSQL) UpdateWithContext(ctx context.Context, id int, codigo_de_transaccion, moneda string, monto float64, emisor, receptor, fecha_de_transaccion string) (models.Transaction, error) {
	transaccionActualizada, err := ser.repository.UpdateWithContext(ctx, id, codigo_de_transaccion, moneda, monto, emisor, receptor, fecha_de_transaccion)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaccionActualizada, nil
}

func (ser *serviceSQL) Delete(id int) error {
	err := ser.repository.Delete(id)
	return err
}

func (ser *serviceSQL) GetTransactionByID(id int) (models.Transaction, error) {
	tr, err := ser.repository.GetTransactionByID(id)
	return tr, err
}
