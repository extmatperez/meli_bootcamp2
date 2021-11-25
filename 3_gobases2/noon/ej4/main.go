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

func minFunc(valores ...int) int {
	var min int = 999999999
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}

	return min
}

func maxFunc(valores ...int) int {
	var max int = -999999999
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}

	return max
}

func promFunc(valores ...int) int {
	var prom int = 0
	for _, valor := range valores {
		prom = prom + valor
	}

	return prom / len(valores)
}

func operacion(operador string) (func(valores ...int) int, error) {
	switch operador {
	case minimo:
		return minFunc, nil
	case promedio:
		return promFunc, nil
	case maximo:
		return maxFunc, nil
	default:
		return nil, errors.New("la operación no existe")
	}
}

func pruebaOperacion(operador string) {
	opFunc, err := operacion(operador)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El valor devuelto por la operacion %s es de %d\n", operador, opFunc(2, 3, 3, 4, 1, 2, 4, 5))
	}
}

func main() {
	pruebaOperacion(minimo)
	pruebaOperacion(promedio)
	pruebaOperacion(maximo)
	pruebaOperacion("otra distinta")
}
