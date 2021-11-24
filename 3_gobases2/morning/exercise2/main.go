package main

import (
	"errors"
	"fmt"
)

func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	var sum float64

	for _, number := range numbers {
		if number < 0 {
			return 0, errors.New("Negative number")
		}
		sum += number
	}
	return sum / float64(len(numbers)), nil
}
func main() {
	/*
		Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
	*/

	average, err := average(5, 6, 7, 8, 10, 9, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nEl promedio es: %.2f\n", average)
	}

}
