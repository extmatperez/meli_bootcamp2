package main

import (
	"fmt"
)

func calculateSalary(minutes float64, category string) float64 {
	hours := minutes / 60

	switch category {
	case "A":
		return hours * 1000
	case "B":
		return hours * 1500 * 1.2
	case "C":
		return hours * 3000 * 1.5
	default:
		return 0
	}
}

func main() {
	/*
		Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

		Si es categoría C, su salario es de $1.000 por hora
		Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
		Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

		Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
	*/
	var minutes float64
	var categoty string

	fmt.Printf("Ingrese los minutos trabajados del empleado: ")
	fmt.Scanf("%f", &minutes)

	fmt.Printf("Ingrese la categoria del empleado: ")
	fmt.Scanf("%s", &categoty)

	salary := calculateSalary(minutes, categoty)

	fmt.Printf("\nEl salario del empleado es: %.2f\n", salary)
}
