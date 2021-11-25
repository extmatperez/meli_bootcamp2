package main

import (
	"errors"
	"fmt"
)

func promedio(calificacion ...float64) (float64, error) {

	suma := 0.0

	for _, valor := range calificacion {
		if valor < 0 {
			return 0.0, errors.New("Calificacion no puede ser negativa")
		}
		suma += valor
	}
	promedio := suma / float64(len(calificacion))

	return promedio, nil
}

func main() {

	respuesta, err := promedio(3, 5, 1, 0, 2, 2, 2, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de calificaciones fue: %.2f\n", respuesta)
	}

}
