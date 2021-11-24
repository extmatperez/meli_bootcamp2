package main

import "fmt"

func main() {

	precio := 1000.0
	descuento := 20.0

	fmt.Printf("El precio con descuento es: %.2f\n", (precio*(100-descuento))/100)
}
