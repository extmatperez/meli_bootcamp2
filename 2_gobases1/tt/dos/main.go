package main

import "fmt"

var precio float64 = 35.50
var descuento float64 = 10

func main() {
	total := precio - precio*descuento/100
	fmt.Println("El precio final es ", total, "con un descuento del ", descuento, "%")
}
