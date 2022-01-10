package models

type Sales struct {
	ID         int     `json:"id"`
	ID_Invoice int     `json:"id_invoice"`
	ID_Product int     `json:"id_product"`
	Quantity   float64 `json:"quantity"`
}
type RequestSales struct {
	ID_Invoice int     `json:"id_invoice"`
	ID_Product int     `json:"id_product"`
	Quantity   float64 `json:"quantity"`
}
