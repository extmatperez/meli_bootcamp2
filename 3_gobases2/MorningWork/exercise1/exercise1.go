/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/
package main

import "fmt"

func calculate_taxes(salary float64) float64 {
	var impuesto float64 = 0.0
	if salary >= 150000.00 {
		impuesto = salary*0.83 - salary*0.1
	} else if salary > 50000 && salary < 150000 {
		impuesto = salary * 0.83
	} else {
		impuesto = salary
	}
	return impuesto
}

func main() {
	var salary float64 = 151000.00

	var impuesto float64 = calculate_taxes(salary)

	fmt.Printf("El valor es: %f", impuesto)
}
