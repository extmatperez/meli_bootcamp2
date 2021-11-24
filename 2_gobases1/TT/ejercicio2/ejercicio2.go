package main

import "fmt"

func main() {
	var price float32
	var discount float32 = 25

	fmt.Printf("\n Ingrese el valor del producto:")
	fmt.Scanf("%f", &price)
	fmt.Printf("Precio con descuento: %.2f", price*(1-discount))
}
