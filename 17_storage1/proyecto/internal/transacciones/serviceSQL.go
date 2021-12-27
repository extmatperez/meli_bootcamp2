package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(CodigoTransaccion int, Moneda string, Monto float64, Emisor, Receptor, FechaTransaccion string) (models.Transaccion, error)
	GetByName(name string) models.Transaccion
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) GetByName(name string) models.Transaccion {
	return ser.repository.GetByName(name)
}

func (ser *serviceSQL) Store(CodigoTransaccion int, Moneda string, Monto float64, Emisor, Receptor, FechaTransaccion string) (models.Transaccion, error) {
	newTransaccion := models.Transaccion{CodigoTransaccion: CodigoTransaccion, Moneda: Moneda, Monto: Monto, Emisor: Emisor, Receptor: Receptor, FechaTransaccion: FechaTransaccion}

	transaccionCreada, err := ser.repository.Store(newTransaccion)

	if err != nil {
		return models.Transaccion{}, err
	}
	return transaccionCreada, nil
}