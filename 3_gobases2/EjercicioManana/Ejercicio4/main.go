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
		return minimoValor, nil
	case promedio:
		return promedioValor, nil
	case maximo:
		return maximoValor, nil
	}

	return nil, errors.New("Operacion no encontrada")
}

func minimoValor(valores ...int) float64 {
	min := float64(valores[0])
	for _, value := range valores {
		if float64(value) < min {
			min = float64(value)
		}
	}
	return min
}

func maximoValor(valores ...int) float64 {
	max := float64(valores[0])
	for _, value := range valores {
		if float64(value) > max {
			max = float64(value)
		}
	}
	return max
}

func promedioValor(valores ...int) float64 {
	suma := 0.0
	for _, value := range valores {
		suma += float64(value)
	}
	return suma / float64(len(valores))
}

func main() {

	minFunc, err := operacion(minimo)
	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("\nEl minimo es %.2f\n", valorMinimo)

	promFunc, err := operacion(promedio)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("\nEl minimo es %.2f\n", valorPromedio)

	maxFunc, err := operacion(maximo)

	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("\nEl minimo es %.2f\n", valorMaximo)

	if err != nil {
		fmt.Println(err)
	}

}
