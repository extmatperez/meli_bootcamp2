package internal

type Service interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	GetByReceiver(receiver string) (Transaction, error)
	//CreateTransaction(transaction Transaction) (Transaction, error)
	Store(transactionCode string, currency string, amount float64,
		receiver string, sender string, transactionDate string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) GetByID(id int) (Transaction, error) {

	tran, err := s.repository.GetByID(id)

	if err != nil {
		return Transaction{}, err
	}

	return tran, nil

}

func (s *service) GetByReceiver(receiver string) (Transaction, error) {

	tran, err := s.repository.GetByReceiver(receiver)

	if err != nil {
		return Transaction{}, err
	}

	return tran, nil

}

func (s *service) Store(transactionCode string, currency string, amount float64,
	receiver string, sender string, transactionDate string) (Transaction, error) {

	tran, err := s.repository.Store(transactionCode, currency, amount,
		receiver, sender, transactionDate)

	if err != nil {
		return Transaction{}, err
	}

	return tran, nil
}
