package internal

import "github.com/extmatperez/meli_bootcamp2/17_storage1/internal/models"

type ServiceDB interface {
	Store(amount float64, transactionCode, currency, receiver, sender, transactionDate string) (models.Transaction, error)
	GetOne(id int) models.Transaction
	Update(transaction models.Transaction) (models.Transaction, error)
	GetBySender(sender string) (models.Transaction, error)
}

type serviceDB struct {
	repository RepositoryDB
}

func NewServiceDB(repo RepositoryDB) ServiceDB {
	return &serviceDB{repository: repo}
}

func (service *serviceDB) Store(amount float64, transactionCode, currency, receiver, sender, transactionDate string) (models.Transaction, error) {
	transactionCreated, err := service.repository.Store(
		models.Transaction{
			TransactionCode: transactionCode,
			Currency:        currency,
			Amount:          amount,
			Receiver:        receiver,
			Sender:          sender,
			TransactionDate: transactionDate,
		})

	if err != nil {
		return models.Transaction{}, err
	}
	return transactionCreated, nil
}

func (service *serviceDB) GetOne(id int) models.Transaction {
	return service.repository.GetOne(id)
}

func (service *serviceDB) Update(transaction models.Transaction) (models.Transaction, error) {
	return service.repository.Update(transaction)
}

func (service *serviceDB) GetBySender(sender string) (models.Transaction, error) {
	return service.repository.GetBySender(sender)
}
