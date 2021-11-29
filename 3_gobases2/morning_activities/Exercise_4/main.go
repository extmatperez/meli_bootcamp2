/* Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso,
requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra
función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se
indicó en la función anterior
*/

package main

import "fmt"

func minimum(values ...float64) float64 {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func maximum(values ...float64) float64 {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func promedy(values ...float64) float64 {
	add := 0.0
	size := float64(len(values))

	for i := 0; i < len(values); i++ {
		add += values[i]
	}
	return add / size
}

func operations(operation string) func(values ...float64) float64 {

	switch operation {
	case "min":
		return minimum
	case "max":
		return maximum
	case "promedy":
		return promedy
	}
	return nil
}

func main() {
	operation := operations("promedy")
	result := operation(8.5, 9, 9.5, 8.8, 3.2, 9.9)
	fmt.Println(result)
}
