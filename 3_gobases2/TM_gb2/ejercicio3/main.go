package main

import "fmt"

// Ejercicio 2 - calcular Salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func calculate_total (minutes int, rate int) int {
	return minutes * rate

}

func category_rate (category string) int {
	
}


func final_income (category string, total int) int {
	switch category {
	case category == "C":
		return total
	case category == "B":
		return total + total*0.2
	case category == "A":
		return total + total*0.5
	} 
	return 0

}
 


func main () {

	// func income_calculator(minutes int, category string) int

}