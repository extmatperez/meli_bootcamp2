package main

import "fmt"

func main() {

	var descuento float32 = 25.5
	var precio float32 = 1000.00
	descuentoAplicado := (precio * descuento) / 100
	precioConDescuento := precio - descuentoAplicado
	fmt.Println("El precio con el descuento aplicado es de ", precioConDescuento)

}
