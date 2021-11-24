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

func operacion(op string) (func(valores ...int) float64, error) {
	switch op {
	case minimo:
		return opMinimo, nil
	case promedio:
		return opPromedio, nil
	case maximo:
		return opMaximo, nil
	default:
		return nil, errors.New("invalid operacion")
	}
}

func opMaximo(valores ...int) float64 {
	var max int = valores[0]
	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	return float64(max)
}
func opMinimo(valores ...int) float64 {
	var min int = valores[0]
	for _, val := range valores {
		if val < min {
			min = val
		}
	}
	return float64(min)
}

func opPromedio(valores ...int) float64 {
	var total int = len(valores)
	var suma float64 = 0
	for _, val := range valores {
		suma += float64(val)
	}
	var promedio float64 = suma / float64(total)
	return promedio
}

func main() {
	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Printf("Hubo un error: %v \n", err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5, 7)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Printf("minimo: %v \n", valorMinimo)
		fmt.Printf("promedio: %v \n", valorPromedio)
		fmt.Printf("maximo %v \n", valorMaximo)
	}
}
