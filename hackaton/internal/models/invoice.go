package models

type Invoice struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	IdCustomer int     `json:"idCustomer"`
	Total      float64 `json:"total"`
}
