package main

import "fmt"

func main()  {
	
	var precio float32
	var descuento uint
	var precioFinal float32

	precio = 800
	descuento = 25

	precioFinal = precio - precio * float32(descuento) / 100.0
	fmt.Println("El precio final es", precioFinal)
}