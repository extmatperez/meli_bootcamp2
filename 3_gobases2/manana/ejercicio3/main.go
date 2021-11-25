package main

import (
	"fmt"
)

func main() {
	salario := calcSalario(10000, "C")
	fmt.Println("el salario a pagar es :", salario)
}

func calcSalario(min float64, categoria string) float64 {
	salario := 0.0
	switch categoria {
	case "A":
		salario = min * (1000 / 60)
	case "B":
		salario = min * (1500 / 60)
		salario += (1500 * 160 * 0.2)
	case "C":
		salario = min * (3000 / 60)
		salario += (3000 * 160 * 0.50)
	default:
		salario = -999999
	}

	return salario
}
