package main

import "fmt"

func calcularImpuesto(salario float64) float64 {

	impuesto := 0.0

	if salario > 50000 {
		impuesto = salario * 0.17
	}
	if salario > 150000 {
		impuesto = impuesto + salario * 0.1
	}

	return impuesto
}

func main() {

	var salario float64
	fmt.Println("Ingresa el salario: ")
	fmt.Scanf("%f", &salario)


	fmt.Printf("El impuesto es de %2.f \n", calcularImpuesto(float64(salario)))

}
