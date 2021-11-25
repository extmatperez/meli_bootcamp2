package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedio(2, 4, 6, 2, 10, 4, 6)
	fmt.Printf("Promedio válido: %.2f. %s\n", promedio, err)
	promedio, err = calcularPromedio(2, 4, 6, -2, 10, 4, 6)
	fmt.Printf("Promedio inválido: %.2f. %s\n", promedio, err)

}

func calcularPromedio(notas ...int) (float64, error) {
	promedio := 0
	hayNegativo := false
	for _, nota := range notas {
		if nota < 0 {
			hayNegativo = true
		}
		promedio += nota
	}
	if hayNegativo {
		return 0, errors.New("Las notas no pueden ser negativas")
	} else {
		return float64(promedio) / float64(len(notas)), nil
	}
}
