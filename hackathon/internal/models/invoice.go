package models

type Invoice struct {
	ID       int      `json:"id"`
	Datetime string   `json:"datetime"`
	Customer Customer `json:"customer"`
	Total    float64  `json:"precio" `
}
