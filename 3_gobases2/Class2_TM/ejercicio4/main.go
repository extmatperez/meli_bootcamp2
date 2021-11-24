package main

import (
	"errors"
	"fmt"
)

/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo,
máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
(mínimo, máximo o promedio) y que devuelva otra función ( y un error en caso que el
cálculo no esté definido) que se le puede pasar una cantidad N de enteros y
devuelva el cálculo que se indicó en la función anterior

*/

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	minFunc, err := operaciones(minimo)
	promFunc, err := operaciones(promedio)
	maxFunc, err := operaciones(maximo)
	if err != nil {
		fmt.Printf("Hubo un error: %v \n", err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5, 7)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Printf("minimo: %v \n", valorMinimo)
		fmt.Printf("promedio: %v \n", valorPromedio)
		fmt.Printf("maximo %v \n", valorMaximo)
	}
}

func operaciones(operacion string) (func(valores ...int) float64, error) {
	switch operacion {
	case minimo:
		return minFunc, nil
	case maximo:
		return maxFunc, nil
	case promedio:
		return promFunc, nil
	default:
		return nil, errors.New("invalid operacion")
	}
}

func minFunc(valores ...int) float64 {
	var min int = valores[0]
	for _, val := range valores {
		if val < min {
			min = val
		}
	}
	return float64(min)
}

func maxFunc(valores ...int) float64 {
	var max int = valores[0]
	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	return float64(max)
}

func promFunc(valores ...int) float64 {
	var total int = len(valores)
	var suma float64 = 0
	for _, val := range valores {
		suma += float64(val)
	}
	var promedio float64 = suma / float64(total)
	return promedio
}
