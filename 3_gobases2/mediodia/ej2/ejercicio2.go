package main

import (
	"errors"
	"fmt"
)

func calcPromedio(notas []float64) (float64, error) {
	var suma float64 = 0.0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Se ingreso una nota negativa")
		} else {
			suma = suma + nota
		}
	}

	return suma / (float64(len(notas))), nil
}

func pruebaPromedio(notas ...float64) {
	promedio, err := calcPromedio(notas)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de notas es %.2f\n", promedio)
	}
}

func main() {
	pruebaPromedio(5, 4, 8, 6, 1, 5, 5, 6, 9, 2)
	pruebaPromedio(5, 4, 8, -6, 1, 5, 5, 6, 9, 2)
}
