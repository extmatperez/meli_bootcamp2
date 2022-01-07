package models

type Invoicer struct {
	ID          int     `json:"id"`
	Date_Time   string  `json:"date_time"`
	ID_Customer int     `json:"id_customer"`
	Total       float64 `json:"total"`
}

type RequestInvoicer struct {
	Date_Time   string  `json:"date_time"`
	ID_Customer int     `json:"id_customer"`
	Total       float64 `json:"total"`
}
