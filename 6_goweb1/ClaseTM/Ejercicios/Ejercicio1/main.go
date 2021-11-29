package main

type Fecha struct {
	Dia  int
	Mes  int
	Anio int
}
type Producto struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion Fecha   `json:"fecha_creacion"`
}

func main() {

}
