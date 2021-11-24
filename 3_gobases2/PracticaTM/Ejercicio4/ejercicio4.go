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
	min := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] < min {
			min = valores[i]
		}
	}
	return min
}

func maxFunc(valores ...int) int {
	max := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] > max {
			max = valores[i]
		}
	}
	return max
}

func promFunc(valores ...int) int {
	contador := 0
	acum := 0
	for _, valor := range valores {
		contador++
		acum += valor
	}
	return acum / contador
}

func main() {
	minFunc, err := operacion(minimo)
	if err == nil {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El valor minimo es %v\n", valorMinimo)
	} else {
		fmt.Println(err)
	}

	promFunc, err := operacion(promedio)
	if err == nil {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El valor promedio es %v\n", valorPromedio)
	} else {
		fmt.Println(err)
	}

	maxFunc, err := operacion(maximo)
	if err == nil {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El valor maximo es %v\n", valorMaximo)
	} else {
		fmt.Println(err)
	}

	pruebaFunc, err := operacion("prueba")
	if err == nil {
		valorPrueba := pruebaFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El valor maximo es %v\n", valorPrueba)
	} else {
		fmt.Println(err)
	}
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
		return nil, errors.New("No existe ninguna función asociada al operador " + operador)
	}
}
