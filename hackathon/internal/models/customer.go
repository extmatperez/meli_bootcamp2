package models

type Condition string

const (
	ACTIVO    = "Activo"
	INACTIVO  = "Inactivo"
	BLOQUEADO = "Bloqueado"
)

type Customer struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Condition Condition `json:"condition" `
}
