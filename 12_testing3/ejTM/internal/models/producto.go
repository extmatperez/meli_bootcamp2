package models

type Producto struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Color  string  `json:"color" `
	Precio float64 `json:"precio" `
	Ciudad Ciudad  `json:"ciudad" `
}

type Ciudad struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre" `
}
