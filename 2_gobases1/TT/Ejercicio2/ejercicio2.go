package main

import "fmt"

func main() {
	precio_producto := 2599.00
	porcentaje := 10.00
	fmt.Println(descuento(precio_producto, porcentaje))
}

func descuento(precio, porcentaje float64) float64 {
	precio_final := precio - precio*(porcentaje/100)
	return precio_final
}
