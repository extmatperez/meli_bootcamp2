package main

import "fmt"

func main() {
	var initial_price, disc_porc float64
	fmt.Println("Precio: ")
	fmt.Scanf("%f", &initial_price)
	fmt.Println("Descuento: ")
	fmt.Scanf("%f", &disc_porc)
	final_price := initial_price - (initial_price * (disc_porc / 100))
	fmt.Printf("\nPrecio final: %v pesos.", final_price)
}
