package main

import "fmt"

func descuento(precio float64, porcentaje float64) {
	var descuento float64 = precio * porcentaje / 100
	var precioFinal float64 = precio - descuento
	fmt.Printf("El precio con los descuentos es de %v\n", precioFinal)
}

func main() {
	descuento(700, 10)
}
