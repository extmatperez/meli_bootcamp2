package models

type Product struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"precio"`
}
type RequestProduct struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
