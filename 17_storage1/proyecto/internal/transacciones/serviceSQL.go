package internal

import (
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(CodigoTransaccion, Emisor, Receptor, Moneda string, Monto float64) (models.Transaccion, error)
	GetOne(id int) models.Transaccion
	Update(trans models.Transaccion) (models.Transaccion, error)
	GetByCode(CodigoTransaccion string) ([]models.Transaccion, error)
	//GetAll() ([]models.Transaccion, error)
}

type serviceSql struct {
	repository RepositorySql
}

func NewServiceSql(repo RepositorySql) ServiceSQL {
	return &serviceSql{repository: repo}
}

func (ser *serviceSql) Store(CodigoTransaccion, Emisor, Receptor, Moneda string, Monto float64) (models.Transaccion, error) {
	newTrans := models.Transaccion{CodigoTransaccion: CodigoTransaccion, Emisor: Emisor, Receptor: Receptor, Moneda: Moneda, Monto: Monto}
	transCreated, err := ser.repository.Store(newTrans)
	if err != nil {
		return models.Transaccion{}, nil
	}
	return transCreated, nil
}

func (ser *serviceSql) GetOne(id int) models.Transaccion {
	return ser.repository.GetOne(id)
}

func (ser *serviceSql) Update(trans models.Transaccion) (models.Transaccion, error) {
	return ser.repository.Update(trans)
}

func (ser *serviceSql) GetByCode(CodigoTransaccion string) ([]models.Transaccion, error) {
	trans, err := ser.repository.GetByCode(CodigoTransaccion)
	if err != nil {
		return nil, err
	}

	return trans, nil
}

/* func (ser *serviceSql) GetAll() ([]models.Usuario, error) {
	usuario, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return usuario, nil
} */
