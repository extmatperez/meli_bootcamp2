package models

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type PersonaGetAllDTO struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Domicilio Ciudad `json:"ciudad"`
}

type Ciudad struct {
	Ciudad string `json:"ciudad"`
	Pais   string `json:"pais"`
}
