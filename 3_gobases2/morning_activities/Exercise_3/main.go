/* Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

func salary(minutes float64, category string) float64 {
	var hours float64
	var salary_total float64
	hours = minutes / 60

	switch category {
	case "a":
		salary_total = (hours * 3000) * 1.50
	case "b":
		salary_total = (hours * 1500) * 1.20
	case "c":
		salary_total = (hours * 1000)
	default:
		fmt.Println("Something went wrong, please try again!")
	}
	return salary_total
}

func main() {
	/* var category string */
	/* var minutes float64 */

	/* fmt.Printf("Enter the employee category: ")
	fmt.Scanf("%f\n", &category) */
	/* fmt.Printf("Enter the employee minutes worked: ")
	fmt.Scanf("%f\n", &minutes) */

	fmt.Println(salary(600, "a"))
}
