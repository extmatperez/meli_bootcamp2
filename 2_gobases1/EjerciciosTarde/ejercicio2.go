package main

import "fmt"

func main() {
	var precio, descuento float64
	precio = 100
	descuento = 15

	fmt.Printf("Precio con descuento: %.2f \n", (precio - (precio * (descuento / 100))))

}
