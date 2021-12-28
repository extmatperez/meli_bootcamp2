package internal

import (
	"github.com/extmatperez/meli_bootcamp2/18_storage2/TTarde/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(CodigoTransaccion int, Moneda string, Monto float64, Emisor, Receptor, FechaTransaccion string) (models.Transaccion, error)
	GetByName(name string) models.Transaccion
	GetAll() []models.Transaccion
	Update(transaccion models.Transaccion)(models.Transaccion, error)
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


func (ser *serviceSQL) GetAll() []models.Transaccion {
	return ser.repository.GetAll()
}

func (ser *serviceSQL) Update(transaccion models.Transaccion) (models.Transaccion, error) {
	return ser.repository.Update(transaccion)
}
