package main

import "fmt"

func main() {
	var salario float64 = 50001
	fmt.Println("El impuesto del empleado es: ", impuesto(salario))
}

func impuesto(salario float64) float64 {
	var descuento float64
	if salario > 50000 {
		descuento = salario * (27.0 / 100.0)
	} else {
		descuento = salario * (17.0 / 100.0)
	}
	return descuento
}
