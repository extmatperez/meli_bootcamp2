package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (float64, error) {
	var promedio int
	for _, valor := range notas {
		if valor < 0 {
			return 0, errors.New("existe una nota negativa")
		}
		promedio += valor
	}
	return float64(promedio) / float64(len(notas)), nil
}

func main() {

	prom, err := promedio(1, 2, 3, 4, -5)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("El promedio de 1,2,3,4,5 es: ", prom)
	}
}
