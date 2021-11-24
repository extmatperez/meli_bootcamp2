package main

import "fmt"

func main() {

	fmt.Printf("Ejercicio 2\n")
	fmt.Println("Ingrese precio del producto a comprar")
	var precio int
	fmt.Scanf("%d", &precio)
	var descuento float64 = 0.3
	var precioFinal = float64(precio) - (float64(precio) * descuento)

	fmt.Println("el precio del articulo es", precio, "el descuento es del", descuento*100, "por ciento y en total te queda en", precioFinal)

}
