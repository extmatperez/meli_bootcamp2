/* Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

func category_a(minutes float64) float64 {
	hours := minutes / 60
	salary := hours * 3000
	response := salary + (salary * 0.5)
	fmt.Printf("The corresponding salary to the employee category A is:  $%.2f \n", response)
	return response
}

func category_b(minutes float64) float64 {
	hours := minutes / 60
	salary := hours * 1500
	response := salary + (salary * 0.2)
	fmt.Printf("The corresponding salary to the employee category B is: $%.2f \n", response)
	return response
}
func category_c(minutes float64) float64 {
	hours := minutes / 60
	salary := hours * 1000
	response := salary
	fmt.Printf("The corresponding salary to the employee category C is: $%.2f \n", response)
	return response
}

func main() {
	var category string
	var minutes float64
	fmt.Printf("Enter the employee category: ")
	fmt.Scanf("%f\n", &category)
	fmt.Printf("Enter the employee minutes worked: ")
	fmt.Scanf("%f\n", &minutes)

	/* Fix the bug in switch statement */

	switch category {
	case "a":
		fmt.Println(category_a(minutes))
	case "b":
		fmt.Println(category_b(minutes))
	case "c":
		fmt.Println(category_c(minutes))
	default:
		fmt.Println("Something went wrong, please try again!")
	}

}
