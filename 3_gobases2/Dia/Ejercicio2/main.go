/*
Ejercicio 2 - Calcular promedio

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el
promedio y un error en caso que uno de los números ingresados sea negativo

*/

package main

import (
	"errors"
	"fmt"
)

func alumnos(valores ...int) (int, error) {

	sum := 0
	for i := 0; i < len(valores); i++ {
		if valores[i] < 0 {
			return 0, errors.New("Los valores no pueden ser negativos")
		}
		sum = sum + valores[i]
	}
	sum = sum / len(valores)
	return sum, nil
}

func main() {

	result, error := alumnos(10, 0, 10, 8, 9)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("El promedio es: ", result)
	}
}
