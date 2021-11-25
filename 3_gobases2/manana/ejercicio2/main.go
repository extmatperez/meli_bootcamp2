package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, error := calcPromedio(5, 5, -4)
	if error != nil {
		fmt.Println("error :", error)

	} else {
		fmt.Println("el promedio de notas es de :", promedio)

	}
}

func calcPromedio(notas ...float64) (float64, error) {
	promedio := 0.0
	for _, nota := range notas {
		promedio += nota
		if nota < 0 {
			return 0, errors.New("error se ingreso numero negativo")
		}
	}
	promedio = promedio / float64(len(notas))
	return promedio, nil
}
