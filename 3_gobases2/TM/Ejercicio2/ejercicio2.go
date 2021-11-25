package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := promedios(10, 10, 10, 9, 8, 8, 7, 10)
	if err == nil {
		fmt.Printf("El promedio de calificaciones es: %d\n", promedio)
	} else {
		fmt.Printf("Error: %v", err)
	}

}

func promedios(calificaciones ...int) (int, error) {
	promedio := 0
	for _, c := range calificaciones {
		if c > 0 {
			promedio += c / len(calificaciones)
		} else {
			return 0, errors.New("Una de las calificaciones ingresadas es un numero negativo")
		}
	}
	return promedio, nil
}
