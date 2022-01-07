package loader

type Service interface {
	CustomersLoader() error
	// InvoicersLoader() error
	// ProductsLoader() error
	// SalesLoader() error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CustomersLoader() error {
	return s.repo.CustomersLoader()
}
