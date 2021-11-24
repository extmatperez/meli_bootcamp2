package main

import "fmt"

func main() {
	fmt.Println(calcularImpuestoSalario(200000))
}

func calcularImpuestoSalario(salario float64) float64 {
	const impuestoMas50 = 0.17
	const impuestoMas150 = 0.1

	if (salario < 50000) {
		return 0
	} else {
		descuento := impuestoMas50;
		if (salario > 150000) {
			descuento += impuestoMas150; 
		}
		return descuento * salario
	}
} 