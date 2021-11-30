package internal

type Service interface {
	GetAll() ([]Transaction, error)
	Store(code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Transaction, error) {
	transactions, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ser *service) Store(code string, currency string, amount int, sender string, receiver string, date string) (Transaction, error) {
	lastId, err := ser.repository.LastId()

	if err != nil {
		return Transaction{}, err
	}

	transaction, err := ser.repository.Store(lastId+1, code, currency, amount, sender, receiver, date)

	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}
