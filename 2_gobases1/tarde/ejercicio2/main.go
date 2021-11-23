package main

import "fmt"

func main() {
	var precio, precioFianl, descuento float64

	precio = 1000
	descuento = 10
	precioFianl = precio - precio*(descuento/100)
	fmt.Println("precio :", precio, "descuento :", descuento, "precio con descuento :", precioFianl)
}
