package main

import "fmt"

func main() {
	var precio float64 = 45.5
	var descuento int = 25

	var total float64 = precio - (precio * float64(descuento) / 100)
	fmt.Printf("El precio final es: %.2f\n", total)
}