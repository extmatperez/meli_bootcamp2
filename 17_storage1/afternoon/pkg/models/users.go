package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Active   bool    `json:"active"`
	Created  string  `json:"created"`
}
