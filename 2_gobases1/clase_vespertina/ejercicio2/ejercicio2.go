package main

import (
	"fmt"
)

func main() {
	var precio, descuento float64
	fmt.Println("Ingrese el precio:")
	fmt.Scanf("%f", &precio)
	fmt.Println("Ingrese el descuento (%):")
	fmt.Scanf("%f", &descuento)
	calcDescuento(precio, descuento)
}

func calcDescuento(precio, descuento float64) {
	descuentoTotal := (descuento * precio) / 100
	fmt.Println("Se realizo un descuento del", descuento, "%")
	fmt.Printf("Se desconto un total de $ %.2f", descuentoTotal)
	fmt.Printf("\n")
}
