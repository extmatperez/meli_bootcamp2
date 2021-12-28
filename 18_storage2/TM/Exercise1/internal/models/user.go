package models

type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
}

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
	Address     City   `json:"address"`
}
