package loader

type Service interface {
	CustomersLoader() error
	InvoicersLoader() error
	ProductsLoader() error
	SalesLoader() error
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

func (s *service) InvoicersLoader() error {
	return s.repo.InvoicersLoader()
}

func (s *service) ProductsLoader() error {
	return s.repo.ProductsLoader()
}

func (s *service) SalesLoader() error {
	return s.repo.SalesLoader()
}
