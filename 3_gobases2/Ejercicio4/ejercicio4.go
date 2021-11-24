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

func opMin(valores ...float64) float64 {
	var min = 1000000.0
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func opMax(valores ...float64) float64 {
	var max = -999999.0
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

func opProm(valores ...float64) float64 {
	var promedio float64
	for _, valor := range valores {
		if valor < 0 {
			return 0
		}
		promedio += valor
	}
	return float64(promedio) / float64(len(valores))
}

func operacion(operador string) (func(valores ...float64) float64, error) {
	switch operador {
	case minimo:
		return opMin, nil
	case maximo:
		return opMax, nil
	case promedio:
		return opProm, nil
	}
	return nil, errors.New("la operacion no existe")
}

func main() {
	minFunc, err := operacion("minimo")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		valorMinimo := minFunc(1.0, 2.0, 3.0, 4.0, 5.0)
		fmt.Println("El valor minimo es:", valorMinimo)
	}

	maxFunc, err := operacion("maximo")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		valorMaximo := maxFunc(1.0, 2.0, 3.0, 4.0, 5.0)
		fmt.Println("El valor maximo es:", valorMaximo)
	}

	promFunc, err := operacion(promedio)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		valorPromedio := promFunc(10, 2.0, 3.0, 4.0, 5.0)
		fmt.Println("El valor promedio es:", valorPromedio)
	}
}
