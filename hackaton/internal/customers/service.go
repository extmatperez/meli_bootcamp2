package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

type CustomerService interface {
	ImportAllCustomers() error
	StoreCustomer(lastName, firstName, condition string) (models.Customer, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
}

type service_customer struct {
	repository_customer CustomerRepository
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &service_customer{repository_customer: repo}
}

func (s *service_customer) ImportAllCustomers() error {
	return s.repository_customer.ImportAllCustomers()
}

func (s *service_customer) StoreCustomer(lastName, firstName, condition string) (models.Customer, error) {

	new_customer := models.Customer{Id: 0,
		LastName:  lastName,
		FirstName: firstName,
		Condition: condition,
	}
	c, err := s.repository_customer.StoreCustomer(new_customer)

	if err != nil {
		return models.Customer{}, err
	}

	return c, nil
}

func (s *service_customer) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	return s.repository_customer.UpdateCustomer(customer)
}
