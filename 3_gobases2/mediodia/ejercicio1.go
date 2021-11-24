package main

import "fmt"

func calcImpuesto(salario float64) float64 {
	if salario > 150000 {
		return salario * 0.1
	}

	if salario > 50000 {
		return salario * 0.17
	}

	return 0
}

func main() {
	var salario1, salario2, salario3 float64 = 35000, 55000, 165000

	fmt.Printf("El impuesto a descontar es %v\n", calcImpuesto(salario1))
	fmt.Printf("El impuesto a descontar es %v\n", calcImpuesto(salario2))
	fmt.Printf("El impuesto a descontar es %v\n", calcImpuesto(salario3))
}
