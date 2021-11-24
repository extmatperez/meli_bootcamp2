package main

import "fmt"

func calcularImpuesto(salario float64) float64 {
	impuesto := 0.0
	if salario >= 150000 {
		impuesto = 27
	} else if salario >= 50000 {
		impuesto = 17
	}
	return salario * impuesto / 100
}

func main() {
	fmt.Println(calcularImpuesto(78000))
}
