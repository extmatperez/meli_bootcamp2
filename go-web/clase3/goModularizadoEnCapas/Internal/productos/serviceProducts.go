package internal

type Service interface {
	GetAll() ([]Product, error)
	Store(productoAux Product) (Product, error)
	Update(varID int, productoAux Product) (Product, error)
	UpdateName(varID int, nameUpdate string) error
	Delete(varID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Product, error) {
	personas, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (ser *service) Store(productoAux Product) (Product, error) {
	ultimoId, err := ser.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	productoAux.Id = ultimoId + 1
	producto, err := ser.repository.Store(productoAux)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (ser *service) Update(varID int, producto Product) (Product, error) {
	producto, err := ser.repository.Update(varID, producto)
	if err != nil {
		return Product{}, err
	} else {
		return producto, nil
	}
}

func (ser *service) UpdateName(varID int, nameUpdate string) error {
	err := ser.repository.UpdateName(varID, nameUpdate)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (ser *service) Delete(varID int) error {
	err := ser.repository.Delete(varID)
	if err != nil {
		return err
	} else {
		return nil
	}
}
