package main

import "fmt"

func main() {
	var valor float32 = 500000
	var descuento float32 = 35
	fmt.Printf("Total a Pagar: %v", valor-(valor*(descuento/100)))
}
