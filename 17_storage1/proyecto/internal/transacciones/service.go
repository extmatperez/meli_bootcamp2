package internal

import "github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"

type Service interface {
	Load() ([]models.Transaccion, error)
	GetAll() ([]models.Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error)
	FindById(id int) (models.Transaccion, error)
	FilterBy(valores ...string) ([]models.Transaccion, error)
	Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error)
	UpdateCod(id int, codigotransaccion string) (models.Transaccion, error)
	UpdateMon(id int, monto float64) (models.Transaccion, error)
	Delete(id int) error
	//DeleteAll() error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Load() ([]models.Transaccion, error) {
	trans, err := s.repository.Load()

	if err != nil {
		return nil, err
	}

	return trans, nil
}

func (s *service) GetAll() ([]models.Transaccion, error) {
	trans, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return trans, nil
}

func (s *service) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {

	if len(transacciones) == 0 || id == 0 {

		id = 1
	} else {
		id = transacciones[len(transacciones)-1].ID + 1
	}
	trans, err := s.repository.Store(id, codigotransaccion, moneda, monto, emisor, receptor, fechacreacion)
	if err != nil {
		return models.Transaccion{}, err
	}
	return trans, nil
}

func (s *service) FindById(id int) (models.Transaccion, error) {
	return s.repository.FindById(id)
}

func (s *service) FilterBy(valores ...string) ([]models.Transaccion, error) {
	return s.repository.FilterBy(valores...)
}

func (s *service) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {
	return s.repository.Update(id, codigotransaccion, moneda, monto, emisor, receptor, fechacreacion)
}

func (s *service) UpdateCod(id int, codigotransaccion string) (models.Transaccion, error) {
	return s.repository.UpdateCod(id, codigotransaccion)
}
func (s *service) UpdateMon(id int, monto float64) (models.Transaccion, error) {
	return s.repository.UpdateMon(id, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

/* func (s *service) DeleteAll() error {
	return s.repository.DeleteAll()
} */
