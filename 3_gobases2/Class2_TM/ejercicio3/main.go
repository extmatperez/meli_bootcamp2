package main

import "fmt"

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

*/
func calulate(category string, minutes int) int {
	salaries := map[string]int{
		"A": (3000 / 60) * 1.5,
		"B": (1500 / 60) * 1.2,
		"C": (1000 / 60),
	}
	fmt.Println(minutes * salaries[category])
	return minutes * salaries[category]
}

func main() {
	calulate("A", 1800)

}
