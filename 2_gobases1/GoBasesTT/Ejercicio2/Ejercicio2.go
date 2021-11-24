package main

import "fmt"

func main() {
	var precio float64 //= 100.50
	var descuento int = 20

	// & apunta a la ubicacion en memoria de precio
	fmt.Scanf("%f", &precio)

	calcularDescuento(precio, descuento)
}

func calcularDescuento(precio float64, porcentajeDescuento int) {

	fmt.Println(precio - (precio * float64(porcentajeDescuento) / 100))
}
