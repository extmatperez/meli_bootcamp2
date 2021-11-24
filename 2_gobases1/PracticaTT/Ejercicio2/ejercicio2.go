package main

import "fmt"

func main() {
	var precio, descuento float32
	fmt.Printf("Ingrese el precio:\n")
	fmt.Scanf("%f", &precio)
	fmt.Printf("Ingrese el descuento:\n")
	fmt.Scanf("%f", &descuento)

	precioConDto := precio - precio*descuento/100

	fmt.Printf("El precio original de la prenda es $%v\n", precio)
	fmt.Printf("El precio con un descuento del %v%% es de $%v\n", descuento, precioConDto)
}
