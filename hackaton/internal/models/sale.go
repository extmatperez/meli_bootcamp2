package models

type Sale struct {
	Id        int `json:"id"`
	IdProduct int `json:"idProduct"`
	IdInvoice int `json:"idInvoice"`
	Quantity  int `json:"quantity"`
}
