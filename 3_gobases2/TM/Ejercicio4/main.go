/* Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un
curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva
otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el
cálculo que se indicó en la función anterior
*/

package main

import (
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	op := calc(maximo)
	res := op(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(res)
}

func calc(operacion string) func(valores ...int) int {
	switch operacion {
	case "minimo":
		return minFunc
	case "maximo":
		return maxFunc
	case "promedio":
		return promFunc
	default:
		fmt.Println("El cálculo no está definido")
	}
	return nil
}

func minFunc(valores ...int) int {
	min := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] < min {
			min = valores[i]
		}
	}
	return min
}
func maxFunc(valores ...int) int {
	max := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] > max {
			max = valores[i]
		}
	}
	return max
}
func promFunc(valores ...int) int {
	sum := 0
	length := len(valores)
	for i := 0; i < len(valores); i++ {
		sum += valores[i]
	}
	prom := sum / length
	return prom
}

/* minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...

valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
*/
