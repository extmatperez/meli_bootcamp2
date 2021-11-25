package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func main() {
	minFunc, err := operacion(minimo)
	maxFunc, err := operacion(maximo)
	promFunc, err := operacion(promedio)

	if err != nil {
		fmt.Println(err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Printf("Minimo: %v\n", valorMinimo)
		fmt.Printf("Maximo %v\n", valorMaximo)
		fmt.Printf("Promedio: %v\n", valorPromedio)
	}
}

func operacion(operacion string) (func(valores ...float64) float64, error) {
	switch operacion {
	case minimo:
		return operacionMin, nil
	case maximo:
		return operacionMax, nil
	case promedio:
		return operacionProm, nil
	default:
		return nil, errors.New("operacion no definida")
	}
}

func operacionMin(valores ...float64) float64 {
	min := valores[0]
	for _, numero := range valores {
		if numero < min {
			min = numero
		}
	}
	return min
}

func operacionMax(valores ...float64) float64 {
	max := 0.0
	for _, numero := range valores {
		if numero > max {
			max = numero
		}
	}
	return max
}

func operacionProm(valores ...float64) float64 {
	prom := 0.0
	cant := 0.0
	for _, numero := range valores {
		prom += numero
		cant++
	}
	prom = prom / cant
	return prom
}
