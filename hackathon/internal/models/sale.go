package models

type User struct {
	ID       int     `json:"id"`
	Invoice  Invoice `json:"invoice"`
	Product  Product `json:"product"`
	Quantity float64 `json:"quantity"`
}
