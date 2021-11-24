package main

import (
	"errors"
	"fmt"
)

func promedioAlumno(notas ...int) (float64, error) {
	suma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Hay una nota que es negativa")
		}
		suma += nota
	}
	cantidadTotal := len(notas)
	var promedio float64 = float64(suma) / float64(cantidadTotal)
	return promedio, nil
}

func main() {
	promedio, err := promedioAlumno(2, 4, -7, 8, 9)

	if err == nil {
		fmt.Printf("El promedio del alumno es: %v \n", promedio)
	} else {
		fmt.Printf("Hay un error: %v \n", err)
	}
}
