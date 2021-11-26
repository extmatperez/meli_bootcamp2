// Ejercicio 4 - Calcular estadísticas

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
// Ejemplo:

// const (
// 	minimo = "minimo"
// 	promedio = "promedio"
// 	maximo = "maximo"
//  )

//  ...

//  minFunc, err := operacion(minimo)
//  promFunc, err := operacion(promedio)
//  maxFunc, err := operacion(maximo)

//  ...

//  valorMinimo := minFunc(2,3,3,4,1,2,4,5)
//  valorPromedio := promFunc(2,3,3,4,1,2,4,5)
//  valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

package main

import "fmt"

func calculate_min(values ...float64) float64 {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func calculate_max(values ...float64) float64 {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func calculate_prom(values ...float64) float64 {
	sum := 0
	length := len(values)
	for i := 0; i < len(values); i++ {
		sum += int(values[i])
	}
	prom := sum / length
	return float64(prom)
}

func operation(operation string) func(values ...float64) float64 {
	switch operation {
	case "minimo":
		return calculate_min

	case "maximo":
		return calculate_max

	case "promedio":
		return calculate_prom
	}
	return nil
}

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	op := operation(minimo)
	res := op(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(res)
}
