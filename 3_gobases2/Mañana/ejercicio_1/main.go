package main

import (
	"fmt"
)

func impuesto(salario float64) (impuesto float64) {
	if salario > 150000 {
		impuesto = (0.1 * salario)

	} else if salario > 50000 {
		impuesto = (0.17 * salario)
	} else {
		impuesto = 0
	}
	return
}

func main() {
	salario := 155000.00
	fmt.Printf("Salario = %v \n", salario)
	fmt.Printf("Impuesto a pagar = %v\n", impuesto(salario))
}
