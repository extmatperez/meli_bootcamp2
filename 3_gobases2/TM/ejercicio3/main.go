/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en
la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos
trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import (
	"fmt"
	"math"
)

func minutosAHoras(minutos float64) float64 {
	return math.RoundToEven(minutos / 60)
}

func calcularSalario(categoría string, minutos float64) float64 {

	switch categoría {
	case "C":
		return 1000 * minutosAHoras(minutos)
	case "B":
		return 1500*minutosAHoras(minutos) + (1500*minutosAHoras(minutos))*0.20
	case "A":
		return 3000*minutosAHoras(minutos) + (3000*minutosAHoras(minutos))*0.50
	}

	return 0.0
}

func main() {
	fmt.Printf("%.2f\n", calcularSalario("A", 122))
}
