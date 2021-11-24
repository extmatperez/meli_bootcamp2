/*
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva
el promedio y un error en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(calificaciones ...int) (float64, error) {

	acumulador := 0.0

	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("se ha ingresado un número negativo")
		}

		acumulador += float64(calificacion)
	}

	promedio := acumulador / float64(len(calificaciones))

	return promedio, nil
}

func main() {

	promedio, error := calcularPromedio(8, -9, 10, 8)

	if error != nil {
		fmt.Printf("Error: %v\n", error)
	} else {
		fmt.Printf("El promedio del alumno es: %v\n", promedio)
	}
}