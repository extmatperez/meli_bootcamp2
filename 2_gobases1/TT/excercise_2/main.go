package main

import "fmt"

func main() {
	var precio float64
	var cantdescuento float64

	fmt.Println("ingrese el valor de producto:")
	fmt.Scanln(&precio)
	fmt.Println("ingrese el descuento:")
	fmt.Scanln(&cantdescuento)
	fmt.Println("El precio final con el descuento es de:")

	var descuento float64 = precio * (cantdescuento / 100)
	var precioFinal float64 = precio - descuento
	fmt.Println(precioFinal)
}
