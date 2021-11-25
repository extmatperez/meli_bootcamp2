// Ejercicio 2 - Calcular promedio

// Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

package main

import (
	"errors"
	"fmt"
)

func calculate_average(notes ...int) (int, error) {
	var resultado int
	length := len(notes)
	for _, note := range notes {
		if note < 0 {
			return 0, errors.New("La nota no puede ser un numero negativo")
		}
		resultado += note
	}
	return resultado / length, nil
}

func main() {
	res, err := calculate_average(10, 5, 6, -1, 8)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", res)
	}
}
