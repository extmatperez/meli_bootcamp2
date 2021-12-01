package internal

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(codigo int, moneda string, monto float64, emisor string, receptor string, fecha string)(Transaccion, error)
	Update(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string)(Transaccion, error)
	UpdateEmisor(id int, emisor string)(Transaccion, error)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service{
	return &service{repository: r}
}

func (s *service) GetAll() ([]Transaccion, error){
	ps, err := s.repository.GetAll()
	if err != nil{
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(codigo int, moneda string, monto float64, emisor string, receptor string, fecha string)(Transaccion, error){
	trID, err := s.repository.LastID()

	if err!=nil{
		return Transaccion{}, err
	}

	trID++

	transaccion, err := s.repository.Store(trID, codigo, moneda, monto, emisor, receptor, fecha)

	if err != nil {
		return Transaccion{}, nil
	}

	return transaccion, nil
}


func(s *service) Update(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string)(Transaccion, error){
	return s.repository.Update(id, codigo, moneda, monto, emisor, receptor, fecha)
}

func(s *service) UpdateEmisor(id int, emisor string)(Transaccion, error){
	return s.repository.UpdateEmisor(id, emisor)
}
