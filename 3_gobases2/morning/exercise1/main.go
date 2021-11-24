package main

import "fmt"

func discount(salary float64) float64 {
	switch {
	case salary > 150000:
		return salary * 0.1
	case salary > 50000:
		return salary * 0.17
	default:
		return 0
	}
}

func main() {
	/*
		Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
		Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
	*/
	var salary float64

	fmt.Println("Ingrese el sueldo: ")
	fmt.Scanf("%f", &salary)

	fmt.Println(discount(salary))
}
