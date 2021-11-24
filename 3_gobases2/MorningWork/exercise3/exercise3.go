/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/
package main

import "fmt"

func salary_calculator(minutes int, category string) float64 {
	switch category {
	case "C":
		constant_mult := float64((1000.0 / 60.0))
		return constant_mult * float64(minutes)
	case "B":
		constant_mult := float64((1500.0 / 60.0) * 1.2)
		return constant_mult * float64(minutes)
	default:
		constant_mult := float64((3000.0 / 60.0) * 1.5)
		return constant_mult * float64(minutes)
	}
}

func main() {
	salary := salary_calculator(100, "A")

	fmt.Printf("El salario resultante es: %f", salary)

}
