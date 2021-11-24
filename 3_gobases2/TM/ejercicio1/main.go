/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento
de depositar el sueldo, para cumplir el objetivo es necesario crear una función que
devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del
sueldo y si gana más de $150.000 se le descontará además un 10%.
*/

package main

import "fmt"

func calcularImpuesto(sueldo float64) float64 {
	if sueldo > 50000 {
		return sueldo * 0.17
	}

	if sueldo > 150000 {
		return sueldo * 0.27
	}

	return 0.0
}

func main() {
	fmt.Printf("El impuesto es de $%.2f", calcularImpuesto(160000.0))
}
