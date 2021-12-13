package internal

type Service interface {
	GetAll() ([]Transaction, error)
	Store(transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error)
	Update(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error)
	UpdateReceptor(id int, receptor string) (Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Transaction, error) {
	transaction, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (ser *service) Store(transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Transaction{}, err
	}

	trans, err := ser.repository.Store(ultimoId+1, transaction_code, coin, emitor, receptor, transaction_date, amount)

	if err != nil {
		return Transaction{}, err
	}

	return trans, nil
}

func (ser *service) Update(id int, transaction_code, coin, emitor, receptor, transaction_date string, amount float64) (Transaction, error) {
	return ser.repository.Update(id, transaction_code, coin, emitor, receptor, transaction_date, amount)
}

func (ser *service) UpdateReceptor(id int, nombre string) (Transaction, error) {
	return ser.repository.UpdateReceptor(id, nombre)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
