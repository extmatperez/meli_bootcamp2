package internal

type Service interface {
	Load() ([]Transaccion, error)
	GetAll() ([]Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	FindById(id int) (Transaccion, error)
	FilterBy(valores ...string) ([]Transaccion, error)
	Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	UpdateCod(id int, codigotransaccion string) (Transaccion, error)
	UpdateMon(id int, monto float64) (Transaccion, error)
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

func (s *service) Load() ([]Transaccion, error) {
	transaccion, err := s.repository.Load()

	if err != nil {
		return nil, err
	}

	return transaccion, nil
}

func (s *service) GetAll() ([]Transaccion, error) {
	transaccion, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return transaccion, nil
}

func (s *service) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {

	if len(transacciones) == 0 {
		id = 1
	} else {
		id = transacciones[len(transacciones)-1].ID + 1
	}
	transaccion, err := s.repository.Store(id, codigotransaccion, moneda, monto, emisor, receptor, fechacreacion)
	if err != nil {
		return Transaccion{}, err
	}
	return transaccion, nil
}

func (s *service) FindById(id int) (Transaccion, error) {
	return s.repository.FindById(id)
}

func (s *service) FilterBy(valores ...string) ([]Transaccion, error) {
	return s.repository.FilterBy(valores...)
}

func (s *service) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {
	return s.repository.Update(id, codigotransaccion, moneda, monto, emisor, receptor, fechacreacion)
}

func (s *service) UpdateCod(id int, codigotransaccion string) (Transaccion, error) {
	return s.repository.UpdateCod(id, codigotransaccion)
}
func (s *service) UpdateMon(id int, monto float64) (Transaccion, error) {
	return s.repository.UpdateMon(id, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

/* func (s *service) DeleteAll() error {
	return s.repository.DeleteAll()
} */
