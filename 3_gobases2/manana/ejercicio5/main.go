package main

import "fmt"

func main() {
	impuesto := calcImpuesto(50000.0)
	fmt.Println("el impuesto a pagar es de :", impuesto)
}

func calcImpuesto(salario float64) float64 {

	if salario > 50000 {

		return salario - salario*0.17
	}
	if salario > 150000 {
		return salario - salario*0.27
	}
	return 0
}
