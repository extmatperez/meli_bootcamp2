package main

import "fmt"

func descuento() {
	var precio float64
	var porcentaje float64

	fmt.Printf("Ingrese el precio del producto: ")
	fmt.Scanf("%f", &precio)

	fmt.Printf("Ingrese el descuento: ")
	fmt.Scanf("%f", &porcentaje)

	var descuento float64 = precio * porcentaje / 100
	var precioFinal float64 = precio - descuento

	fmt.Printf("El precio con los descuentos es de %v\n", precioFinal)
}

func main() {
	descuento()
}
