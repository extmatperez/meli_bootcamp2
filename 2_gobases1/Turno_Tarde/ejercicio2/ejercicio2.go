package main

import "fmt"

func main() {
	var price float64 = 100.50
	var descuento float64 = 20

	fmt.Println("Precio del Producto: ", price)
	fmt.Println("Descuento: ", descuento, "%")

	fmt.Println("Cantidad Descuento: ", (price * float64(descuento/100)))
}
