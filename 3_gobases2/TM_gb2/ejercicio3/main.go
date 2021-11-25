package main

import "fmt"

// Ejercicio 2 - calcular Salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.


func calculate_parcial (minutes float64, rate float64) float64 {
	return minutes * rate
}


func category_rate (category string) float64 {
	switch category {
		case "C":
			return 1000
		case "B":
			return 1500
		case "A":
			return 3000
	} 
	return 0

}

func calculate_total_income (employee_category string, minutes float64) float64 {

		switch employee_category {
			case "C":
				return calculate_parcial(category_rate(employee_category), minutes)

			case "B":
				parcial := calculate_parcial(category_rate(employee_category), minutes) 
				return parcial + parcial*0.2	

			case "A":
				parcial := calculate_parcial(category_rate(employee_category), minutes) 
				return parcial + parcial*0.5	
			} 
		return 0

}


 


func main () {

	fmt.Println(calculate_total_income("A", 120))

}