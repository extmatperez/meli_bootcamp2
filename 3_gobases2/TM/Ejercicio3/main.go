/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la
categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que
devuelva su salario.
*/
package main

import (
	"fmt"
)

var salario float64

func salary(categoria string, minutos float64) {
	if categoria == "a" {
		salario = minutos * 50.00 * 1.50
		fmt.Printf("Su salario es de %f\n", salario)
	} else if categoria == "b" {
		salario = minutos * 25.00 * 1.20
		fmt.Printf("Su salario es de %f\n", salario)
	} else if categoria == "c" {
		salario = minutos * (1000.00 / 60.00)
		fmt.Printf("Su salario es de %f\n", salario)
	}

	//return salario
}

func main() {
	var categoria string
	var minutos float64

	/*categoria = "c"
	minutos = 50*/
	fmt.Println("Indique la categoría del trabajador:")
	fmt.Scanf("%s", &categoria)
	fmt.Println("Indique los minutos trabajados")
	fmt.Scanf("%f", &minutos)
	salary(categoria, minutos)

}
