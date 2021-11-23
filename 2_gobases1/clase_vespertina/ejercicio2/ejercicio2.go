package main

import (
	"fmt"
)

func main() {
	precio, descuento := 858.21, 18.0
	calcDescuento(precio, descuento)
}

func calcDescuento(precio, descuento float64) {
	descuentoTotal := (descuento * precio) / 100
	fmt.Println("Se realizo un descuento del", descuento, "%")
	fmt.Printf("Se desconto un total de $ %.2f", descuentoTotal)
	fmt.Printf("\n")
}
