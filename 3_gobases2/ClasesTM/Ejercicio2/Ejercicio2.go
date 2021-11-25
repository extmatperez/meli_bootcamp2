package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calculadorDePromedios(8, 8, 8, 9, 10, 10, -2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", promedio)

	}
}

func calculadorDePromedios(notas ...float64) (float64, error) {

	notasSum := 0.0

	for _, value := range notas {

		if value < 0 {
			return 0, errors.New("No se puede ingresar numeros negativos")
		}
		notasSum = notasSum + value
	}
	return notasSum / float64(len(notas)), nil
}
