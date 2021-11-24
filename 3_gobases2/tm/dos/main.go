package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float64) (float64, error) {
	promedio := 0.0
	for _, nota := range notas {
		if nota < 0 {
			return 0.0, errors.New("nota negativa")
		} else {
			promedio += nota
		}
	}
	return promedio / float64(len(notas)), nil
}

func main() {
	promedio, err := (calcularPromedio(7.5, 8, 9.5, 8, 10))
	if err != nil {
		fmt.Println("El promedio es: ", promedio)
	} else {
		fmt.Println(err)
	}
}
