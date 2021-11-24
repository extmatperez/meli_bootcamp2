package main

import "fmt"

/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si
gana más de $150.000 se le descontará además un 10%.
*/

func main() {
	salario := 130000.0

	impuesto := calcularImpuesto(salario)

	fmt.Printf("El impuesto del salario $%v es $%v\n", salario, impuesto)
}

func calcularImpuesto(salario float64) float64 {
	impuesto := 0.0
	if salario > 50000 {
		impuesto = salario * 0.17
	}
	if salario > 150000 {
		impuesto += salario * 0.1
	}

	return impuesto
}
