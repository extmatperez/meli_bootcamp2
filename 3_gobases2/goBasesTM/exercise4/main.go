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

func operacion(operador string) (func(valores ...int) float64, error) {
	switch operador {
	case minimo:
		return minFunc, nil
	case maximo:
		return maxFunc, nil
	case promedio:
		return promFunc, nil
	default:
		return nil, errors.New("operacion invalida")
	}
}

func minFunc(valores ...int) float64 {
	var minimo int = valores[0]
	for _, valor := range valores {
		if valor < minimo {
			minimo = valor
		}
	}
	return float64(minimo)

}

func maxFunc(valores ...int) float64 {
	var maximo int = valores[0]
	for _, valor := range valores {
		if valor > maximo {
			maximo = valor
		}
	}
	return float64(maximo)
}

func promFunc(valores ...int) float64 {
	var cant int = len(valores)
	var suma int
	for _, valor := range valores {
		suma += valor
	}
	return float64(suma / cant)

}
