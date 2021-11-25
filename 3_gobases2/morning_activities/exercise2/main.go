/* Ejercicio 2 - Calcular promedio

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los
números ingresados sea negativo. */

package main

import (
	"errors"
	"fmt"
)

func notes_promedy(notes ...float64) (float64, error) {
	result := 0.0
	notes_len := float64(len(notes))
	var promedy float64

	for _, note := range notes {
		if note < 0 {
			fmt.Println(errors.New("Negative Number..."))
		} else {
			result += note
		}
	}
	promedy = result / notes_len
	fmt.Println("Promedy: ", promedy)
	return promedy, nil
}

func main() {
	notes_promedy(1, 2, 3, 4, 5, 6, 7)
	notes_promedy(1, 2, 3, 4, -5)
}
