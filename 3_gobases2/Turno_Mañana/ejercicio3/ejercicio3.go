/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y
la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
y que devuelva su salario.*/

package main

import "fmt"

func calcularSueldo(categoria string, minutos int) float64 {

	cant_horas := float64(minutos / 60)
	salario := 0.00

	switch categoria {
	case "A":
		salario = 3000.0 * cant_horas
		salario += salario * 0.5
	case "B":
		salario = 1500.0 * cant_horas
		salario += salario * 0.2
	case "C":
		salario = 1000.0 * cant_horas
	default:
		salario = 0
	}

	return salario
}

func main() {

	minutos := 0
	categoria := ""

	fmt.Println("Ingrese la categoría: ")
	fmt.Scanf("%f", &categoria)

	fmt.Println("Ingrese los minutos trabajados: ")
	fmt.Scanf("%f", &minutos)

	fmt.Println("El sueldo correspondiente es: ", calcularSueldo(categoria, minutos))
}
