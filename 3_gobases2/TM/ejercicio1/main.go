package main

import "fmt"

func impuestoSalarial(salario float64) (float64, int) {
	var impuesto float64
	if salario > 50000 {
		impuesto = salario * 0.17
		return impuesto, 17
	} else if salario > 150000 {
		impuesto = salario * 0.27
		return impuesto, 27
	} else {
		return 0.0, 0
	}
}

func main() {
	fmt.Println("Ingrese su salario")
	var salario int
	fmt.Scanf("%d", &salario)
	impuesto, descuento := impuestoSalarial(float64(salario))
	fmt.Printf("El impuesto es de un %d porciento y son: $%.2f  \n", descuento, impuesto)
}
