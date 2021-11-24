package main

import "fmt"

/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes
y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
y que devuelva su salario.
*/

func main() {
	categoria := "A"
	minutosTrabajados := 12000

	fmt.Printf("El salario es %v\n", calcularSalario(categoria, minutosTrabajados))
}

func calcularSalario(categoria string, minutos int) float64 {
	horas := minutos / 60
	switch categoria {
	case "C":
		return float64(minutos / 60 * 1000)
	case "B":
		salario := horas * 1500
		return float64(salario) * 1.2
	case "A":
		salario := horas * 3000
		return float64(salario) * 1.5
	}
	return 0
}
