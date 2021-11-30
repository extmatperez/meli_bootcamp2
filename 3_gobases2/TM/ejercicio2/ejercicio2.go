/*Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo*/

package main

import (
	"errors"
	"fmt"
)

func promedios(notas ...float64) (float64, error) {

	suma := 0.00
	promedio := 0.00

	for _, n := range notas {

		if n <= 0 {
			return 0, errors.New("Numero Negativo")
		}
		suma += n
	}

	cant_notas := len(notas)

	promedio = suma / float64(cant_notas)

	return promedio, nil
}

func main() {

	promedio, errores := promedios(4.0, 5.0, 6.0, 7.0, 8.0, 9.0)

	if errores == nil {
		fmt.Println("El promedio es: ", promedio)
	} else {
		fmt.Println(errores)
	}

}
