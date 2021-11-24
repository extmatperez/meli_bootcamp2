package main

import "fmt"

func main() {
	salario := 0.0
	fmt.Println("Introduzca su salario")
	fmt.Scanf("%f", &salario)
	fmt.Printf("Debera pagar $%.2f de impuestos para un salario de %.2f\n", cobrarImpuestos(salario), salario)
}

func cobrarImpuestos(salario float64) float64 {
	if salario >= 17000 {
		if salario >= 150000 {
			return (salario * 0.27)
		}
		return (salario * 0.17)
	}
	return 0
}
