/*
? Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una
cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

*/

package main

import (
	"errors"
	"fmt"
)

func operacion(operador string) (func(val ...int) float64, error) {
	switch operador {
	case "minimo":
		return min, nil
	case "maximo":
		return max, nil
	case "promedio":
		return prom, nil
	default:
		return nil, errors.New("error: No existe la operación")
	}

}

func min(val ...int) float64 {
	min := val[0]
	for _, v := range val {
		if v < min {
			min = v
		}
	}
	return float64(min)

}
func max(val ...int) float64 {
	max := val[0]
	for _, v := range val {
		if v > max {
			max = v
		}
	}
	return float64(max)
}
func prom(val ...int) float64 {
	var prom float64
	sum := 0.0
	for _, v := range val {
		sum = sum + float64(v)
	}
	prom = sum / float64(len(val))
	return prom
}

func main() {

	operMin, _ := operacion("minimo")
	valores := operMin(5, 3, 4, 2, 7, 1, 8)
	fmt.Printf("El valor minimo es : %v\n", valores)

	operMax, _ := operacion("maximo")
	valores = operMax(5, 3, 4, 2, 7, 1, 8)
	fmt.Printf("El valor maximo es : %v\n", valores)

	operProm, err := operacion("promedio")
	valores = operProm(5, 3, 4, 2, 7, 1, 8)
	// valores = operProm(8, 8, 6, 6, 6, 8)
	fmt.Printf("El promedio es : %4.3v\n", valores)

	if err != nil {
		fmt.Println(err)
	}

}
