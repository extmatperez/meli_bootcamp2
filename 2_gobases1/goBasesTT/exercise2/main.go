package main

import "fmt"

func main() {

	fmt.Printf("Ejercicio 2\n")

	var precio int = 1000
	var descuento int = 30
	var precioFinal = precio - (precio * (descuento / 100))

	fmt.Printf("el precio del articulo es %d el descuento es del %d por ciento y en total te queda en %d", precio, descuento, precioFinal)

}
