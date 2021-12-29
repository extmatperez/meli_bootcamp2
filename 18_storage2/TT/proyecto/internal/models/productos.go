package models

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
	Tipo          Tipo   `json:"tipo"`
}

type Tipo struct {
	ID          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}
