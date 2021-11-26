// Ejercicio 3 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

package main

import (
	"fmt"
)

func calculate_salari(time_work int, porc int, salari_hour int) int {
	porcent := porc
	salari := time_work * salari_hour
	if porcent != 0 {
		salari_month := (salari_hour * 8) * 22
		salari = salari + salari_month*porc/100
		return salari
	} else {
		return salari
	}
}

func manager(time_work int, category string) int {
	porc := 0
	switch category {
	case "C":
		return calculate_salari(time_work, porc, 1000)
	case "B":
		return calculate_salari(time_work, 20, 1500)
	case "A":
		return calculate_salari(time_work, 50, 3000)
	default:
		return 0
	}
}

func main() {
	fmt.Println(manager(300, "A"))
}
