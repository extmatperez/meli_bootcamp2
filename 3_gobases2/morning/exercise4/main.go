package main

import (
	"errors"
	"fmt"
)

func getMin(numbers []float64) float64 {
	var min float64 = numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}
	return min
}
func getMax(numbers []float64) float64 {
	var max float64 = numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	return max
}
func getAvg(numbers []float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func getErr(numbers []float64) float64 {
	return 0
}

func calcType(data string) (func(numbers []float64) float64, error) {

	switch data {
	case "minimo":
		return getMin, nil
	case "maximo":
		return getMax, nil
	case "promedio":
		return getAvg, nil
	default:
		return getErr, errors.New("Valor incorrecto")
	}

}

func main() {
	/*
		Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

		Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
	*/

	var data string
	var numbers []float64
	var cant int

	fmt.Printf("Ingrese el tipo de cálculo que desea realizar: ")
	fmt.Scanf("%s", &data)

	fmt.Printf("Ingrese la cantidad de números que desea ingresar: ")
	fmt.Scanf("%d", &cant)

	for i := 0; i < cant; i++ {
		var number float64
		fmt.Printf("Ingrese el número %d: ", i+1)
		fmt.Scanf("%f", &number)
		numbers = append(numbers, number)
	}

	calc, err := calcType(data)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El resultado es: %.2f\n", calc(numbers))
	}
}
