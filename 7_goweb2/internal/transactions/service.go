package internal

type Service interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Transaction, error){
	transactions, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	} else {
		return transactions, nil
	}
}

func (ser *service) Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	newID, err := ser.repository.LastID()
	if err != nil {
		return Transaction{}, err
	} else {
		transaction, err := ser.repository.Store(newID+1, code, currency, amount, remitter, receptor, date)
		if err != nil {
			return Transaction{}, err
		} else {
			return transaction, nil
		}
	}
}