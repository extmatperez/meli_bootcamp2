package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

type SaleService interface {
	ImportAllSales() error
	StoreSale(id_product, id_invoice, quantity int) (models.Sale, error)
	UpdateSale(sale models.Sale) (models.Sale, error)
}

type service_sale struct {
	repository_sale SaleRepository
}

func NewSaleService(repo SaleRepository) SaleService {
	return &service_sale{repository_sale: repo}
}

func (s *service_sale) ImportAllSales() error {
	return s.repository_sale.ImportAllSales()
}

func (s *service_sale) StoreSale(id_product, id_invoice, quantity int) (models.Sale, error) {

	new_sale := models.Sale{Id: 0,
		IdProduct: id_product,
		IdInvoice: id_invoice,
		Quantity:  quantity,
	}
	c, err := s.repository_sale.StoreSale(new_sale)

	if err != nil {
		return models.Sale{}, err
	}

	return c, nil
}

func (s *service_sale) UpdateSale(sale models.Sale) (models.Sale, error) {
	return s.repository_sale.UpdateSale(sale)
}
