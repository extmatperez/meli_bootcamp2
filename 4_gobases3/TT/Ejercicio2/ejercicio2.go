package main

func main() {

}

type Usuario struct {
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Correo    string     `json:"correo"`
	Productos []Producto `json:"producto"`
}

type Producto struct {
	Nombre   string `json:"nombre"`
	Precio   int    `json:"precio"`
	Cantidad int    `json:"cantidad"`
}

func nuevo_producto(nombre string, precio int) Producto {

}
