package internal

import "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"

var sales []models.Sale

type SaleRepository interface {
	ImportAllSales() error
	StoreSale(models.Sale) (models.Sale, error)
	UpdateSale(models.Sale) (models.Sale, error)
}

type repository_sale struct{}

func NewSaleRepository() SaleRepository {
	return &repository_sale{}
}

func (r *repository_sale) ImportAllSales() error {

}

func (r *repository_sale) StoreSale(sale models.Sale) (models.Sale, error) {

}

func (r *repository_sale) UpdateSale(sale models.Sale) (models.Sale, error) {

}
