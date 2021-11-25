/*Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos
de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y
que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros
y devuelva el cálculo que se indicó en la función anterior*/

package main

import "fmt"

func minimo(notas ...int) float64 {

	min := notas[0]
	for _, v := range notas {
		if v < min {
			min = v
		}
	}
	return float64(min)
}

func maximo(notas ...int) float64 {
	max := notas[0]
	for _, v := range notas {
		if v > max {
			max = v
		}
	}
	return float64(max)
}

func promedio(notas ...int) float64 {

	suma := 0
	promedio := 0.00

	for _, n := range notas {

		suma += n
	}

	cant_notas := len(notas)

	promedio = float64(suma) / float64(cant_notas)

	return promedio
}

func operacion(operador string) func(notas ...int) float64 {

	switch operador {
	case "MIN":
		return minimo
	case "MAX":
		return maximo
	case "PROM":
		return promedio
	}

	return nil
}

func main() {

	fmt.Println(operacion("PROM")(5, 3, 7, 8, 9, 2))
}
