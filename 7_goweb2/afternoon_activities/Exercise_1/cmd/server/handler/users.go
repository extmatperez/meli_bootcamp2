// Agregamos el package handler
package handler

// importo el package handler
import (
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/7_goweb2/afternoon_activities/Exercise_1/internal/users"
)

// Creamos la request struct
type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

// Creamos Users struct
type Users struct {
	service users.Service
}
