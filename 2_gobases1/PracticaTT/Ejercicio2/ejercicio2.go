package main

import "fmt"

func main() {
	precio := 20.5
	descuento := 0.15

	precioConDto := precio - precio*descuento

	fmt.Printf("El precio original de la prenda es $%v\n", precio)
	fmt.Printf("El precio con un descuento del %v %% es de $%v\n", descuento*100, precioConDto)
}
