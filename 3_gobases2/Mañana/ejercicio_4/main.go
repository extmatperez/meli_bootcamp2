package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func operacion(opera string) (func(valores ...int) float64, error) {
	switch opera {
	case minimo:
		return operaMin, nil
	case promedio:
		return operaProm, nil
	case maximo:
		return operaMax, nil
	default:
		return nil, errors.New("Operacion no generada")
	}
}

func operaMin(valores ...int) float64 {
	min := valores[0]
	for _, numero := range valores {
		if numero < min {
			min = numero
		}
	}
	return float64(min)
}

func operaMax(valores ...int) float64 {
	max := valores[0]
	for _, numero := range valores {
		if numero > max {
			max = numero
		}
	}
	return float64(max)
}
func operaProm(valores ...int) float64 {
	var suma float64 = 0
	for _, numero := range valores {
		suma += float64(numero)
	}
	return (suma / float64(len(valores)))
}

func main() {
	minFunc, err := operacion(minimo)
	maxFunc, err := operacion(maximo)
	promFunc, err := operacion(promedio)

	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Printf("Minimo: %v\n", valorMinimo)
		fmt.Printf("Maximo %v\n", valorMaximo)
		fmt.Printf("Promedio: %v\n", valorPromedio)

	}
}
