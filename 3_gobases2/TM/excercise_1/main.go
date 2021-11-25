package main

import "fmt"

func main() {

	var salario float64
	fmt.Println("ingresa el salario")
	fmt.Scanln(&salario)
	fmt.Println("el empleado debe pagar:", CalcularImpuesto(salario), "de impuesto")
}

func CalcularImpuesto(salario float64) float64 {
	var impuesto float64
	if salario > 50000 {
		if salario > 1500000 {
			impuesto = salario * 0.27
		} else {
			impuesto = salario * 0.17
		}
	} else {
		impuesto = 0
	}
	return impuesto
}
