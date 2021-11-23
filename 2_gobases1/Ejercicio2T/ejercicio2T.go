package main

import "fmt"

func main() {

	var precio float64
	var descuento float64

	precio = 250000
	descuento = 10
	totalDescuento := precio * (descuento / 100)
	precio -= totalDescuento
	fmt.Println("precio:", precio, "-- descuento:", descuento, "%")
	fmt.Println("total con descuento:", precio)

}
