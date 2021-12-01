package internal

import "fmt"

type Service interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error)
	Update(id int, code string, currency string, amount float64, remitter string, receptor string, date string) (Transaction, error)
	Delete(id int) error
	ModifyTransactionCode(id int, code string) (Transaction, error)
	ModifyAmount(id int, amount float64) (Transaction, error)
}

type service struct {
	repository Repository
}

func (ser *service) ModifyAmount(id int, amount float64) (Transaction, error) {
	if amount > 0 {
		transaction, err := ser.repository.ModifyAmount(id, amount)
		if err != nil {
			return Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return Transaction{}, fmt.Errorf("El codigo de transaccion no puede estar vacio")
	}
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) ModifyTransactionCode(id int, transactionCode string) (Transaction, error) {
	if transactionCode != "" {
		transaction, err := ser.repository.ModifyTransactionCode(id, transactionCode)
		if err != nil {
			return Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return Transaction{}, fmt.Errorf("El codigo de transaccion no puede estar vacio")
	}
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

func (ser *service) Update(id int, code, currency string, amount float64, remitter, receptor, date string) (Transaction, error) {
	if validateFields(id, code, currency, amount, remitter, receptor, date) {
		transaction, err := ser.repository.Update(id, code, currency, amount, remitter, receptor, date)
		if err != nil {
			return Transaction{}, err
		} else {
			return transaction, nil
		}
	} else {
		return Transaction{}, fmt.Errorf("Alguno de los campos son incorrectos")
	}
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}

func validateFields(id int, code string, currency string, amount float64, remitter string, receptor string, date string) bool {
	if id <= 0 || amount <= 0 {
		return false
	}
	if code == "" || currency == "" || remitter == "" || receptor == "" || date == "" {
		return false
	}
	return true
}

