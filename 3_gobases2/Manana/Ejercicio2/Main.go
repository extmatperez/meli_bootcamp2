package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float64) (float64, error) {
	var resultado float64
	resultado = 0

	fmt.Println(notas, len(notas))

	for _, nota := range notas {
		if nota < 0 {
			return 0.0, errors.New("nota negativa")
		} else {
			resultado += nota
		}
	}

	return resultado / float64(len(notas)), nil

}

func main() {

	fmt.Println("Bienvenidos al ejercicio 2")

	fmt.Println("CASO 1 ALUMNO ARIEL")
	promedio1, err := calcularPromedio(1, 6, 9, 6, 2, 2, -1, 6)
	if err == nil {
		fmt.Println("El promedio del alumno1 es", promedio1)
	} else {
		fmt.Println(err)
	}

}
