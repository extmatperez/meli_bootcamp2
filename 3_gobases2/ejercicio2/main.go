package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := promedio(1, 2, -3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de este alumno es: %v\n", promedio)
	}
}

func promedio(notas ...int) (int, error) {
	suma := 0
	for _, nota := range notas {
		suma += nota
		if nota < 0 {
			return 0, errors.New("no puede haber calificaciones negativas")
		}
	}
	return (suma / len(notas)), nil
}
