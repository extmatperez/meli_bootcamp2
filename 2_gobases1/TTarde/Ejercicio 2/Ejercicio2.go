package main

import "fmt"

func main() {
	var precio float64
	var descuento float64
	var d_final float64

	fmt.Println("Introduzca el precio y el descuento:")
	fmt.Scanln(&precio, &descuento)

	d_final = precio/100*descuento
	p_final := precio - d_final

	fmt.Println(p_final)

}