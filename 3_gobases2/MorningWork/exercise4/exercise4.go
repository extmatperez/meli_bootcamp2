/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso,
requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función
( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

Ejemplo:
const (
   minimo = "minimo"
   promedio = "promedio"
   maximo = "maximo"
)

...

minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...

valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
*/
package main

import "fmt"

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func getValueMin(numbers ...int) int {
	var minimo int = numbers[0]
	for _, number := range numbers {
		if number < minimo {
			minimo = number
		}
	}
	return minimo
}

func getValueMax(numbers ...int) int {
	var maximo int = numbers[0]
	for _, number := range numbers {
		if number > maximo {
			maximo = number
		}
	}
	return maximo
}

func getAverage(numbers ...int) int {
	var promedio int = 0
	var count int
	for index, number := range numbers {
		promedio += number
		count = index
	}
	return promedio / count
}

func errFunc(number ...int) int {
	return -9999
}

func operacion(tipoOperacion string) (func(number ...int) int, string) {
	var message string = "OK"
	switch tipoOperacion {
	case minimo:
		return getValueMin, message
	case maximo:
		return getValueMax, message
	case promedio:
		return getAverage, message
	default:
		message = "Funcion no definida"
		return errFunc, message
	}
}

func main() {
	minFunc, errMinimo := operacion(minimo)
	maxFunc, errMaximo := operacion(maximo)
	promFunc, errPromedio := operacion(promedio)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("valor minimo: %d, error minimo: %s \n", valorMinimo, errMinimo)
	fmt.Printf("valor maximo: %d, error maximo: %s \n", valorMaximo, errMaximo)
	fmt.Printf("valor promedio: %d, error promedio: %s \n", valorPromedio, errPromedio)

}
