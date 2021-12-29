package models

// Estructura de los datos que voy a manipular
type City struct {
	ID          int    `json:"id"`
	CityName    string `json:"city_name"`
	CountryName string `json:"country_name"`
}

type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
	Address   City   `json:"address"`
}
