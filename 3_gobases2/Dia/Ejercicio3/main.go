/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en
la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y
la categoría, y que devuelva su salario.

*/

package main

import "fmt"

func calcular(min float64, cat string) float64 {
	var hora float64 = min / 60
	switch cat {
	case "c":
		return 1000 * hora
	case "b":
		return 1500*hora + 1500*1.2
	case "a":
		return 3000*hora + 3000*1.5
	}
	return 0
}

func main() {
	salarioC := calcular(100, "c")
	fmt.Printf("El salario para este empleado es: %.2f \n", salarioC)
	salarioB := calcular(500, "b")
	fmt.Printf("El salario para este empleado es: %.2f \n", salarioB)
	salarioA := calcular(1000, "a")
	fmt.Printf("El salario para este empleado es: %.2f \n", salarioA)
}
