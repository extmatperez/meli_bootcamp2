package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

type InvoiceService interface {
	ImportAllInvoices() error
	StoreInvoice(datetime string, id_costumer int, total float64) (models.Invoice, error)
	UpdateInvoice(invoice models.Invoice) (models.Invoice, error)
}

type service_invoice struct {
	repository_invoice InvoiceRepository
}

func NewInvoiceService(repo InvoiceRepository) InvoiceService {
	return &service_invoice{repository_invoice: repo}
}

func (s *service_invoice) ImportAllInvoices() error {
	return s.repository_invoice.ImportAllInvoices()
}

func (s *service_invoice) StoreInvoice(datetime string, id_costumer int, total float64) (models.Invoice, error) {

	new_invoice := models.Invoice{Id: 0,
		Datetime:   datetime,
		IdCustomer: id_costumer,
		Total:      total,
	}
	c, err := s.repository_invoice.StoreInvoice(new_invoice)

	if err != nil {
		return models.Invoice{}, err
	}

	return c, nil
}

func (s *service_invoice) UpdateInvoice(invoice models.Invoice) (models.Invoice, error) {
	return s.repository_invoice.UpdateInvoice(invoice)
}
