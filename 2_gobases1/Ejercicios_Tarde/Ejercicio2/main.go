package main

import "fmt"

func div(precio float64, descuento float64) float64 {
	return precio - ((precio * descuento) / 100)
}

func main() {
	result := div(100, 30)
	fmt.Println(result)
}
