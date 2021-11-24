package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := promedio(20, 50, 50, 15, 12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", res)
	}
}

func promedio(calificaciones ...float64) (float64, error) {
	suma := 0.0
	resultado := 0.0
	for _, value := range calificaciones {
		if value < 0 {
			return resultado, errors.New("se encontró un valor negtivo")
		}
		suma += value
	}

	resultado = suma / float64(len(calificaciones))

	return resultado, nil
}
