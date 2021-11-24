package main

import (
	"errors"
	"fmt"
)

/* Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y
un error en caso que uno de los números ingresados sea negativo
*/

func main() {
	promedio, err := promedio(7, 8, 6, 7, 9, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio es %v\n", promedio)
	}
}

func promedio(notas ...int) (int, error) {
	contador := 0
	suma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("hay alguna nota que es negativa")
		}
		suma += nota
		contador++
	}

	return suma / contador, nil
}
