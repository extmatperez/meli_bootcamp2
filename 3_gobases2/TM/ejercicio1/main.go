package main

import "fmt"

func main() {
	salario := 200000
	fmt.Printf("El monto a descontar es: %.2f", CalcularImpuesto(salario))
}

func CalcularDescuento(sueldo int) int {
	switch {
	case sueldo <= 50000:
		return 0
	case sueldo <= 150000:
		return 17
	default:
		return 27
	}
}

func CalcularImpuesto(sueldo int) float64 {
	descuento := CalcularDescuento(sueldo)
	return float64(sueldo) * float64(descuento) / 100.0
}
