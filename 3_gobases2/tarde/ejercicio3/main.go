package main

import (
	"fmt"
)

type tienda struct {
	producto producto
}
type producto struct {
	Nombre string
	Tipo   string
}
type Producto interface {
	calcularCosto() float64
}
type Ecommerce interface {
	total() float64
	agregar() float64
}

func nuevoProducto(nombre, tipo string) producto {
	producto := producto{Nombre: nombre, Tipo: tipo}
	return producto
}

func nuevaTienda() {
	nuevProd := nuevoProducto("pantalon", "ropa")
	tienda := tienda{producto: nuevProd}
	fmt.Println(tienda)
}
func main() {
	nuevaTienda()

}
