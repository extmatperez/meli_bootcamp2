package main 

import "fmt"

func main() {
	precio := 1300.0
	descuento := 25.0 //En porcentaje

	fmt.Println(precio * (1 - descuento/100))
}