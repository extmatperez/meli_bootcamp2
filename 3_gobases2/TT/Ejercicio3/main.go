package main

import "fmt"

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar()
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Tienda struct {
	prod []producto
}

func nuevoProducto(tipo, nombre string, precio float64) (p producto) {
	p = producto{tipo, nombre, precio}
	return
}

// func nuevaTienda() (e Tienda){
// 	t := Tienda{}
// 	return
// }

func main() {
	fmt.Println(nuevoProducto(pequeño, "chocolate", 750))
}
