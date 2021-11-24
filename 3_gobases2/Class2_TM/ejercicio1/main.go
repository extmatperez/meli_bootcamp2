package main

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados
al momento de depositar el sueldo, para cumplir el objetivo es necesario
crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará
un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/

func impuesto(salario float64) float64 {
	if salario > 50000 && salario < 150000 {
		ret := ((salario / 100) * 17)
		return salario - ret
	}
	if salario > 150000 {
		ret := ((salario / 100) * 27)
		return salario - ret
	}
	return 0
}

func main() {

	salario := 180000.0

	calc := impuesto(salario)
	fmt.Printf("El salario con impuestos descontados es: %.2f\n", calc)
}
